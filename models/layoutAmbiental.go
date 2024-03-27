package models

type RESULTAMBIENTAL struct {
	CODCONTROLE string
	DATASET     string
	TIPO        string
	CPFCNPJ     string
	NOME        string
	Data        string
	Municipio   string
	UF          string
	NumProcesso string
	Descricao   string
	NomeAux     string
}

type AutosInfracaoIbama struct {
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

type SuspensaoIbama struct {
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

type AreaEmbargadaIbama struct {
	SEQ_TAD                 string
	NUM_TAD                 string
	SER_TAD                 string
	DAT_EMBARGO             string
	DAT_IMPRESSAO           string
	DAT_ULT_ALTERACAO       string
	NOME_PESSOA_EMBARGADA   string
	CPF_CNPJ_EMBARGADO      string
	NUM_PROCESSO            string
	DES_TAD                 string
	COD_MUNICIPIO           string
	MUNICIPIO               string
	COD_UF                  string
	UF                      string
	DES_LOCALIZACAO         string
	NUM_LONGITUDE_TAD       string
	NUM_LATITUDE_TAD        string
	NUM_LONGITUDE_GMS_TAD   string
	NUM_LATITUDE_GMS_TAD    string
	QTD_AREA_EMBARGADA      string
	NOME_IMOVEL             string
	SIT_DESMATAMENTO        string
	WKT_GEOM_AREA_EMBARGADA string
	DAT_ULT_ALTER_GEOM      string
	NUM_AUTO_INFRACAO       string
	SER_AUTO_INFRACAO       string
	QTD_AREA_DESMATADA      string
	DES_INFRACAO            string
	COD_TIPO_BIOMA          string
	DES_TIPO_BIOMA          string
	ACAO_FISCALIZATORIA     string
	ORDEM_FISCALIZACAO      string
	OPERACAO                string
}

type InfracaoIbama struct {
	Numero           string
	TipoInfracao     string
	DataInfracao     string
	Bioma            string
	Estado           string
	Municipio        string
	CPFouCNPJ        string
	NomeAutuado      string
	NumeroAI         string
	SerieAI          string
	ValorMulta       string
	NumeroProcesso   string
	StatusDebito     string
	SancoesAplicadas string
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

type EmbargoICMBIO struct {
	ID                 string `csv:"ID"`
	NumeroEmbargo      string `csv:"Número Embargo"`
	Serie              string `csv:"Série"`
	Origem             string `csv:"Origem"`
	NumeroAutoInfracao string `csv:"Número Auto Infração"`
	CPFCNPJ            string `csv:"CPF/CNPJ"`
	Autuado            string `csv:"Autuado"`
	DescricaoInfracao  string `csv:"Descrição da Infração"`
	DescricaoSancao    string `csv:"Descrição da Sanção"`
	Art1               string `csv:"Art 1 (Dec n° 6.514/08)"`
	Art2               string `csv:"Art 2 (Dec n° 6.514/08)"`
	TipoInfracao       string `csv:"Tipo de Infração"`
	NomeUC             string `csv:"Nome UC"`
	CNUC               string `csv:"CNUC"`
	Municipio          string `csv:"Município"`
	UF                 string `csv:"UF"`
	DataAuto           string `csv:"Data do Auto"`
	Area               string `csv:"area"`
	NumProcesso        string `csv:"N° do Processo"`
}

type Aneel struct {
	DatGeracaoConjuntoDados     string `csv:"DatGeracaoConjuntoDados"`
	SigAgenteFiscalizador       string `csv:"SigAgenteFiscalizador"`
	NumAutoInfracao             string `csv:"NumAutoInfracao"`
	DatLavraturaAutoInfracao    string `csv:"DatLavraturaAutoInfracao"`
	NomNaturezaFiscalizacao     string `csv:"NomNaturezaFiscalizacao"`
	DscObjetoFiscalizado        string `csv:"DscObjetoFiscalizado"`
	CodObjetoFiscalizado        string `csv:"CodObjetoFiscalizado"`
	NomAgenteFiscalizado        string `csv:"NomAgenteFiscalizado"`
	NumCPFCNPJAgenteFiscalizado string `csv:"NumCPFCNPJAgenteFiscalizado"`
	NumProcessoPunitivo         string `csv:"NumProcessoPunitivo"`
	NumProcessoPunitivoANEEL    string `csv:"NumProcessoPunitivoANEEL"`
	DatRecebimentoAutoInfracao  string `csv:"DatRecebimentoAutoInfracao"`
	DscTipoPenalidade           string `csv:"DscTipoPenalidade"`
	VlrPenalidade               string `csv:"VlrPenalidade"`
	DtRecebimentoRecurso        string `csv:"DtRecebimentoRecurso"`
	DatDecisaoJuizo             string `csv:"DatDecisaoJuizo"`
	DscDecisaoCompletaJuizo     string `csv:"DscDecisaoCompletaJuizo"`
	DscAtoJuizo                 string `csv:"DscAtoJuizo"`
	VlrMultaAposJuizo           string `csv:"VlrMultaAposJuizo"`
	DatDecisaoDiretoria         string `csv:"DatDecisaoDiretoria"`
	DscDecisaoCompletaDiretoria string `csv:"DscDecisaoCompletaDiretoria"`
	DscAtoDiretoria             string `csv:"DscAtoDiretoria"`
	VlrMultaAposDiretoria       string `csv:"VlrMultaAposDiretoria"`
	NumTermoEncerramento        string `csv:"NumTermoEncerramento"`
	DatLavraturaTE              string `csv:"DatLavraturaTE"`
	DscEnquadramentoAI          string `csv:"DscEnquadramentoAI"`
	NumTermoNotificacao         string `csv:"NumTermoNotificacao"`
	NumProcessoFiscalizacao     string `csv:"NumProcessoFiscalizacao"`
}
