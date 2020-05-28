package gmaps

import "git.resultys.com.br/motor/models/gmaps"

// Protocol ...
type Protocol struct {
	Code    int           `json:"code"`
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    gmaps.Company `json:"company"`
}
