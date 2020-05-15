package logger

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/paulocsilvajr/go-code/server_restful_json/v2/helper"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		helper.CriarDiretorioSeNaoExistir("logs")

		// log em arquivo
		nomeArquivo := fmt.Sprintf("logs/%04d%02d%02d.log", start.Year(), start.Month(), start.Day())
		arquivo, err := os.OpenFile(nomeArquivo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Erro ao abrir arquivo de log[%s]", err)
		}
		defer arquivo.Close()

		// log em arquivo(arquivo) e tela(Stdout)
		multiplaSaida := io.MultiWriter(os.Stdout, arquivo)
		log.SetOutput(multiplaSaida)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}
