package google

import (
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/lib/convert"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
)

// SearchTelefones pesquisa telefones na pagina principal do google
// Retorna array de telefones
func SearchTelefones(nome string, cidade string) []string {
	telefones := []string{}

	response := request.Get(str.Format("http://35.198.20.221/google/search/phone?nome={0}&cidade={1}", nome, cidade))
	if response != nil {
		telefones = convert.ArrayInterfaceToArrayString(response.([]interface{}))
	}

	return telefones
}
