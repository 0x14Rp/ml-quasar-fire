package usecase

import (
	"errors"
	"strings"
)

func ProcessMessage(messages ...[]string) (string, error) {

	maxLength := 0
	for _, msg := range messages {
		if len(msg) > maxLength {
			maxLength = len(msg)
		}
	}

	finalMessage := make([]string, maxLength)

	for _, msg := range messages {
		offset := maxLength - len(msg)
		for i, word := range msg {
			if word != "" {
				finalMessage[i+offset] = word
			}
		}
	}
	// Validar que no haya posiciones vacías en el mensaje final
	for _, word := range finalMessage {
		if word == "" {
			return "", errors.New("insuficiente información para reconstruir el mensaje")
		}
	}

	return strings.Join(finalMessage, " "), nil
}
