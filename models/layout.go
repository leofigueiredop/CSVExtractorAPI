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

type AutosInfracaoIbama struct {
	UUID                string
	SeqAutoInfracao     string `csv:"SEQ_AUTO_INFRACAO"`
	NumAutoInfracao     string `csv:"NUM_AUTO_INFRACAO"`
	SerAutoInfracao     string `csv:"SER_AUTO_INFRACAO"`
	TipoAuto            string `csv:"TIPO_AUTO"`
	TipoMulta           string `csv:"TIPO_MULTA"`
	ValAutoInfracao     string `csv:"VAL_AUTO_INFRACAO"`
	PatrimonioApuracao  string `csv:"PATRIMONIO_APURACAO"`
	GravidadeInfracao   string `csv:"GRAVIDADE_INFRACAO"`
	UnidArrecadacao     string `csv:"UNID_ARRECADACAO"`
	DesAutoInfracao     string `csv:"DES_AUTO_INFRACAO"`
	DatHoraAutoInfracao string `csv:"DAT_HORA_AUTO_INFRACAO"`
	DatCienciaAutuacao  string `csv:"DAT_CIENCIA_AUTUACAO"`
	Municipio           string `csv:"MUNICIPIO"`
	Uf                  string `csv:"UF"`
	NumProcesso         string `csv:"NUM_PROCESSO"`
	CodInfracao         string `csv:"COD_INFRACAO"`
	DesInfracao         string `csv:"DES_INFRACAO"`
	TipoInfracao        string `csv:"TIPO_INFRACAO"`
	NomeInfrator        string `csv:"NOME_INFRATOR"`
	CpfCnpjInfrator     string `csv:"CPF_CNPJ_INFRATOR"`
	DesLocalInfracao    string `csv:"DES_LOCAL_INFRACAO"`
	TipoAcao            string `csv:"TIPO_ACAO"`
	Operacao            string `csv:"OPERACAO"`
	DatLancamento       string `csv:"DAT_LANCAMENTO"`
}

type AutosInfracaoICMBIO struct {
	UUID             string
	ID               string `csv:"ID"`
	NumeroAI         string `csv:"Número AI"`
	Serie            string `csv:"Série"`
	Origem           string `csv:"Origem"`
	Tipo             string `csv:"Tipo"`
	ValorMulta       string `csv:"Valor da Multa"`
	Embargo          string `csv:"Embargo"`
	Apreensao        string `csv:"Apreensão"`
	Autuado          string `csv:"Autuado"`
	CPFCNPJ          string `csv:"CPF/CNPJ"`
	DescricaoAI      string `csv:"Descrição AI"`
	DescricaoSancoes string `csv:"Descrição das Sanções"`
	Data             string `csv:"Data"`
	Ano              string `csv:"Ano"`
	Artigo1          string `csv:"Artigo 1"`
	Artigo2          string `csv:"Artigo 2"`
	TipoInfracao     string `csv:"Tipo de Infração"`
	NomeUC           string `csv:"Nome da UC"`
	CNUC             string `csv:"CNUC"`
	Municipio        string `csv:"Município"`
	UF               string `csv:"UF"`
	TermosEmbargo    string `csv:"Termos Embargo"`
	TermosApreensao  string `csv:"Termos Apreensão"`
	Processo         string `csv:"Processo"`
	Julgamento       string `csv:"Julgamento"`
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

type Suspensaobama struct {
	UUID                      string
	SEQ_TAD                   string `csv:"SEQ_TAD"`
	STATUS_FORMULARIO         string `csv:"STATUS_FORMULARIO"`
	SIT_CANCELADO             string `csv:"SIT_CANCELADO"`
	NUM_TAD                   string `csv:"NUM_TAD"`
	SER_TAD                   string `csv:"SER_TAD"`
	DAT_TAD                   string `csv:"DAT_TAD"`
	DAT_IMPRESSAO             string `csv:"DAT_IMPRESSAO"`
	NUM_PESSOA_SUSPENSAO      string `csv:"NUM_PESSOA_SUSPENSAO"`
	NOM_PESSOA_SUSPENSAO      string `csv:"NOM_PESSOA_SUSPENSAO"`
	CPF_CNPJ_PESSOA_SUSPENSAO string `csv:"CPF_CNPJ_PESSOA_SUSPENSAO"`
	NUM_PROCESSO              string `csv:"NUM_PROCESSO"`
	DES_TAD                   string `csv:"DES_TAD"`
	NOM_MUNICIPIO             string `csv:"NOM_MUNICIPIO"`
	SIG_UF                    string `csv:"SIG_UF"`
	DES_LOCALIZACAO           string `csv:"DES_LOCALIZACAO"`
	DES_JUSTIFICATIVA         string `csv:"DES_JUSTIFICATIVA"`
	UNID_CONTROLE             string `csv:"UNID_CONTROLE"`
	SEQ_AUTO_INFRACAO         string `csv:"SEQ_AUTO_INFRACAO"`
}

type ApreensaoIbama struct {
	UUID                      string
	SEQ_TAD                   string `csv:"SEQ_TAD"`
	STATUS_FORMULARIO         string `csv:"STATUS_FORMULARIO"`
	SIT_CANCELADO             string `csv:"SIT_CANCELADO"`
	NUM_TAD                   string `csv:"NUM_TAD"`
	SER_TAD                   string `csv:"SER_TAD"`
	DAT_TAD                   string `csv:"DAT_TAD"`
	DAT_IMPRESSAO             string `csv:"DAT_IMPRESSAO"`
	NUM_PESSOA_SUSPENSAO      string `csv:"NUM_PESSOA_SUSPENSAO"`
	NOM_PESSOA_SUSPENSAO      string `csv:"NOM_PESSOA_SUSPENSAO"`
	CPF_CNPJ_PESSOA_SUSPENSAO string `csv:"CPF_CNPJ_PESSOA_SUSPENSAO"`
	NUM_PROCESSO              string `csv:"NUM_PROCESSO"`
	DES_TAD                   string `csv:"DES_TAD"`
	NOM_MUNICIPIO             string `csv:"NOM_MUNICIPIO"`
	SIG_UF                    string `csv:"SIG_UF"`
	DES_LOCALIZACAO           string `csv:"DES_LOCALIZACAO"`
	DES_JUSTIFICATIVA         string `csv:"DES_JUSTIFICATIVA"`
	SEQ_AUTO_INFRACAO         string `csv:"SEQ_AUTO_INFRACAO"`
	SEQ_NOTIFICACAO           string `csv:"SEQ_NOTIFICACAO"`
}
