package config

import (
	"encoding/json"
	"os"
)

// https://play.golang.org/p/6dX5SMdVtr

const arquivo = "config/config.json"

func criarConfigPadrao() {
	configuracoes := make(map[string]string)
	configuracoes["porta"] = "8080"
	configuracoes["host"] = "localhost"

	encodeFile, err := os.Create(arquivo)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(encodeFile)

	if err := encoder.Encode(configuracoes); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

func AbrirConfiguracoes() map[string]string {
	decodeFile, err := os.Open(arquivo)
	if err != nil {
		criarConfigPadrao()

		decodeFile, err = os.Open(arquivo)
		if err != nil {
			panic(err)
		}
	}
	defer decodeFile.Close()

	decoder := json.NewDecoder(decodeFile)

	configuracoes := make(map[string]string)

	decoder.Decode(&configuracoes)

	return configuracoes
}
