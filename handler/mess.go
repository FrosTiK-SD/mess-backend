package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateMess(ctx *fiber.Ctx) error {
	var Mess models.Mess

	if err := ctx.BodyParser(&Mess); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := controller.CreateMess(handler.MongikClient, &Mess); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"mess": Mess,
	})

}

func (h *Handler) GetAllMesses(ctx *fiber.Ctx) error {

	messes, err := controller.GetAllMesses(h.MongikClient)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"messes": messes,
	})
}

func (handler *Handler) GetMess(ctx *fiber.Ctx) error {
	var Mess models.Mess
	messID, errObjID := primitive.ObjectIDFromHex(ctx.Get("messID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": messID}).Decode(&Mess); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Mess with the given Mess ID", Mess, nil))
}

func (handler *Handler) UpdateMess(ctx *fiber.Ctx) error {
	var updateData models.Mess
	messID, errObjID := primitive.ObjectIDFromHex(ctx.Get("messID"))
	if errObjID != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errObjID.Error()})
	}
	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	update := bson.M{
		"$set": updateData,
	}
	result, err := collection.UpdateOne(ctx.Context(), bson.M{"_id": messID}, update)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if result.MatchedCount == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Mess not found"})
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Mess updated successfully", nil, nil))
}

func (handler *Handler) DeleteMess(ctx *fiber.Ctx) error {
	messID, errObjID := primitive.ObjectIDFromHex(ctx.Get("messID"))
	if errObjID != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errObjID.Error()})
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if result, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": messID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else {
		if result.DeletedCount == 0 {
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Mess not found"})
		}
		return ctx.JSON(interfaces.GetGenericResponse(true, "Mess deleted successfully", nil, nil))
	}
}

func (handler *Handler) GetMessDashboard(ctx *fiber.Ctx) error {
	userID, errObjID := primitive.ObjectIDFromHex(ctx.Get("userID"))
	if errObjID != nil {
		return errObjID
	}

	type MessDashboard struct {
		AllocatedToMess []struct {
			ID primitive.ObjectID `json:"_id" bson:"_id"`
		} `json:"allocatedToMess" bson:"allocatedToMess"`
		ManagedMess models.Mess
	}

	var results []MessDashboard

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userID}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "messes"},
					{Key: "localField", Value: "managingDetails.messes"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "managedMesses"},
				},
			},
		},
		bson.D{{Key: "$unwind", Value: "$managedMesses"}},
		bson.D{{Key: "$project", Value: bson.D{{Key: "managedMess", Value: "$managedMesses"}}}},
		bson.D{{Key: "$unset", Value: "_id"}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "let", Value: bson.D{{Key: "mess_id", Value: "$managedMess._id"}}},
					{Key: "as", Value: "allocatedToMess"},
					{Key: "pipeline",
						Value: bson.A{
							bson.D{
								{Key: "$match",
									Value: bson.D{
										{Key: "$expr",
											Value: bson.D{
												{Key: "$eq",
													Value: bson.A{
														"$allocationDetails.mess",
														"$$mess_id",
													},
												},
											},
										},
									},
								},
							},
							bson.D{{Key: "$project", Value: bson.D{{Key: "_id", Value: 1}}}},
						},
					},
				},
			},
		},
	}

	if cursor, err := collection.Aggregate(ctx.Context(), pipeline); err != nil {
		return err
	} else {
		if errBind := cursor.All(ctx.Context(), &results); errBind != nil {
			return errBind
		}
		return ctx.JSON(interfaces.GetGenericResponse(true, "Fetched Mess Dashboard", results, nil))
	}
}
