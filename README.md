
#App import text/csv (golang/postgres)
=
Aplicação tem como finalidade a importação de um determinado arquivo text/csv persistindo os dados em uma base de dados Postgres.

A mesma executará com container(Docker) utilizando como linguagem Golang e base de dados Postgres.

Nota: Os testes foram realizados em ambiente Windows 10 e DockerContainer linux alpine.

## Pré-requisto

- golang version (version 1.12) - https://golang.org/dl/
- Base de dados Postgres 11 - https://www.postgresql.org/download/windows
- Docker Tools -https://docs.docker.com/v17.12/toolbox/toolbox_install_windows/

## Fontes

- gitHub - https://github.com/antoniny/go_lang_text.git

## Instalação e execução pelo Windows

- 1 - Caso não tenha instalado Postgres, o mesmo deverá ser instalado.
- 2 - Considerando que a base de dados esteja acessível com usuário com privilegios.

### Configurando Postgres Database:
Execute os comandos abaixo em ```psql``` com usuário com privilegios "master" para configurar a base de dados.
Será criada uma base e usuário exclusivo para execução do programa.
```
-- Comando para criar a base de dados do app
CREATE DATABASE nw_app;
-- Comando para criar o usu?rio e senha padr?o para o app
CREATE USER user_app WITH ENCRYPTED PASSWORD 'user_app';
-- concede privilegios na base de dados para o usuário da app
GRANT ALL PRIVILEGES ON DATABASE nw_app TO user_app;
GRANT ALL PRIVILEGES ON DATABASE nw_app TO postgres;
GRANT CONNECT ON DATABASE nw_app TO user_app;```
```

### Instalar e executar no Docker container:
Nota: Considerando o serviço Docker ativo na maquina. Utilizado container linux do Docker

#### Docker Linux Containes


##### Pasta Local Projeto - Windows Via Docker

Para executar o projeto via pasta do projeto no Docker, siga as instruções abaixo
Este ira gerar um docker container linux com a instalaçao Linux, o serviço do docker deverá estar como Linux containers.

##### instalação goLang  
Realizar instalaçao do GO LANG e criar as variaveis de ambiente do windows.
>%GOPATH% -> Diretório de trabalho para o GO_LANG
>%GOROOT% -> Diretório de instalação do golang

##### Execução do programa localmente "sem" utilização do Docker - diretamente via GoLang
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
>$GOPATH$/src/github.com/antoniny/go_csv_nw/

 ```bash
1 - cd %GOPATH%/src/
2 - get -v github.com/antoniny/go_csv_nw/
3 - cd github.com/antoniny/go_csv_nw/
4 - main.exe file_input.txt
Nota: Teste realizado com compilação original
```

##### Execução do programa localmente "com" utilização do Docker
Acesse cmd do windows e digite as linhas de comando abaixo.

##### build docker local
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
$GOPATH$/src/github.com/antoniny/go_csv_nw/

```bash
1 - cd %GOPATH%\src\github.com\antoniny\go_csv_nw\
2 - docker build -t docker_go_csv_nw:DEMO . -f Dockerfile
3 - docker images | grep docker_go_csv_nw
4 - docker run -it docker_go_csv_nw:DEMO
```
##### build docker DOCKERHUB

```bash
1 - docker pull antoniny/go_csv_nw
2 - docker run -it antoniny/go_csv_nw:latest
```

##### docker-compose
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
$GOPATH$/src/github.com/antoniny/go_csv_nw/

```bash
1 - cd %GOPATH%\src\github.com\antoniny\go_csv_nw\
2 - docker-compose up -d
4 - docker-compose up
```
