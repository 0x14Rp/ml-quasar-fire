package usecase

import "errors"

type LocationRequest struct {
	Distances []float32 `json:"distances"`
}

var (
	kenobi    = [2]float32{-500, -200}
	skywalker = [2]float32{100, -100}
	sato      = [2]float32{500, 100}
)

func GetLocation(distances ...float32) (x, y float32, err error) {

	if len(distances) != 3 {
		return 0, 0, errors.New("distancias insuficientes")
	}

	x1, y1 := kenobi[0], kenobi[1]
	x2, y2 := skywalker[0], skywalker[1]
	x3, y3 := sato[0], sato[1]

	r1, r2, r3 := distances[0], distances[1], distances[2]

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2

	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3

	denominator := A*E - B*D
	if denominator == 0 {
		return 0, 0, errors.New("no se puede determinar una ubicación única")
	}

	x = (C*E - F*B) / denominator
	y = (A*F - C*D) / denominator

	return x, y, nil
}
