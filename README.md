# Arquitetura Hexagonal com Mongo e MySQL

## Pré-requisitos
- Golang
- Docker
- Postman (ou Curl)

## Rodando

### Mongo DB

Configure as variáveis de ambiente:
````
$ export MONGO_CONN: mongodb://localhost:27017 (se usar docker, senão a url de conexão que desejar)
$ export MONGO_DB: hexagonal_arch (ou a que desejar)
````

Inicie uma image docker:
````
$ docker run --name mongo -p 27017:27017 mongo
````

### MySQL

Configure as variáveis de ambiente:
````
$ export MYSQL_USR=root
$ export MYSQL_PASS=root
$ export MYSQL_HOST=localhost
$ export MYSQL_PORT=3306
$ export MYSQL_DBNAME=hexagonal_arch
````

Inicie uma image docker:
````
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql
````

Conecte no banco MySQL:
````
$ docker exec -it mysql /bin/bash

## dentro da imagem
# mysql -u root -p (coloque a senha `root` ou outra que tenha definido)

## dentro do banco
mysql> create database hexagonal_arch;
````

Agora, volte para a pasta do projeto, e execute o comando seguinte, para rodar as migrações no banco MySQL:
````
$  go run migrations/main.go 
````

## Iniciando o projeto

Após os passos anteriores, para iniciar o projeto basta apenas executar:
````
$ go run main.go
````

Existem dois parâmetros que podem ser utilizados para iniciar o serviço:
````
- repo: default para `mysql`. Outro valor aceitável é `mongo`. Define em qual base de dados se conectará ao iniciar.

- httpbind: default para `:8080`. Define em qual porta a API será exposta.
````

Exemplos de execução:

- Banco Mongo e porta 9000
````
$ go run main.go -repo=mongo -httpbind=9000
````

- Banco MySQL e porta 3333
````
$ go run main.go -httpbind=3333
````

## APIs
 
As definições das APIs podem ser encontradas [aqui](https://github.com/dexterorion/hexagonal-architecture-mongo-mysql/blob/main/Hexagonal%20Arch.postman_collection.json). Para visualizá-las, utilize o [postman](https://www.postman.com/).

## Testes

Em construção...
