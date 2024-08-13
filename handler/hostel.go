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

func (h *Handler) GetAllHostels(ctx *fiber.Ctx) error {
	hostels, err := controller.GetAllHostels(h.MongikClient)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"hostels": hostels,
	})
}

func (handler *Handler) CreateHostel(ctx *fiber.Ctx) error {
	var hostel models.Hostel

	if err := ctx.BodyParser(&hostel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.CreateHostel(handler.MongikClient, &hostel)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"hostel": hostel,
	})
}

func (handler *Handler) GetHostelById(ctx *fiber.Ctx) error {
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Params("hostelID"))
	if errObjID != nil {
		return errObjID
	}

	hostel, err := controller.GetHostelById(handler.MongikClient, hostelID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"hostel": hostel,
	})
}

func (handler *Handler) GetFullyPopulatedHostel(ctx *fiber.Ctx) error {
	// TODO: add Access Level : Admins and caretakers of that particular hostel
	FPHostel := models.FullyPopulatedHostel{Rooms: []models.PopulatedRoom{}, Caretakers: []models.UserMini{}}
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Params("hostelId"))
	if errObjID != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errObjID.Error()})
	}
	// get hostel from hostel collection
	var Hostel models.Hostel
	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": hostelID}).Decode(&Hostel); errFind != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": errFind.Error()})
	}
	// Get caretakers from users: caretakers are users with managingDetails.hostels containing the hostelID
	caretakersCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if cur, err := caretakersCollection.Find(ctx.Context(), bson.M{"managingDetails.hostels": hostelID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err := cur.All(ctx.Context(), &FPHostel.Caretakers); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	FPHostel.Hostel = Hostel
	// Get rooms from rooms collection: list of Populated rooms with hostelID
	var rooms []models.Room
	roomsCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_ROOMS)
	if cur, err := roomsCollection.Find(ctx.Context(), bson.M{"hostel": hostelID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err := cur.All(ctx.Context(), &rooms); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Get all users allocated to the rooms of the hostel in one query
	var users []models.User
	userCollection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if cur, err := userCollection.Find(ctx.Context(), bson.M{"allocationDetails.hostel": hostelID}); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if err := cur.All(ctx.Context(), &users); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Create a map of room ID to allocated users
	roomToUsers := make(map[primitive.ObjectID][]models.User)
	for _, user := range users {
		roomID := user.AllocationDetails.Room
		roomToUsers[roomID] = append(roomToUsers[roomID], user)
	}

	// update allocatedTo Property of each room in rooms
	for _, room := range rooms {
		/*
			A PopulatedRoom extends Room as follows:
				- A property allocatedTo which is an array of StudentMini

				A UserMini only has firstName, lastName, middleName , email and mobile
				A StudentMini extends UserMini with instituteProfile
		*/
		var PRoom models.PopulatedRoom
		PRoom.Room = room
		var students []models.StudentMini

		for _, user := range roomToUsers[room.ID] {
			var userMini models.UserMini
			userMini.Email = user.Email
			userMini.Mobile = user.Mobile

			var studentMini models.StudentMini
			studentMini.UserMini = userMini
			studentMini.InstituteProfile = user.InstituteProfile

		}

		PRoom.AllocatedTo = students
		FPHostel.Rooms = append(FPHostel.Rooms, PRoom)
	}
	if FPHostel.Caretakers == nil {
		FPHostel.Caretakers = make([]models.UserMini, 0)
	}

	return ctx.JSON(fiber.Map{
		"hostel": FPHostel,
	})
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
