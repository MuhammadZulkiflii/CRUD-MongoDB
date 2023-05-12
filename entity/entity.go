package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Data struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Nama string             `json:"nama" bson:"nama"`
}

func Create(nama string) *Data {
	return &Data{
		Id:   primitive.NewObjectID(),
		Nama: nama,
	}
}
