package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/Etevaldo15/go-nif-validator/internal/router"
)

// @title NIF Validator API
// @version 1.0
// @description API para validação de NIF angolano
// @BasePath /
func main() {
	// Carregar .env (não falha em produção caso não exista)
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env não encontrado, usando variáveis de ambiente do sistema")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configura o roteador da API
	r := router.SetupRouter()

	// Configura CORS (permitir requisições do frontend local)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:36605", // porta do seu frontend
			"http://127.0.0.1:36605",
			"http://localhost:*",     // opcional: aceitar qualquer porta local
			"http://127.0.0.1:*",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Envolve o roteador com o middleware CORS
	handler := c.Handler(r)

	fmt.Printf("🚀 Servidor iniciado na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}