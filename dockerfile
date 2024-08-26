# Escolhe uma imagem base com Go
FROM golang:1.23

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Configura o proxy para evitar problemas de rede (se necessário)
ENV GOPROXY=https://proxy.golang.org,direct

# Copia o arquivo go.mod e go.sum para o diretório de trabalho no container
COPY go.mod ./
COPY go.sum ./

# Faz o download das dependências necessárias
RUN go mod download

# Copia o código fonte para o diretório de trabalho no container
COPY . .

# Compila o código Go
RUN go build -o main .

# Expõe a porta em que a aplicação vai rodar
EXPOSE 8080

# Comando para rodar a aplicação quando o container for iniciado
CMD ["./main"]
