package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"crudapi/database"
	"crudapi/handlers"
	"crudapi/repositories"
	"crudapi/services"

	"github.com/spf13/viper"
	_ "crudapi-ku/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func loadConfig() Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	port := viper.GetString("PORT")
	if port == "" {
		port = "8080" // fallback lokal
	}

	return Config{
		Port:   port,
		DBConn: viper.GetString("DB_CONN"),
	}
}

// @title           Kasir API Service
// @version         1.0
// @description     Dokumentasi API untuk aplikasi Kasir menggunakan Clean Architecture.
// @host            crudapi-ku.up.railway.app
// @BasePath        /

func main() {
	config := loadConfig()

	// Init DB (Supabase)
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Routes
	http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	// Pastikan ada '/' di akhir "/docs/"
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
