package registrobr

import (
	"git.resultys.com.br/sdk/crawlers-golang/registrobr/document"
)

// Protocol ...
type Protocol struct {
	Code    int                `json:"code"`
	Status  string             `json:"status"`
	Data    *document.Document `json:"data"`
	Message string             `json:"message"`
}
