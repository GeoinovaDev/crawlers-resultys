package guiamais

import (
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/lib/convert"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
)

// SearchTelefones pesquisa telefones no guiamas
// Retorna array de telefones
func SearchTelefones(nome string, cidade string, estado string) []string {
	telefones := []string{}

	response := request.Get(str.Format("http://35.198.20.221/guiamais/search?nome={0}&cidade={1}&estado={2}", nome, cidade, estado))
	if response != nil {
		telefones = convert.ArrayInterfaceToArrayString(response.([]interface{}))
	}

	return telefones
}
