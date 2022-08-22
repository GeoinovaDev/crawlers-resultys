package registrobr

import (
	"github.com/GeoinovaDev/crawlers-resultys/registrobr/document"
)

// Protocol ...
type Protocol struct {
	Code    int                `json:"code"`
	Status  string             `json:"status"`
	Data    *document.Document `json:"data"`
	Message string             `json:"message"`
}
