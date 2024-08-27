package usecase

import (
	"ml-quasar-fire/internal/entity"
)

func ProcessTopSecretRequest(distances []float32, messages [][]string) (entity.Position, string, error) {

	x, y, err := GetLocation(distances...)
	if err != nil {
		return entity.Position{}, "", err
	}

	message, err := ProcessMessage(messages...)

	return entity.Position{X: x, Y: y}, message, err
}
