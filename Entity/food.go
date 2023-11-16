package entity

type Food struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Category    int    `json:"category" bson:"category"`
	Description string `json:"description" bson:"description"`
	Status      bool   `json:"status" bson:"status"`
}
