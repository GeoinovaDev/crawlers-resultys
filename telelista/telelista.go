package telelista

import (
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/lib/convert"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
)

// SearchTelefones pesquisa telefones no telelista
// Retorna array de telefones
func SearchTelefones(nome string, estado string) []string {
	telefones := []string{}

	response := request.Get(str.Format("http://35.198.20.221/telelista/search?nome={0}&estado={1}", nome, estado))
	if response != nil {
		telefones = convert.ArrayInterfaceToArrayString(response.([]interface{}))
	}

	return telefones
}
