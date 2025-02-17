package utils

import "github.com/google/uuid"

func GenerateUniqueKey() string {
	uUid := uuid.NewString()
	return NormalizeToken(uUid)
}
