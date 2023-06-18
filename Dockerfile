# Imagem base do Docker
FROM golang:1.16-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos necessários para o contêiner
COPY go.mod go.sum ./

# Executa o download das dependências
RUN go mod download

# Copia o código fonte da API para o contêiner
COPY . .

# Compila o código da API
RUN go build -o main

# Expõe a porta em que a API estará ouvindo
EXPOSE 8080

# Comando para executar a API quando o contêiner for iniciado
CMD ["./main"]