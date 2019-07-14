
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

### Instalar e executar no Docker container:
Nota: Considerando o serviço Docker ativo na maquina. Utilizado container linux do Docker

#### Docker Linux Containes


##### Pasta Local Projeto - Windows Via Docker

Para executar o projeto via pasta do projeto no Docker, siga as instruções abaixo
Este ira gerar um docker container linux com a instalaçao Linux, o serviço do docker deverá estar como Linux containers.

##### Execução do programa localmente "com" utilização do Docker
Acesse cmd do windows e digite as linhas de comando abaixo.

##### docker-compose
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
Via prompt de comando "CMD" execute o comando abaixo dentro da pasta onde foi realizado o download.

```bash
> docker-compose up -d
```

##### build docker local
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
$GOPATH$/src/github.com/antoniny/go_csv_nw/

```bash
1 - cd %GOPATH%\src\github.com\antoniny\go_csv_nw\
2 - docker build -t docker_go_csv_nw:DEMO . -f Dockerfile
3 - docker images | grep docker_go_csv_nw
4 - docker run -it docker_go_csv_nw:DEMO
```

##### Execução do programa localmente "sem" utilização do Docker - diretamente via GoLang
Faça o download do projeto para a estrutura abaixo. (https://github.com/antoniny/go_csv_nw.git)
>$GOPATH$/src/github.com/antoniny/go_csv_nw/

 ```bash
1 - cd %GOPATH%/src/
2 - get -v github.com/antoniny/go_csv_nw/
3 - cd github.com/antoniny/go_csv_nw/
4 - main.exe file_input.txt
```

