package helpers

import "github.com/google/uuid"

func RandomUUIDAsString() string {
	return uuid.New().String()
}

func SafeUUIDFromString(s string) uuid.UUID {
	val, _ := uuid.Parse(s)
	return val
}
