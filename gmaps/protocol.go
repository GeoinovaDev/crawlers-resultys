package gmaps

// Protocol ...
type Protocol struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    Company `json:"company"`
}
