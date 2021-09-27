# Go code
### Exemplos de códigos desenvolvidos na linguagem Go.

Este repositório contém pacotes desenvolvido em golang para várias aplicações. O objetivo principal é aperfeiçoar o conhecimento na linguagem e testar os pacotes padrões e de terceiros.

### Pré-requisitos

Instalar a linguagem [Go](https://golang.org/dl/) e seguir os [passos](https://golang.org/doc/install) para adicionar o binário ao PATH do sistema.

Pode-se obter este repositório pelo utilitário go get:
```
$ go get -v github.com/paulocsilvajr/go-code
```
Após o download, este repositório estará disponível em ~/go/src/github.com/paulocsilvajr/go-code.

Para acessar a docstring, execute:
```
$ go doc github.com/paulocsilvajr/go-code/util  # lista funções de pacote util
$ go doc github.com/paulocsilvajr/go-code/util Dir # docstring de função Dir
```

Execute o teste unitário pelo comando
```
$ go test -cover github.com/paulocsilvajr/go-code/util  # teste para o pacote util
$ ./test.sh       # script de teste
```

### Arquivos

```
list: Pacote que implementa lista aos moldes do tipo list da linguagem Python.
list_test.go: Arquivo do teste unitário do pacote list.
server_restful_json: Teste de servidor REST JSON baseado em https://thenewstack.io/make-a-restful-json-api-go/
```

### Licença

[Licença GPL](https://github.com/paulocsilvajr/go-code/blob/master/license_gpl.txt), arquivo em anexo no repositório.

### Contato

Paulo Carvalho da Silva Junior - pauluscave@gmail.com
