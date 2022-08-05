
<div style="display: inline_block"><br>
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" />
  <img align="center" alt="Spring" height="80" width="70" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" />
  <img align="center" alt="RabbitMQ" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original-wordmark.svg" />
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" />
</div>

# Minha Terceira API usando Gin - Go Lang
Pequena API desenvolvido em Go Lang junto com a biblioteca Gin e autenticação com Jwt, que permite criar, consultar, editar e deletar um livro.

## Requisitos

- Go Lang
- Docker

## Configurando o Ambiente

### Docker
1) Instalar o Docker (Tutorial via Windows). [Download](https://www.docker.com/products/docker-desktop/) e  [Instalação](https://sh-tsang.medium.com/tutorial-docker-installation-in-wsl-2-of-windows-f4471fc3e1d4)

2) No diretório raiz, executar o comando abaixo:
```shell
docker-compose up
```

### Go Lang

1) [Baixar](https://go.dev/doc/install) e [instalar](https://medium.com/@rafaelmoraisdev/como-instalar-go-no-windows-10-7787faac3a7f) o Go Lang.

## Executando o projeto
Dentro da pasta raiz, rodar o seguinte comando:
```shell
go run main.go  
```
O projeto estará rodando em localhost na porta 5000
```
http://localhost:5000/
```

Vale ressaltar que para acessar os endpoints de book, você devera gerar um token no endpoint abaixo:

```
http://localhost:5000/api/v1/users/login
```

Obs: Para criar um usuario, utilize seu respectivo endpoint conforme as requisições do Postman.
## Endpoints

<b><span style="color:yellow"> POST</span></b> `http://localhost:5000/api/v1/users`

<b><span style="color:yellow"> POST</span></b> `http://localhost:5000/api/v1/users/login`

<b><span style="color:green"> GET</span></b> `http://localhost:5000/api/v1/books`

<b><span style="color:green"> GET</span></b> `http://localhost:5000/api/v1/books/1`

<b><span style="color:yellow"> POST</span></b> `http://localhost:5000/api/v1/books`

<b><span style="color:blue"> PUT</span></b> `http://localhost:5000/api/v1/books`

<b><span style="color:red"> DELETE</span></b> `http://localhost:5000/api/v1/books/1`

## POSTMAN

Para fazer o import no Postman, siga os passos abaixo:

1) Baixe o arquivo [aqui](./postman)

2) No Postman, crie um novo workspace em Workspace > + New Workspace e defina um nome

3) Clique em Import > File > Upload Files e escolha o arquivo baixado

4) Confirme o Import

5) Pronto! Agora é só acessar a coleção.


## Referências
- [Youtube](https://www.youtube.com/)