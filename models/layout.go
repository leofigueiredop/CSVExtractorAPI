package models

import "encoding/json"

type Result struct {
	Key   string          `json:"key"`
	Index string          `json:"index"`
	Data  json.RawMessage `json:"data"`
}

type CadastroBasico struct {
	UUID       string
	TipoPessoa string
	CPF_CNPJ   string `csv:"CPF/CNPJ"`
}

type PEP struct {
	UUID                  string
	CPF                   string `csv:"CPF"`
	Nome_PEP              string `csv:"Nome_PEP"`
	Sigla_Funcao          string `csv:"Sigla_Função"`
	Descricao_Funcao      string `csv:"Descrição_Função"`
	Nivel_Funcao          string `csv:"Nível_Função"`
	Nome_Orgao            string `csv:"Nome_Órgão"`
	Data_Inicio_Exercicio string `csv:"Data_Início_Exercício"`
	Data_Fim_Exercicio    string `csv:"Data_Fim_Exercício"`
	Data_Fim_Carencia     string `csv:"Data_Fim_Carência"`
}

type CNEP struct {
	UUID                          string
	Cadastro                      string `csv:"CADASTRO"`
	CodigoSancao                  string `csv:"CÓDIGO DA SANÇÃO"`
	CPFCNPJSancionado             string `csv:"CPF OU CNPJ DO SANCIONADO"`
	NomeSancionado                string `csv:"NOME DO SANCIONADO"`
	NomeInformadoOrgaoSancionador string `csv:"NOME INFORMADO PELO ÓRGÃO SANCIONADOR"`
	RazaoSocialCadastroReceita    string `csv:"RAZÃO SOCIAL - CADASTRO RECEITA"`
	NomeFantasiaCadastroReceita   string `csv:"NOME FANTASIA - CADASTRO RECEITA"`
	NumeroProcesso                string `csv:"NÚMERO DO PROCESSO"`
	CategoriaSancao               string `csv:"CATEGORIA DA SANÇÃO"`
	ValorMulta                    string `csv:"VALOR DA MULTA"`
	DataInicioSancao              string `csv:"DATA INÍCIO SANÇÃO"`
	DataFinalSancao               string `csv:"DATA FINAL SANÇÃO"`
	DataPublicacao                string `csv:"DATA PUBLICAÇÃO"`
	Publicacao                    string `csv:"PUBLICAÇÃO"`
	Detalhamento                  string `csv:"DETALHAMENTO"`
	DataTransitoJulgado           string `csv:"DATA DO TRÂNSITO EM JULGADO"`
	AbrangenciaDecisaoJudicial    string `csv:"ABRANGÊNCIA DEFINIDA EM DECISÃO JUDICIAL"`
	OrgaoSancionador              string `csv:"ÓRGÃO SANCIONADOR"`
	UfOrgaoSancionador            string `csv:"UF ÓRGÃO SANCIONADOR"`
	FundamentacaoLegal            string `csv:"FUNDAMENTAÇÃO LEGAL"`
}

type CEIS struct {
	UUID                          string
	Cadastro                      string `csv:"CADASTRO"`
	CodigoSancao                  string `csv:"CÓDIGO DA SANÇÃO"`
	CPFCNPJSancionado             string `csv:"CPF OU CNPJ DO SANCIONADO"`
	NomeSancionado                string `csv:"NOME DO SANCIONADO"`
	NomeInformadoOrgaoSancionador string `csv:"NOME INFORMADO PELO ÓRGÃO SANCIONADOR"`
	RazaoSocialCadastroReceita    string `csv:"RAZÃO SOCIAL - CADASTRO RECEITA"`
	NomeFantasiaCadastroReceita   string `csv:"NOME FANTASIA - CADASTRO RECEITA"`
	NumeroProcesso                string `csv:"NÚMERO DO PROCESSO"`
	CategoriaSancao               string `csv:"CATEGORIA DA SANÇÃO"`
	DataInicioSancao              string `csv:"DATA INÍCIO SANÇÃO"`
	DataFinalSancao               string `csv:"DATA FINAL SANÇÃO"`
	DataPublicacao                string `csv:"DATA PUBLICAÇÃO"`
	Publicacao                    string `csv:"PUBLICAÇÃO"`
	Detalhamento                  string `csv:"DETALHAMENTO"`
	DataTransitoJulgado           string `csv:"DATA DO TRÂNSITO EM JULGADO"`
	AbrangenciaDecisaoJudicial    string `csv:"ABRANGÊNCIA DEFINIDA EM DECISÃO JUDICIAL"`
	OrgaoSancionador              string `csv:"ÓRGÃO SANCIONADOR"`
	UfOrgaoSancionador            string `csv:"UF ÓRGÃO SANCIONADOR"`
	FundamentacaoLegal            string `csv:"FUNDAMENTAÇÃO LEGAL"`
}

type TrabalhoEscravo struct {
	UUID                         string
	ID                           string `csv:"ID"`
	AnoAcaoFiscal                string `csv:"Ano da ação fiscal"`
	UF                           string `csv:"UF"`
	Empregador                   string `csv:"Empregador"`
	CNPJCPF                      string `csv:"CNPJ/CPF"`
	Estabelecimento              string `csv:"Estabelecimento"`
	TrabalhadoresEnvolvidos      string `csv:"Trabalhadores envolvidos"`
	CNAE                         string `csv:"CNAE"`
	DecisaoAdministrativa        string `csv:"Decisão administrativa de procedência"`
	InclusaoCadastroEmpregadores string `csv:"Inclusão no Cadastro de Empregadores"`
}
