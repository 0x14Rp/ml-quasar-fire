package usecase

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLocation_InvalidInput(t *testing.T) {

	distances := []float32{538.52, 141.42, 122.34, 122.4}
	x, y, err := GetLocation(distances...)
	assert.NotNil(t, err)
	assert.Equal(t, "distancias insuficientes", err.Error())
	assert.InDelta(t, 0, x, 0.01)
	assert.InDelta(t, 0, y, 0.01)
}

func TestGetLocation(t *testing.T) {

	distances := []float32{100, 115.5, 142.7}
	x, y, err := GetLocation(distances...)
	assert.Nil(t, err)
	assert.InDelta(t, -487.2859, x, 0.0001)
	assert.InDelta(t, 1557.0142, y, 0.0001)

}
