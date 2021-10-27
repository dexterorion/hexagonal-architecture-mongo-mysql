package helpers

import "go.mongodb.org/mongo-driver/bson/primitive"

func SafeObjectIdFromString(s string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(s)
	return oid
}
