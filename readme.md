## Anão API

Descrição
Uma API RESTful em Go para gerenciar um catálogo de anões. Utilize PostgreSQL como banco de dados e Docker para facilitar a execução e o gerenciamento.

* ## :wrench: technologies used
<div>
<img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
<img src="https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white" />
<img src="https://img.shields.io/badge/DBeaver-8C8C8C?style=for-the-badge&logo=dbeaver&logoColor=white" />
<img src="https://img.shields.io/badge/Dockerfile-2496ED?style=for-the-badge&logo=docker&logoColor=white" />
<img src="https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white" />
</div>

Instalação
Requisitos:
Docker
Docker Compose
Go (opcional)

## Clonando o Repositório:
```
git clone https://github.com/seu-usuario/anao-api.git
cd anao-api
```
Executando a Aplicação:
````
Bash
docker-compose up --build
TABELA DE PREÇOS DE ANÕES:
````
Uso

1. Iniciar o Servidor Go
Primeiro, certifique-se de que o servidor está rodando. No terminal, dentro do diretório do seu projeto, execute:
Bash
```
go run main.go
```

2. Configurar o Postman
Abra o Postman e siga as instruções para testar os diferentes endpoints.

a. Criar um Anão (POST /anoes)
Selecione o método POST.
Insira a URL: `http://localhost:8080/anoes.`
Vá para a aba Body e selecione raw e JSON.
Insira o JSON com os dados do anão, como no exemplo abaixo:

```
{
  "nome": "Anão Exemplo",
  "altura": 1.45,
  "idade": 30,
  "raca": "Branca",
  "regiao": "Anão Paulista",
  "valor_venda": 10000.00,
  "valor_aluguel": 500.00
}
```
Todas as rotas
``
1- 
Criar um Anão (POST /anoes)
**URL**: `http://localhost:8080/anoes`

Método: POST
Corpo (Body): JSON

2- Listar Todos os Anões (GET /anoes) **URL**: `http://localhost:8080/anoes`
Método: GET

3- Buscar um Anão por Nome (GET /anoes/{nome}) **URL**: `http://localhost:8080/anoes/{nome}`
Método: GET

4- Atualizar um Anão (PUT /anoes/{nome}) **URL**: `http://localhost:8080/anoes/{nome}`
Método: PUT

5- Deletar um Anão (DELETE /anoes/{nome}) **URL**: `http://localhost:8080/anoes/{nome}`
Método: DELETE

```
Resumo das Rotas Atualizadas:
POST /anoes: Cria um novo anão.
GET /anoes: Lista todos os anões (com paginação opcional).
GET /anoes/{nome}: Busca um anão pelo nome.
PUT /anoes/{nome}: Atualiza as informações de um anão existente.
DELETE /anoes/{nome}: Deleta um anão pelo nome.
````

Outras tabelas por preço 

````
TABELA DE PREÇOS DE ANÕES:

- Anão boliviano: R$ 170,86  
- Anão russo: R$ 104,57  
- Anão argentino: R$ 10,99  
- Anão colombiano: R$ 260,45  
- Anão rebaixado (edição especial): R$ 456,79  
- Anão latino: R$ 159,99  
- Anão kppoper:   
- Anão baiano: R$ 10.000,00  
- Anão blindado: R$ 345,49  
- Anão cachaceiro: R$ 68,99  
- Anão amadeirado: R$ 112,10  
- Filhote de anão: $20  
- Anão cabeçudo cearense: R$ 1.059,49  
- Anão petista: R$ 468,45  
- Anão orgânico: R$ 420,00  
- Anão com pau: R$ 6.969,00  
- Anão lombrado: R$ 40,20  
- Anão industrializado (dano maciço na camada de ozônio): R$ 690,00  
- Anão importado: US$ 200,00 (R$ 1.052)  
- Anão que toca tuba: R$ 340  
- Anão Goldfinger: 10 mil em ouro  
- Anão mineiro: 2 queijos de 350g  
- Anão 12 anos (envelhecido no tonel de carvalho): US$ 32.000,00 (R$ 168.230,40)  
- Anão piloto de fuga: R$ 4.500,09  
- Anão Judeu: R$ 50,00  
- Anão alto: R$ 100.000,00  
- Anão careca: R$ 145,00
````








