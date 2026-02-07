package utils

import (
	"log"

	"github.com/google/uuid"
)

func GenerateUuid() string {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Fatal("failed to generate uuid")
	}

	return uuid.String()
}
