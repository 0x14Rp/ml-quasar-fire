package errs

import (
	"errors"
)

var (
	ErrInvalidSatelliteData = errors.New("Numero de satellites insuficientes")
	ErrNoSolutionFound      = errors.New("no unique solution found for the location")
)
