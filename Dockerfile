# builder
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Instalar git (necessário para módulos privados ou dependências via git)
RUN apk add --no-cache git

# Copiar módulos e baixar
COPY go.mod go.sum ./
RUN go mod download

# Copiar código e buildar
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go-nif-validator ./cmd/server

# final
FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /root

# Copiar binário
COPY --from=builder /go-nif-validator .

# Criar usuário não-root (boa prática de segurança)
RUN adduser -D -s /bin/sh validator
USER validator

EXPOSE 8080
ENTRYPOINT ["./go-nif-validator"]