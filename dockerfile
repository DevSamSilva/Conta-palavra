# Usa a imagem oficial do Go
FROM golang:1.22.2-alpine

# Diretório de trabalho
WORKDIR /app

# Copia o código para o container
COPY . .

# Baixa as dependências (se go.mod existir)
RUN go mod download

# Compila o programa
RUN go build -o word-counter

# Define o comando padrão
CMD ["./word-counter"]