package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/Nivedithabp/receipt-processor-challenge/docs"
	"github.com/spf13/viper"
	"github.com/Nivedithabp/receipt-processor-challenge/routes"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Cors struct {
		AllowedOrigins []string `mapstructure:"allowedOrigins"`
	} `mapstructure:"cors"`
	Swagger struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"swagger"`
}

var config Config

// LoadConfig loads config.json
func LoadConfig(filename string) error {
	viper.SetConfigFile(filename)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return viper.Unmarshal(&config)
}

func main() {
	// Load config.json
	if err := LoadConfig("config.json"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	router := mux.NewRouter()

	// Register Routes
	routes.RegisterRoutes(router)

	// Enable Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// CORS settings
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	origins := handlers.AllowedOrigins(config.Cors.AllowedOrigins)

	// Start the server
	fmt.Printf("🚀 Server running on http://localhost:%s\n", config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+config.Server.Port, handlers.CORS(headers, methods, origins)(router)))
}
