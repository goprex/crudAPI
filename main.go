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
	_ "crudapi/docs" 
    	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/spf13/viper"
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

	dbConn := viper.GetString("DB_CONN")
	if dbConn == "" {
		log.Fatal("CRITICAL ERROR: DB_CONN is not set in .env or Environment Variable")

	}

	return Config{
		Port:   port,
		DBConn: dbConn,
	//	DBConn: viper.GetString("DB_CONN"),
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
	
	//
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	defer db.Close()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Routes
	//http.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
	//	w.Write([]byte("OK"))
	//})

	// Health Check yang memverifikasi koneksi API dan Database
   	 http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        // Cek koneksi ke database
		err := db.Ping()

		if err != nil {
		    // Jika database bermasalah, kirim status 503 (Service Unavailable)
		    w.Header().Set("Content-Type", "application/json")
		    w.WriteHeader(http.StatusServiceUnavailable)
		    fmt.Fprintf(w, `{"status": "down", "database": "disconnected", "error": "%s"}`, err.Error())
		    return
		}

        	// Jika semua oke, kirim status 200
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "up", "database": "connected", "message": "Everything is Awesome!"}`))
    	})

	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)
   	http.HandleFunc("/docs/", httpSwagger.WrapHandler)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	http.HandleFunc("/api/checkout", transactionHandler.HandleCheckout) // POST

	// Report Hari ini
	reportRepo := repositories.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)

	http.HandleFunc("/api/report/hari-ini", reportHandler.GetToday)


	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}
