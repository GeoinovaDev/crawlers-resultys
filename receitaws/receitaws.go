package receitaws

import (
	"strings"

	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/prospecta/models/email"
	"git.resultys.com.br/prospecta/models/empresa"
	"git.resultys.com.br/prospecta/models/socio"
	"git.resultys.com.br/prospecta/models/telefone"
	"git.resultys.com.br/prospecta/models/telefone/fonte"
)

// GetEmpresa busca informações da empresa na receitaws
func GetEmpresa(cnpj string) *empresa.Empresa {
	empresa := &empresa.Empresa{}
	json := make(map[string]interface{})

	url := str.Format("https://www.receitaws.com.br/v1/cnpj/{0}", cnpj)
	err := request.New(url).GetJSON(&json)
	if err != nil {

	}

	empresa.CNPJ = json["cnpj"].(string)
	empresa.RazaoSocial = json["nome"].(string)
	empresa.Fantasia = json["fantasia"].(string)
	empresa.Abertura = json["cnpj"].(string)

	empresa.NaturezaJuridica = json["natureza_juridica"].(string)
	empresa.Logradouro = json["logradouro"].(string)
	empresa.Numero = json["numero"].(string)
	empresa.Complemento = json["complemento"].(string)
	empresa.CEP = json["cep"].(string)
	empresa.Bairro = json["bairro"].(string)
	empresa.Municipio = json["municipio"].(string)
	empresa.UF = json["uf"].(string)
	empresa.EFR = json["efr"].(string)
	empresa.Situacao = json["situacao"].(string)
	empresa.DataSituacao = json["data_situacao"].(string)
	empresa.MotivoSituacao = json["motivo_situacao"].(string)
	empresa.SituacaoEspecial = json["situacao_especial"].(string)
	empresa.DataSituacaoEspecial = json["data_situacao_especial"].(string)
	empresa.CapitalSocial = json["capital_social"].(string)

	empresaEmail := json["email"].(string)
	if len(empresaEmail) > 0 {
		empresa.Emails = []email.Email{
			email.Email{
				Fonte: "ReceitaWS",
				Email: empresaEmail,
			},
		}
	}

	numeros := json["telefone"].(string)
	if len(numeros) > 0 {
		arr := strings.Split(numeros, "/")
		for i := 0; i < len(arr); i++ {
			tel := telefone.New(arr[i])
			tel.Fonte = fonte.RECEITAWS
			empresa.Telefones = append(empresa.Telefones, tel)
		}
	}

	populeAtividades(json["atividade_principal"].([]interface{}), &empresa.AtividadesPrincipais)
	populeAtividades(json["atividades_secundarias"].([]interface{}), &empresa.AtividadesSecundarias)
	populeSocios(json["qsa"].([]interface{}), &empresa.Socios)

	return empresa
}

func populeAtividades(src []interface{}, dst *[]empresa.Atividade) {
	for _, value := range src {
		m := value.(map[string]interface{})
		at := empresa.Atividade{}
		at.Codigo = m["code"].(string)
		at.Texto = m["text"].(string)
		*dst = append(*dst, at)
	}
}

func populeSocios(src []interface{}, dst *[]socio.Socio) {
	for _, value := range src {
		m := value.(map[string]interface{})
		at := socio.Socio{}
		at.Qualificacao = m["qual"].(string)
		at.Nome = m["nome"].(string)
		*dst = append(*dst, at)
	}
}
