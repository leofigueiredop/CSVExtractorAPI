package models

type Pessoa struct {
	Tipo_pessoa string    `json:"TipoPessoa"`
	CPF_CNPJ    string    `json:"CPFCNPJ"`
	Empresas    []Empresa `json:"Empresas"`
}

type Empresa struct {
	CnpjBase         string            `json:"CnpjBase"`
	RazaoSocial      string            `json:"RazaoSocial"`
	Natureza         string            `json:"Natureza"`
	Qualificacao     string            `json:"Qualificacao"`
	CapitalSocial    string            `json:"CapitalSocial"`
	Porte            string            `json:"Porte"`
	EnteFederativo   string            `json:"EnteFederativo"`
	Estabelecimentos []Estabelecimento `json:"Estabelecimentos"`
}

type Estabelecimento struct {
	CnpjBase             string `json:"CnpjBase"`
	CnpjOrdem            string `json:"CnpjOrdem"`
	CnpjDv               string `json:"CnpjDv"`
	IdentificadorMatriz  string `json:"IdentificadorMatriz"`
	NomeFantasia         string `json:"NomeFantasia"`
	SituacaoCadastral    string `json:"SituacaoCadastral"`
	DataSituacao         string `json:"DataSituacao"`
	MotivoSituacao       string `json:"MotivoSituacao"`
	CidadeExterior       string `json:"CidadeExterior"`
	Pais                 string `json:"Pais"`
	DataInicioAtividade  string `json:"DataInicioAtividade"`
	CnaeFiscalPrincipal  string `json:"CnaeFiscalPrincipal"`
	CnaeFiscalSecundaria string `json:"CnaeFiscalSecundaria"`
	TipoLogradouro       string `json:"TipoLogradouro"`
	Logradouro           string `json:"Logradouro"`
	Numero               string `json:"Numero"`
	Complemento          string `json:"Complemento"`
	Bairro               string `json:"Bairro"`
	Cep                  string `json:"Cep"`
	Uf                   string `json:"Uf"`
	Municipio            string `json:"Municipio"`
	Ddd1                 string `json:"Ddd1"`
	Telefone1            string `json:"Telefone1"`
	Ddd2                 string `json:"Ddd2"`
	Telefone2            string `json:"Telefone2"`
	DddFax               string `json:"DddFax"`
	Fax                  string `json:"Fax"`
	Email                string `json:"Email"`
	SituacaoEspecial     string `json:"SituacaoEspecial"`
	DataSituacaoEspecial string `json:"DataSituacaoEspecial"`
}
