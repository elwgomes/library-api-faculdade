# Biblioteca API

Uma API RESTful para gerenciamento de um sistema de biblioteca, construída em Go (Golang), utilizando o ORM GORM para operações de banco de dados e Docker para conteinerização.

## Funcionalidades

- Gerenciar recursos da biblioteca (ex.: livros e autores).
- Integração com banco de dados MySQL utilizando o GORM.
- Arquitetura modular para fácil manutenção e escalabilidade.

## Tecnologias Utilizadas

- **Linguagem de Programação:** Go (Golang)
- **Framework Web:** Gorilla Mux
- **Banco de Dados:** MySQL (via Docker)
- **ORM:** GORM
- **Conteinerização:** Docker

## Estrutura do Projeto

```
library-api-faculdade/
├── cmd/                 # Pontos de entrada da aplicação
├── core/                # Lógica principal da aplicação
├── infra/               # Configuração de infraestrutura (ex.: conexão com o banco de dados)
├── server/              # Servidor HTTP e rotas
├── .env                 # Variáveis de ambiente
├── docker-compose.yml   # Configuração do Docker Compose
├── go.mod               # Dependências do módulo Go
├── go.sum               # Checksums das dependências
└── README.md            # Documentação do projeto
```

## Pré-requisitos

- [Go](https://golang.org/) (versão 1.23.2 ou superior)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Configuração e Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-repo/library-api-faculdade.git
   cd library-api-faculdade
   ```

2. Configure o arquivo `.env` com as variáveis de ambiente apropriadas. Exemplo:

   ```env
   MYSQL_USER=mysql
   MYSQL_PASSWORD=mysql
   MYSQL_DATABASE=library_db
   MYSQL_ROOT_PASSWORD=mysql
   ```

3. Inicie o banco de dados utilizando o Docker Compose:

   ```bash
   docker-compose up -d
   ```

4. Execute a aplicação:

   ```bash
   go run cmd/main.go
   ```

## Endpoints da API

Aqui estão alguns exemplos de endpoints (supondo que o servidor esteja rodando em `http://localhost:8080`):

- **GET /books**: Retorna todos os livros.
- **POST /books**: Adiciona um novo livro.
- **PUT /books/{id}**: Atualiza os detalhes de um livro.
- **DELETE /books/{id}**: Remove um livro.

## Utilizando Docker

Para executar a aplicação em um ambiente completamente conteinerizado:

1. Crie a imagem do container da aplicação (se necessário).
2. Atualize o `docker-compose.yml` para incluir o serviço da aplicação.
3. Inicie todos os serviços:

   ```bash
   docker-compose up --build
   ```

## Diretrizes para Contribuição

1. Faça um fork do repositório e clone-o localmente.
2. Crie um novo branch para sua funcionalidade ou correção de bug.
3. Faça commit e push das suas alterações.
4. Abra um pull request para revisão.



## Colaboradores

Este projeto foi desenvolvido com a colaboração dos seguintes membros:

- [Edson Leonardo](https://github.com/elwgomes)
- [Kalil Farias](https://github.com/KalilFarias)
- [Larissa Marques](https://github.com/Laryhop)
- [Frankilin Eduardo](https://github.com/mancharger)
- [Roberto Beltrão](https://github.com/robertobeltraoo)
- [Pedro Henrique](https://github.com/)

## Licença

Este projeto está licenciado sob a Licença MIT.
