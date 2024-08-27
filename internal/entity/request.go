package entity

type TopSecretRequest struct {
	Satellites []Satellite `json:"satellites"`
}

type SatelliteData struct {
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
