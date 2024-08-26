package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

// Anão representa um anão
type Anao struct {
	Nome                  string  `json:"nome"`                    // Nome do anão é obrigatorio
	Altura                float64 `json:"altura"`                  // Aqui só pode ser tiver 1,50 ou abaixo, fora de 1,50 não pode participar
	Idade                 int     `json:"idade"`                   // A idade abaixo dos 18 anos não é permitido participar
	Raca                  string  `json:"raca"`                    // Tanto faz a cor
	Regiao                string  `json:"regiao"`                  // Região pode ser tanto do Brasil como de fora
	ValorVenda            float64 `json:"valor_venda"`             // Valor para venda
	ValorAluguel          float64 `json:"valor_aluguel"`           // Valor para aluguel
	ValorVendaFormatado   string  `json:"valor_venda_formatado"`   // Valor formatado para venda
	ValorAluguelFormatado string  `json:"valor_aluguel_formatado"` // Valor formatado para aluguel
}

var anoes []Anao

func conectarBancoDeDados() *sql.DB {
	// Obtendo os valores das variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Conectando ao banco de dados PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Banco de dados não está respondendo:", err)
	}

	fmt.Println("Conexão com banco de dados estabelecida com sucesso!")
	return db
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/anoes", criarAnao).Methods("POST")
	r.HandleFunc("/anoes", listarAnoes).Methods("GET")
	r.HandleFunc("/anoes/{nome}", buscarAnao).Methods("GET")
	r.HandleFunc("/anoes/{nome}", deletarAnao).Methods("DELETE")
	r.HandleFunc("/anoes/{nome}", atualizarAnao).Methods("PUT")

	http.ListenAndServe(":8080", r)
}

// criarAnao adiciona um novo anão
func criarAnao(w http.ResponseWriter, r *http.Request) {
	var anao Anao
	_ = json.NewDecoder(r.Body).Decode(&anao)

	if anao.Nome == "" || anao.Altura <= 0 || anao.Idade <= 0 || anao.Raca == "" || anao.Regiao == "" {
		http.Error(w, "Todos os campos devem ser preenchidos corretamente", http.StatusBadRequest)
		return
	}

	if anao.Altura > 1.50 {
		http.Error(w, "A altura deve ser 1.50 ou abaixo", http.StatusBadRequest)
		return
	}

	if anao.ValorVenda < 0 || anao.ValorAluguel < 0 {
		http.Error(w, "Os valores de venda e aluguel não podem ser negativos", http.StatusBadRequest)
		return
	}

	anao.ValorVendaFormatado = fmt.Sprintf("R$ %.2f", anao.ValorVenda)
	anao.ValorAluguelFormatado = fmt.Sprintf("R$ %.2f", anao.ValorAluguel)

	anoes = append(anoes, anao)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(anao)
}

// listarAnoes retorna todos os anões
func listarAnoes(w http.ResponseWriter, r *http.Request) {
	// Exemplo básico de paginação
	limite, _ := strconv.Atoi(r.URL.Query().Get("limite"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if limite <= 0 {
		limite = 10
	}

	inicio := offset
	fim := inicio + limite

	if fim > len(anoes) {
		fim = len(anoes)
	}

	json.NewEncoder(w).Encode(anoes[inicio:fim])
}

// buscarAnao retorna um anão específico pelo nome
func buscarAnao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range anoes {
		if item.Nome == params["nome"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Anão não encontrado", http.StatusNotFound)
}

// deletarAnao deleta um anão pelo nome
func deletarAnao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range anoes {
		if item.Nome == params["nome"] {
			anoes = append(anoes[:index], anoes[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

// AtualizaAnão que você cadastrou
func atualizarAnao(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var novoAnao Anao
	_ = json.NewDecoder(r.Body).Decode(&novoAnao)

	if novoAnao.ValorVenda < 0 || novoAnao.ValorAluguel < 0 {
		http.Error(w, "Os valores de venda e aluguel não podem ser negativos", http.StatusBadRequest)
		return
	}

	novoAnao.ValorVendaFormatado = fmt.Sprintf("R$ %.2f", novoAnao.ValorVenda)
	novoAnao.ValorAluguelFormatado = fmt.Sprintf("R$ %.2f", novoAnao.ValorAluguel)

	for index, item := range anoes {
		if item.Nome == params["nome"] {
			anoes[index] = novoAnao
			json.NewEncoder(w).Encode(novoAnao)
			return
		}
	}
	http.Error(w, "Anão não encontrado", http.StatusNotFound)
}
