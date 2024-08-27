package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMessage_Success(t *testing.T) {
	messages := [][]string{
		{"este", "", "", "mensaje", ""},
		{"este", "", "un", "", ""},
		{"", "es", "", "", "secreto"},
	}

	expectedMessage := "este es un mensaje secreto"
	actualMessage := ProcessMessage(messages...)

	assert.Equal(t, expectedMessage, actualMessage)
}

func TestGetMessage_Error(t *testing.T) {
	messages := [][]string{
		{"este", "", "mensaje", ""},
		{"", "es", "", "secreto"},
		{"", "este", "un", "", ""},
	}

	expectedMessage := ""
	actualMessage := ProcessMessage(messages...)

	assert.NotEqual(t, expectedMessage, actualMessage)
}
