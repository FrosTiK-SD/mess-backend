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

func (handler *Handler) CreateHostel(ctx *fiber.Ctx) error {
	var Hostel models.Hostel

	if err := ctx.BodyParser(&Hostel); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Hostel.ID = primitive.NewObjectID()

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if result, err := collection.InsertOne(ctx.Context(), Hostel); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Menu Item Created", result, nil))
	}
}

func (handler *Handler) GetHostel(ctx *fiber.Ctx) error {
	var Hostel models.Hostel
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Get("hostelID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": hostelID}).Decode(&Hostel); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Hostel with the given ID", Hostel, nil))
}

func (handler *Handler) GetFullyPopulatedHostel(ctx *fiber.Ctx) error {
	// TODO: add Access Level : Admins and caretakers of that particular hostel
	var FPHostel models.FullyPopulatedHostel
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Get("hostelID"))
	if errObjID != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errObjID.Error()})
	}
	var Hostel models.Hostel
	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": hostelID}).Decode(&Hostel); errFind != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": errFind.Error()})
	}
	// Get caretakers: caretakers are users with managingDetails.hostels containing the hostelID
	caretakersCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if cur, err := caretakersCollection.Find(ctx.Context(), bson.M{"managingDetails.hostels": hostelID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err := cur.All(ctx.Context(), &FPHostel.Caretakers); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	FPHostel.Hostel = Hostel
	// Get rooms: list of Populated rooms with hostelID
	var rooms []models.Room
	roomsCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_ROOMS)
	if cur, err := roomsCollection.Find(ctx.Context(), bson.M{"hostel": hostelID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err := cur.All(ctx.Context(), &rooms); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// update allocatedTo Property of each room in rooms
	for _, room := range rooms {
		/*
			A PopulatedRoom extends Hostel as follows:
				- A property allocatedTo which is an array of StudentMini

				A UserMini only has firstName, lastName, middleName , email and mobile
				A StudentMini extends UserMini with instituteProfile
		*/
		var PRoom models.PopulatedRoom
		PRoom.Hostel = Hostel
		var students []models.StudentMini

		// TODO: test rooms[i].AllocatedTo field props
		PRoom.AllocatedTo = students

		// first find users that have room allocated to them
		var users []models.User
		userCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
		if cur, err := userCollection.Find(ctx.Context(), bson.M{"allocationDetails.room": room.ID}); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		} else if err := cur.All(ctx.Context(), &users); err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		for _, user := range users {
			// create a usermini from the user
			var userMini models.UserMini
			// TODO: how to fetch these fields from the user
			// userMini.FirstName = user.FirstName
			// userMini.LastName = user.LastName
			// userMini.MiddleName = user.MiddleName
			userMini.Email = user.Email
			userMini.Mobile = user.Mobile
			// create student mini from the usermini and instituteProfile
			var studentMini models.StudentMini
			studentMini.UserMini = userMini
			studentMini.InstituteProfile = user.InstituteProfile
			// append student mini to the allocatedTo property of the PRoom
			PRoom.AllocatedTo = append(PRoom.AllocatedTo, studentMini)

		}
		// append PRoom to FPHostel.Rooms
		FPHostel.Rooms = append(FPHostel.Rooms, PRoom)
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Fully Populated Hostel with the given ID", FPHostel, nil))
}

func (handler *Handler) UpdateHostel(ctx *fiber.Ctx) error {
	var updatedHostel models.Hostel

	if err := ctx.BodyParser(&updatedHostel); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Get("hostelID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	filter := bson.M{"_id": hostelID}
	update := bson.M{"$set": updatedHostel}

	if _, err := collection.UpdateOne(ctx.Context(), filter, update); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Hostel Updated", nil, nil))
}

func (handler *Handler) DeleteHostel(ctx *fiber.Ctx) error {
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Get("hostelID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	filter := bson.M{"_id": hostelID}

	if _, err := collection.DeleteOne(ctx.Context(), filter); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Hostel Deleted", nil, nil))
}
