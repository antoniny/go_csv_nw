package models

import (
	"database/sql"
	"log"

	"github.com/antoniny/go_lang_text/config"

	_ "github.com/lib/pq"
)

//Inicializa conexão com a base de dados
func Init() (db *sql.DB) {
	// conectando a base postgres
	var err error

	for int_for := 0; int_for <= 30; int_for++ {
		db, err = sql.Open(config.Database.Driver, config.Database.Open)
		if err != nil {
			log.Println("[sqlOpen] Aguardando base de dados. (", int_for, "/30)")
		} else {
			break
		}
	}
	if err != nil {
		log.Panic("[database - sqlopen] " + err.Error())
	}

	for int_for := 0; int_for <= 30; int_for++ {
		if err = db.Ping(); err != nil {
			log.Println("[ping] Aguardando base de dados. (", int_for, "/30)")
		} else {
			db, err = sql.Open(config.Database.Driver, config.Database.Open)
			err = db.Ping()
			break
		}
	}

	if err != nil {
		log.Panic("[database - pingdb] " + err.Error())
	}

	log.Println("[database] Conectado com sucesso.")

	//Cria tabela padrão para utilizar no sistema
	log.Println("[database] Recriando Tabela [file_csv]...")
	createTable(db)
	log.Println("[database] Tabela recriada com sucesso.")

	return db
}

//Insere as linhas do arquivo na base. Ref.:FileLayout
func InsertLineFileCSV(db *sql.DB, data FileLayout) int {

	id := 0
	err := db.QueryRow("INSERT INTO file_csv (cpf,private,incomplete,last_purchase,ticket_avg_amount,ticket_last_amount,cnpj_max_frequency,cnpj_last_purchase, cpf_ok, cnpj_ok) VALUES ($1, $2, $3,  $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		data.Cpf,
		data.Private,
		data.Incomplete,
		data.LastPurchase,
		data.TicketAvgAmount,
		data.TicketLastAmount,
		data.CnpjLastPurchase,
		data.CnpjMaxFrequency,
		data.CpfOK,
		data.CnpjOK).Scan(&id)

	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	return id
}

//Corrige as colunas base sem valor para null
func UpdateFiledsNullable(db *sql.DB) int64 {
	var rowsAffected int64 = 0
	//Tratamento campo last_purchase para null no campos sem valor
	updStatement := `UPDATE file_csv SET last_purchase = NULL, last_change_time = NOW() WHERE UPPER(COALESCE(last_purchase,'NULL')) = 'NULL';`
	res, err := db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	count, err := res.RowsAffected()
	rowsAffected = rowsAffected + count
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	//Tratamento do campo ticket_avg_amount para null
	updStatement = `UPDATE file_csv SET ticket_avg_amount = NULL, last_change_time = NOW() WHERE UPPER(COALESCE(ticket_avg_amount,'NULL')) = 'NULL';`
	res, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	count, err = res.RowsAffected()
	rowsAffected = rowsAffected + count
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	//Tratamento do campo ticket_last_amount para null
	updStatement = `UPDATE file_csv SET ticket_last_amount = NULL, last_change_time = NOW() WHERE UPPER(COALESCE(ticket_last_amount,'NULL')) = 'NULL';`
	res, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	count, err = res.RowsAffected()
	rowsAffected = rowsAffected + count
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	//Tratamento do campo cnpj_max_frequency para null
	updStatement = `UPDATE file_csv SET cnpj_max_frequency = NULL, last_change_time = NOW() WHERE UPPER(COALESCE(cnpj_max_frequency,'NULL')) = 'NULL';`
	res, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	count, err = res.RowsAffected()
	rowsAffected = rowsAffected + count
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	//Tratamento do campo cnpj_last_purchase para null
	updStatement = `UPDATE file_csv SET cnpj_last_purchase = NULL, last_change_time = NOW() WHERE UPPER(COALESCE(cnpj_last_purchase,'NULL')) = 'NULL';`
	res, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	count, err = res.RowsAffected()
	rowsAffected = rowsAffected + count
	if err != nil {
		log.Panic("[database] " + err.Error())
	}

	return rowsAffected
}

//Busca status resumo da tabela
func BuscaStatusTable(db *sql.DB) (int64, int64, int64) {
	var qtySQL int64 = 0
	var qtyCPF int64 = 0
	var qtyCNPJ int64 = 0

	row, err := db.Query("SELECT count(1) as linhas, sum(TO_NUMBER(case when cpf_ok = '1' then '0' else '1' end,'9G999g999')) as cpf_invalido, sum(TO_NUMBER(case when cnpj_ok = '1' then '0' else '1' end,'9G999g999')) as cnpj_invalido from file_csv")
	if err != nil {
		log.Panic("[database] " + err.Error())
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&qtySQL, &qtyCPF, &qtyCNPJ)
		if err != nil {
			log.Panic("[database] " + err.Error())
		}
	}
	return qtySQL, qtyCPF, qtyCNPJ
}

func createTable(db *sql.DB) {

	//Tratamento do campo cnpj_last_purchase para null
	updStatement := `DROP TABLE IF EXISTS file_csv;`
	_, err := db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}

	updStatement = `CREATE TABLE file_csv (
					id                 serial    PRIMARY KEY, 
					cpf                character varying(20)              NULL, 
					private            character varying(20)              NULL, 
					incomplete         character varying(20)              NULL, 
					last_purchase      character varying(20)              NULL, 
					ticket_avg_amount  character varying(20)              NULL, 
					ticket_last_amount character varying(20)              NULL, 
					cnpj_max_frequency character varying(20)              NULL, 
					cnpj_last_purchase character varying(20)              NULL, 
						last_change_time   timestamp(6) without time zone     NULL DEFAULT now(), 
					cpf_ok             character varying(1)               NULL, 
					cnpj_ok            character varying(1)               NULL
					)
					`
	_, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}

	updStatement = `COMMENT ON COLUMN file_csv.id IS 'ID Incremental';
				COMMENT ON COLUMN file_csv.cpf IS 'CPF do Cliente';
				COMMENT ON COLUMN file_csv.private IS ' ';
				COMMENT ON COLUMN file_csv.incomplete IS ' ';
				COMMENT ON COLUMN file_csv.last_purchase IS 'Data da ultima compra';
				COMMENT ON COLUMN file_csv.ticket_avg_amount IS 'Valor do Ticket médio';
				COMMENT ON COLUMN file_csv.ticket_last_amount IS 'Valor do ultimo ticket';
				COMMENT ON COLUMN file_csv.cnpj_max_frequency IS 'Loja com maior frequencia';
				COMMENT ON COLUMN file_csv.last_change_time IS 'Ultima alteração no registro';
				`
	_, err = db.Exec(updStatement)
	if err != nil {
		log.Panic("[database] " + err.Error())
	}

}
