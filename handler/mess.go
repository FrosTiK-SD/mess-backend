package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateMess(ctx *fiber.Ctx) error {
	var Mess models.Mess

	if err := ctx.BodyParser(&Mess); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Mess.ID = primitive.NewObjectID()

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if result, err := collection.InsertOne(ctx.Context(), Mess); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}
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
