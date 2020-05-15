package helper

import "os"

func CriarDiretorioSeNaoExistir(nomeDiretorio string) {
	if _, err := os.Stat(nomeDiretorio); os.IsNotExist(err) {
		err = os.MkdirAll(nomeDiretorio, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
