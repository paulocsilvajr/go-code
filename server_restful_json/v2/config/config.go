package config

import (
	"encoding/gob"
	"os"
)

// https://play.golang.org/p/6dX5SMdVtr

func criarConfigPadrao() {
	configuracoes := make(map[string]string)
	configuracoes["porta"] = "8080"
	configuracoes["host"] = "localhost"

	encodeFile, err := os.Create("config/config.gob")
	if err != nil {
		panic(err)
	}

	encoder := gob.NewEncoder(encodeFile)

	if err := encoder.Encode(configuracoes); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

func AbrirConfiguracoes() map[string]string {
	decodeFile, err := os.Open("config/config.gob")
	if err != nil {
		criarConfigPadrao()

		decodeFile, err = os.Open("config/config.gob")
		if err != nil {
			panic(err)
		}
	}
	defer decodeFile.Close()

	decoder := gob.NewDecoder(decodeFile)

	configuracoes := make(map[string]string)

	decoder.Decode(&configuracoes)

	return configuracoes
}
