package gmaps

// Company ...
type Company struct {
	Name     string    `json:"name"`
	LatLng   []string  `json:"latlng"`
	Address  string    `json:"address"`
	Phones   []string  `json:"phone"`
	Sites    []string  `json:"site"`
	Comments []Comment `json:"comments"`
	Status   string    `json:"status"`
}
