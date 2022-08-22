package gmaps

import "github.com/GeoinovaDev/models-resultys/gmaps"

// Protocol ...
type Protocol struct {
	Code    int             `json:"code"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    []gmaps.Company `json:"data"`
}
