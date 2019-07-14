/**
* Author: Anderson Antoniny
* Desafio Neoway
* Data: 2019-03-27
**/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/antoniny/go_lang_text/config"
	"github.com/antoniny/go_lang_text/models"
	"github.com/antoniny/go_lang_text/utils"
)

//Função Principal
func main() {
	/*****************************************************************
	*  Validando argumentos
	*****************************************************************/
	log.Println("[main] begin")

	if len(os.Args) != 2 {
		log.Println("[main] Arquivo não encontrado. [text/csv]. Favor informar path do arquivo. ex.: csv.exe c:\\arquivo.txt ")
		//	os.Exit(1)
	}
	arg := os.Args[1]
	if !utils.FileExists(arg) {
		log.Println("[main] Arquivo não encontrado. [text/csv] ex.: csv.exe arquivo.txt ")
		//	os.Exit(1)
	}
	log.Println("[main] *")
	log.Println("[main] * Author: Anderson Antoniny")
	log.Println("[main] * Nota: Os logs apresentados na execução deste programa são apenas de nível didático para este desafio.")
	log.Println("[main] *")
	arg = "file_input.txt"
	/*****************************************************************
	*  Manipulando arquivo texto
	*****************************************************************/
	log.Println("[main] File:", arg)
	countLines := models.LerArquivos(arg)
	log.Println("[main] Linhas encontradas: ", countLines)

	/*****************************************************************
	*  Persistindo linhas na base de dados
	*****************************************************************/
	//Inicializa config
	//carrega as variaveis do config.yam
	config.Init()

	//Conectar na base de dados.
	db := models.Init()

	//Loop nos registros capturados do arquivo
	//Inserindo na base de dados
	insereRegistrosFile(db)

	//Atualiza a tabela - Tratamento de nulos
	rowsAffected := models.UpdateFiledsNullable(db)
	log.Println("[main]", rowsAffected, " Registros atualizado para [null]")

	//Busca status da tabela
	qtySQL, qtyCPF, qtyCNPJ := models.BuscaStatusTable(db)
	log.Println("[main] ********************************************************************")
	log.Println("[main] " + fmt.Sprintf("* Total de linhas: %d CPF's inválidos: %d CNPJ's inválidos: %d *", qtySQL, qtyCPF, qtyCNPJ))
	log.Println("[main] ********************************************************************")
	log.Println("[main] end")

}

//Função loop de registros File para gravar na base
func insereRegistrosFile(db *sql.DB) {
	id := 0
	countLines := 0
	for _, data := range models.AllRecords {
		//Chamada da função DB Insert
		id = models.InsertLineFileCSV(db, data)
		countLines++
		_, resto := utils.DivMod(int64(countLines), 5000)
		if resto == 0 {
			log.Println("[main] ", countLines, " registros inseridos com sucesso! LastID: ", id)
		}
	}
}
