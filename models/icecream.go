package models

type Icecream struct {
	ID     string `json:"id" bson:"_id"`
	Flavor  string `json:"flavor" bson:"flavor"`
	Stock string `json:"stock" bson:"stock"`
}
