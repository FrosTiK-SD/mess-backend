package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type UpdateMeal struct {
	MessID     primitive.ObjectID   `json:"messID,omitempty"`
	MealType   primitive.ObjectID   `json:"mealTypeID,omitempty"`
	Dates      []primitive.DateTime `json:"dates,omitempty"`
	Menu       []primitive.ObjectID `json:"menu,omitempty"`
	AppendOnly bool                 `json:"appendOnly,omitempty"`
}
