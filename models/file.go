package models

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/antoniny/go_lang_text/utils"
)

//"Layout do arquivo text/csv"
type FileLayout struct {
	Cpf              string
	Private          string
	Incomplete       string
	LastPurchase     string
	TicketAvgAmount  string
	TicketLastAmount string
	CnpjMaxFrequency string
	CnpjLastPurchase string
	CpfOK            string
	CnpjOK           string
}

var OneRecord FileLayout
var AllRecords []FileLayout

func LerArquivos(patchFile string) int32 {
	var countLines int32 = 0
	arquivo, err := os.Open(patchFile)

	if err != nil {
		log.Panic("[file] " + err.Error())
	}
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string

	scanner := bufio.NewScanner(arquivo)

	//Percorrendo arquivo texto
	for scanner.Scan() {
		//ignorando primeira linha do arquivo
		if countLines == 0 {
			countLines++
			continue
		}

		countLines++

		//substituindo os espaços por ponto e virgula
		//Na sequencia Split do arquivo por ponto e virgula
		linha := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ";"), ";")
		linhas = nil
		//Gerar nova linha dos campos ignorando os espaços vazios
		for i := range linha {
			if len(linha[i]) > 0 {
				linhas = append(linhas, strings.TrimSpace(linha[i]))
			}
		}
		//Tratar os campos e cria matriz dos registros
		OneRecord.Cpf = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(linhas[0], "-", ""), "/", ""), ".", "")
		OneRecord.Private = linhas[1]
		OneRecord.Incomplete = linhas[2]
		OneRecord.LastPurchase = strings.ReplaceAll(linhas[3], "NULL", "null")
		OneRecord.TicketAvgAmount = strings.ReplaceAll(linhas[4], ",", ".")
		OneRecord.TicketLastAmount = strings.ReplaceAll(linhas[5], ",", ".")
		OneRecord.CnpjLastPurchase = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(linhas[6], "-", ""), "/", ""), ".", "")
		OneRecord.CnpjMaxFrequency = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(linhas[7], "-", ""), "/", ""), ".", "")
		OneRecord.CpfOK = utils.ValidarCPF(OneRecord.Cpf)
		if (utils.ValidarCNPJ(OneRecord.CnpjLastPurchase) == "1") && (utils.ValidarCNPJ(OneRecord.CnpjMaxFrequency) == "1") {
			OneRecord.CnpjOK = "1"
		} else {
			OneRecord.CnpjOK = "0"
		}
		AllRecords = append(AllRecords, OneRecord)
	}
	return countLines
}
