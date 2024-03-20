package services

import (
	"CSVExtractor/models"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func (database *Database) Open() error {
	var err error
	database.DB, err = sql.Open("sqlite3", "./Ailos.db")
	if err != nil {
		return err
	}
	return nil
}

func (database *Database) Close() error {
	return database.DB.Close()
}

func (database *Database) InitTables() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS CadastroBasico (
		UUID TEXT PRIMARY KEY,
		TipoPessoa TEXT,
		CPF_CNPJ TEXT,
		PEP TEXT,
		CEIS TEXT,
		CNEP TEXT,
		AutosInfracaoIbama TEXT,
		AutosInfracaoICMBIO TEXT,
		TrabalhoEscravo TEXT,
		Suspensaobama TEXT,
		ApreensaoIbama TEXT
	)`,
		`CREATE TABLE IF NOT EXISTS PEP (
            UUID TEXT PRIMARY KEY, CPF TEXT, Nome_PEP TEXT, Sigla_Funcao TEXT, Descricao_Funcao TEXT, 
            Nivel_Funcao TEXT, Nome_Orgao TEXT, Data_Inicio_Exercicio TEXT, Data_Fim_Exercicio TEXT, 
            Data_Fim_Carencia TEXT
        )`,
		`CREATE TABLE IF NOT EXISTS CNEP (
            UUID TEXT PRIMARY KEY, Cadastro TEXT, CodigoSancao TEXT, TipoPessoa TEXT, 
			CPFCNPJSanctioned TEXT, NomeSancionado TEXT, NomeInformadoOrgaoSancionador TEXT,
			RazaoSocialCadastroReceita TEXT, NomeFantasiaCadastroReceita TEXT, NumeroProcesso TEXT,
			CategoriaSancao TEXT, ValorMulta TEXT, DataInicioSancao TEXT, DataFinalSancao TEXT,
			DataPublicacao TEXT, Publicacao TEXT, Detalhamento TEXT, DataTransitoJulgado TEXT,
			AbrangenciaDecisaoJudicial TEXT, OrgaoSancionador TEXT, UfOrgaoSancionador TEXT,
			EsferaOrgaoSancionador TEXT, FundamentacaoLegal TEXT
        )`,
		`CREATE TABLE IF NOT EXISTS CEIS (
            UUID TEXT PRIMARY KEY, Cadastro TEXT, CodigoSancao TEXT, TipoPessoa TEXT, 
			CPFCNPJSanctioned TEXT, NomeSancionado TEXT, NomeInformadoOrgaoSancionador TEXT,
			RazaoSocialCadastroReceita TEXT, NomeFantasiaCadastroReceita TEXT, NumeroProcesso TEXT,
			CategoriaSancao TEXT, DataInicioSancao TEXT, DataFinalSancao TEXT,
			DataPublicacao TEXT, Publicacao TEXT, Detalhamento TEXT,DataTransitoJulgado TEXT,
			AbrangenciaDecisaoJudicial TEXT, OrgaoSancionador TEXT, UfOrgaoSancionador TEXT,
			EsferaOrgaoSancionador TEXT, FundamentacaoLegal TEXT
        )`,
		`CREATE TABLE IF NOT EXISTS AutosInfracaoIbama (
			UUID TEXT PRIMARY KEY, SeqAutoInfracao TEXT, NumAutoInfracao TEXT, SerAutoInfracao TEXT, 
			TipoAuto TEXT, TipoMulta TEXT, ValAutoInfracao TEXT, PatrimonioApuracao TEXT, 
			GravidadeInfracao TEXT, UnidArrecadacao TEXT, DesAutoInfracao TEXT, 
			DatHoraAutoInfracao TEXT, DatCienciaAutuacao TEXT, CodMunicipio TEXT, Municipio TEXT,
			Uf TEXT, NumProcesso TEXT, CodInfracao TEXT, DesInfracao TEXT, TipoInfracao TEXT, 
			NomeInfrator TEXT, CpfCnpjInfrator TEXT, DesLocalInfracao TEXT, NotificacaoVinculada TEXT, 
			AcaoFiscalizatoria TEXT, UnidControle TEXT, TipoAcao TEXT, Operacao TEXT, DatLancamento TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS AutosInfracaoICMBIO (
			UUID TEXT PRIMARY KEY, ID TEXT, NumeroAI TEXT, Serie TEXT, Origem TEXT, Tipo TEXT,			ValorMulta TEXT, Embargo TEXT, Apreensao TEXT, Autuado TEXT, 
			CPFCNPJ TEXT, DescricaoAI TEXT, DescricaoSancoes TEXT, Data TEXT, Ano TEXT, Artigo1 TEXT, 
			Artigo2 TEXT, TipoInfracao TEXT, NomeUC TEXT, CNUC TEXT, Municipio TEXT, UF TEXT, 
			TermosEmbargo TEXT, TermosApreensao TEXT, OrdemFiscalizacao TEXT, Processo TEXT, 
			Julgamento TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS TrabalhoEscravo (
			UUID TEXT PRIMARY KEY, ID TEXT, AnoAcaoFiscal TEXT, UF TEXT, Empregador TEXT, 
			CNPJCPF TEXT, Estabelecimento TEXT, TrabalhadoresEnvolvidos TEXT, CNAE TEXT, 
			DecisaoAdministrativa TEXT, InclusaoCadastroEmpregadores TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS Suspensaobama (
			UUID TEXT PRIMARY KEY, SEQ_TAD TEXT, STATUS_FORMULARIO TEXT, SIT_CANCELADO TEXT, 
			NUM_TAD TEXT, SER_TAD TEXT, DAT_TAD TEXT, DAT_IMPRESSAO TEXT, NUM_PESSOA_SUSPENSAO TEXT, 
			NOM_PESSOA_SUSPENSAO TEXT, CPF_CNPJ_PESSOA_SUSPENSAO TEXT, NUM_PROCESSO TEXT, DES_TAD TEXT, 
			COD_MUNICIPIO TEXT, NOM_MUNICIPIO TEXT, SIG_UF TEXT, DES_LOCALIZACAO TEXT, 
			DES_JUSTIFICATIVA TEXT, FORMA_ENTREGA TEXT, UNID_APRESENTACAO TEXT, UNID_CONTROLE TEXT, 
			SEQ_AUTO_INFRACAO TEXT, SEQ_NOTIFICACAO TEXT, SEQ_ACAO_FISCALIZATORIA TEXT, 
			SEQ_ORDEM_FISCALIZACAO TEXT, NUM_ORDEM_FISCALIZACAO TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS ApreensaoIbama (
			UUID TEXT PRIMARY KEY, SEQ_TAD TEXT, STATUS_FORMULARIO TEXT, SIT_CANCELADO TEXT, 
			NUM_TAD TEXT, SER_TAD TEXT, DAT_TAD TEXT, DAT_IMPRESSAO TEXT, NUM_PESSOA_SUSPENSAO TEXT, 
			NOM_PESSOA_SUSPENSAO TEXT, CPF_CNPJ_PESSOA_SUSPENSAO TEXT, NUM_PROCESSO TEXT, DES_TAD TEXT, 
			COD_MUNICIPIO TEXT, NOM_MUNICIPIO TEXT, SIG_UF TEXT, DES_LOCALIZACAO TEXT, 
			DES_JUSTIFICATIVA TEXT, FORMA_ENTREGA TEXT, UNID_APRESENTACAO TEXT, UNID_CONTROLE TEXT, 
			SEQ_AUTO_INFRACAO TEXT, SEQ_NOTIFICACAO TEXT, SEQ_ACAO_FISCALIZATORIA TEXT, 
			SEQ_ORDEM_FISCALIZACAO TEXT, NUM_ORDEM_FISCALIZACAO TEXT
		)`,
	}

	for _, query := range queries {
		_, err := database.DB.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil

}

func (database *Database) InsertCadastroBasico(person models.CadastroBasico) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO CadastroBasico (
			UUID, TipoPessoa, CPF_CNPJ, PEP, CEIS, CNEP, AutosInfracaoIbama, 
			AutosInfracaoICMBIO, TrabalhoEscravo, Suspensaobama, ApreensaoIbama
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		person.UUID,
		person.TipoPessoa,
		person.CPF_CNPJ,
		person.PEP,
		person.CEIS,
		person.CNEP,
		person.AutosInfracaoIbama,
		person.AutosInfracaoICMBIO,
		person.TrabalhoEscravo,
		person.Suspensaobama,
		person.ApreensaoIbama,
	)
	return err
}

func (database *Database) InsertPEP(pep models.PEP) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO PEP (UUID, CPF, Nome_PEP, Sigla_Funcao, Descricao_Funcao, Nivel_Funcao, Nome_Orgao, 
			Data_Inicio_Exercicio, Data_Fim_Exercicio, Data_Fim_Carencia) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(pep.UUID, pep.CPF, pep.Nome_PEP, pep.Sigla_Funcao, pep.Descricao_Funcao, pep.Nivel_Funcao,
		pep.Nome_Orgao, pep.Data_Inicio_Exercicio, pep.Data_Fim_Exercicio, pep.Data_Fim_Carencia)
	return err
}

func (database *Database) InsertCNEP(cnep models.CNEP) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO CNEP (UUID, Cadastro, CodigoSancao, TipoPessoa, 
			CPFCNPJSanctioned, NomeSancionado, NomeInformadoOrgaoSancionador,
			RazaoSocialCadastroReceita, NomeFantasiaCadastroReceita, NumeroProcesso,
			CategoriaSancao, ValorMulta, DataInicioSancao, DataFinalSancao,
			DataPublicacao, Publicacao, Detalhamento, DataTransitoJulgado,
			AbrangenciaDecisaoJudicial, OrgaoSancionador, UfOrgaoSancionador,
			EsferaOrgaoSancionador, FundamentacaoLegal) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(cnep.UUID, cnep.Cadastro, cnep.CodigoSancao, cnep.TipoPessoa,
		cnep.CPFCNPJSancionado, cnep.NomeSancionado, cnep.NomeInformadoOrgaoSancionador,
		cnep.RazaoSocialCadastroReceita, cnep.NomeFantasiaCadastroReceita, cnep.NumeroProcesso,
		cnep.CategoriaSancao, cnep.ValorMulta, cnep.DataInicioSancao, cnep.DataFinalSancao,
		cnep.DataPublicacao, cnep.Publicacao, cnep.Detalhamento, cnep.DataTransitoJulgado,
		cnep.AbrangenciaDecisaoJudicial, cnep.OrgaoSancionador, cnep.UfOrgaoSancionador,
		cnep.EsferaOrgaoSancionador, cnep.FundamentacaoLegal)

	return err
}

func (database *Database) InsertCEIS(ceis models.CEIS) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO CEIS (UUID, Cadastro, CodigoSancao, TipoPessoa, 
			CPFCNPJSanctioned, NomeSancionado, NomeInformadoOrgaoSancionador,
			RazaoSocialCadastroReceita, NomeFantasiaCadastroReceita, NumeroProcesso,
			CategoriaSancao, DataInicioSancao, DataFinalSancao,
			DataPublicacao, Publicacao, Detalhamento, DataTransitoJulgado,
			AbrangenciaDecisaoJudicial, OrgaoSancionador, UfOrgaoSancionador,
			EsferaOrgaoSancionador, FundamentacaoLegal)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ceis.UUID, ceis.Cadastro, ceis.CodigoSancao, ceis.TipoPessoa,
		ceis.CPFCNPJSancionado, ceis.NomeSancionado, ceis.NomeInformadoOrgaoSancionador,
		ceis.RazaoSocialCadastroReceita, ceis.NomeFantasiaCadastroReceita, ceis.NumeroProcesso,
		ceis.CategoriaSancao, ceis.DataInicioSancao, ceis.DataFinalSancao,
		ceis.DataPublicacao, ceis.Publicacao, ceis.Detalhamento, ceis.DataTransitoJulgado,
		ceis.AbrangenciaDecisaoJudicial, ceis.OrgaoSancionador, ceis.UfOrgaoSancionador,
		ceis.EsferaOrgaoSancionador, ceis.FundamentacaoLegal)

	return err
}

func (database *Database) InsertAutosInfracaoIbama(auto models.AutosInfracaoIbama) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO AutosInfracaoIbama (
			UUID, SeqAutoInfracao, NumAutoInfracao, SerAutoInfracao, TipoAuto, TipoMulta, 
			ValAutoInfracao, PatrimonioApuracao, GravidadeInfracao, UnidArrecadacao,
			DesAutoInfracao, DatHoraAutoInfracao, DatCienciaAutuacao, CodMunicipio, Municipio,
			Uf, NumProcesso, CodInfracao, DesInfracao, TipoInfracao, NomeInfrator, 
			CpfCnpjInfrator, DesLocalInfracao, NotificacaoVinculada, AcaoFiscalizatoria, 
			UnidControle, TipoAcao, Operacao, DatLancamento 
		) 
		VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(auto.UUID, auto.SeqAutoInfracao, auto.NumAutoInfracao, auto.SerAutoInfracao,
		auto.TipoAuto, auto.TipoMulta, auto.ValAutoInfracao, auto.PatrimonioApuracao,
		auto.GravidadeInfracao, auto.UnidArrecadacao, auto.DesAutoInfracao,
		auto.DatHoraAutoInfracao, auto.DatCienciaAutuacao, auto.CodMunicipio,
		auto.Municipio, auto.Uf, auto.NumProcesso, auto.CodInfracao,
		auto.DesInfracao, auto.TipoInfracao, auto.NomeInfrator,
		auto.CpfCnpjInfrator, auto.DesLocalInfracao, auto.NotificacaoVinculada,
		auto.AcaoFiscalizatoria, auto.UnidControle, auto.TipoAcao,
		auto.Operacao, auto.DatLancamento)

	return err
}

func (database *Database) InsertAutosInfracaoICMBIO(auto models.AutosInfracaoICMBIO) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO AutosInfracaoICMBIO (
			UUID, ID, NumeroAI, Serie, Origem, Tipo, ValorMulta, Embargo, Apreensao, Autuado, 
			CPFCNPJ, DescricaoAI, DescricaoSancoes, Data, Ano, Artigo1, Artigo2, TipoInfracao,
			NomeUC, CNUC, Municipio, UF, TermosEmbargo, TermosApreensao, OrdemFiscalizacao,
			Processo, Julgamento
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		auto.UUID, auto.ID, auto.NumeroAI, auto.Serie, auto.Origem, auto.Tipo,
		auto.ValorMulta, auto.Embargo, auto.Apreensao, auto.Autuado, auto.CPFCNPJ,
		auto.DescricaoAI, auto.DescricaoSancoes, auto.Data, auto.Ano, auto.Artigo1, auto.Artigo2,
		auto.TipoInfracao, auto.NomeUC, auto.CNUC, auto.Municipio, auto.UF,
		auto.TermosEmbargo, auto.TermosApreensao, auto.OrdemFiscalizacao, auto.Processo, auto.Julgamento,
	)

	return err
}

func (database *Database) InsertTrabalhoEscravo(trabalho models.TrabalhoEscravo) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO TrabalhoEscravo (
			UUID, ID, AnoAcaoFiscal, UF, Empregador, CNPJCPF, Estabelecimento, TrabalhadoresEnvolvidos,
            CNAE, DecisaoAdministrativa, InclusaoCadastroEmpregadores
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		trabalho.UUID, trabalho.ID, trabalho.AnoAcaoFiscal, trabalho.UF, trabalho.Empregador,
		trabalho.CNPJCPF, trabalho.Estabelecimento, trabalho.TrabalhadoresEnvolvidos,
		trabalho.CNAE, trabalho.DecisaoAdministrativa, trabalho.InclusaoCadastroEmpregadores,
	)

	return err
}

func (database *Database) InsertSuspensaobama(suspensao models.Suspensaobama) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO Suspensaobama (UUID, SEQ_TAD, STATUS_FORMULARIO, SIT_CANCELADO, NUM_TAD,
			SER_TAD, DAT_TAD, DAT_IMPRESSAO, NUM_PESSOA_SUSPENSAO, NOM_PESSOA_SUSPENSAO,
			CPF_CNPJ_PESSOA_SUSPENSAO, NUM_PROCESSO, DES_TAD, COD_MUNICIPIO, NOM_MUNICIPIO,
			SIG_UF, DES_LOCALIZACAO, DES_JUSTIFICATIVA, FORMA_ENTREGA, UNID_APRESENTACAO,
			UNID_CONTROLE, SEQ_AUTO_INFRACAO, SEQ_NOTIFICACAO, SEQ_ACAO_FISCALIZATORIA,
			SEQ_ORDEM_FISCALIZACAO, NUM_ORDEM_FISCALIZACAO) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		suspensao.UUID, suspensao.SEQ_TAD, suspensao.STATUS_FORMULARIO, suspensao.SIT_CANCELADO, suspensao.NUM_TAD,
		suspensao.SER_TAD, suspensao.DAT_TAD, suspensao.DAT_IMPRESSAO, suspensao.NUM_PESSOA_SUSPENSAO, suspensao.NOM_PESSOA_SUSPENSAO,
		suspensao.CPF_CNPJ_PESSOA_SUSPENSAO, suspensao.NUM_PROCESSO, suspensao.DES_TAD, suspensao.COD_MUNICIPIO, suspensao.NOM_MUNICIPIO,
		suspensao.SIG_UF, suspensao.DES_LOCALIZACAO, suspensao.DES_JUSTIFICATIVA, suspensao.FORMA_ENTREGA, suspensao.UNID_APRESENTACAO,
		suspensao.UNID_CONTROLE, suspensao.SEQ_AUTO_INFRACAO, suspensao.SEQ_NOTIFICACAO, suspensao.SEQ_ACAO_FISCALIZATORIA,
		suspensao.SEQ_ORDEM_FISCALIZACAO, suspensao.NUM_ORDEM_FISCALIZACAO,
	)

	return err
}

func (database *Database) InsertApreensaoIbama(apreensao models.ApreensaoIbama) error {
	statement, err := database.DB.Prepare(`
		INSERT INTO ApreensaoIbama (
			UUID, SEQ_TAD, STATUS_FORMULARIO, SIT_CANCELADO, NUM_TAD,
			SER_TAD, DAT_TAD, DAT_IMPRESSAO, NUM_PESSOA_SUSPENSAO, NOM_PESSOA_SUSPENSAO,
			CPF_CNPJ_PESSOA_SUSPENSAO, NUM_PROCESSO, DES_TAD, COD_MUNICIPIO, NOM_MUNICIPIO,
			SIG_UF, DES_LOCALIZACAO, DES_JUSTIFICATIVA, FORMA_ENTREGA, UNID_APRESENTACAO,
			UNID_CONTROLE, SEQ_AUTO_INFRACAO, SEQ_NOTIFICACAO, SEQ_ACAO_FISCALIZATORIA,
			SEQ_ORDEM_FISCALIZACAO, NUM_ORDEM_FISCALIZACAO
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(
		apreensao.UUID, apreensao.SEQ_TAD, apreensao.STATUS_FORMULARIO, apreensao.SIT_CANCELADO, apreensao.NUM_TAD,
		apreensao.SER_TAD, apreensao.DAT_TAD, apreensao.DAT_IMPRESSAO, apreensao.NUM_PESSOA_SUSPENSAO, apreensao.NOM_PESSOA_SUSPENSAO,
		apreensao.CPF_CNPJ_PESSOA_SUSPENSAO, apreensao.NUM_PROCESSO, apreensao.DES_TAD, apreensao.COD_MUNICIPIO, apreensao.NOM_MUNICIPIO,
		apreensao.SIG_UF, apreensao.DES_LOCALIZACAO, apreensao.DES_JUSTIFICATIVA, apreensao.FORMA_ENTREGA, apreensao.UNID_APRESENTACAO,
		apreensao.UNID_CONTROLE, apreensao.SEQ_AUTO_INFRACAO, apreensao.SEQ_NOTIFICACAO, apreensao.SEQ_ACAO_FISCALIZATORIA,
		apreensao.SEQ_ORDEM_FISCALIZACAO, apreensao.NUM_ORDEM_FISCALIZACAO,
	)

	return err
}
