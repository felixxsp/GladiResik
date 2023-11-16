package entity

type Orders struct {
	Uuid       string `json:"uuid" bson:"uuid"`
	Contents   []Food `json:"contents" bson:"contents"`
	Quantity   []int  `json:"qty" bson:"qty"`
	Completion []bool `json:"completion" bson:"completion"`
}

type Incoming struct {
	FoodID   int `json:"id" bson:"id"`
	Quantity int `json:"qty" bson:"qty"`
}
