package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"os"
	"strings"
	"fmt"
	"log"

	"crudapi/database"
	"crudapi/repositories"
	"crudapi/services"
	"crudapi/handlers"

	"github.com/spf13/viper"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	categories []Category
	nextID     = 1
	mu         sync.RWMutex
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")

	switch r.Method {
	case "GET":
		if idStr == "" {
			mu.RLock()
			json.NewEncoder(w).Encode(categories)
			mu.RUnlock()
			return
		}
		id, _ := strconv.Atoi(idStr)
		mu.RLock()
		for _, c := range categories {
			if c.ID == id {
				json.NewEncoder(w).Encode(c)
				mu.RUnlock()
				return
			}
		}
		mu.RUnlock()
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)

	case "POST":
		var c Category
		if json.NewDecoder(r.Body).Decode(&c) != nil || c.Name == "" {
			http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
			return
		}
		mu.Lock()
		c.ID = nextID
		nextID++
		categories = append(categories, c)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(c)

	case "PUT":
		id, _ := strconv.Atoi(idStr)
		var input Category
		if json.NewDecoder(r.Body).Decode(&input) != nil || input.Name == "" {
			http.Error(w, `{"error":"bad request"}`, http.StatusBadRequest)
			return
		}
		mu.Lock()
		for i := range categories {
			if categories[i].ID == id {
				categories[i].Name = input.Name
				categories[i].Description = input.Description
				json.NewEncoder(w).Encode(categories[i])
				mu.Unlock()
				return
			}
		}
		mu.Unlock()
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)

	case "DELETE":
		id, _ := strconv.Atoi(idStr)
		mu.Lock()
		for i := range categories {
			if categories[i].ID == id {
				categories = append(categories[:i], categories[i+1:]...)
				mu.Unlock()
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		mu.Unlock()
		http.Error(w, `{"error":"not found"}`, http.StatusNotFound)

	default:
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

type Config struct {
	Port string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","-"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := Config {
		Port: viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	//Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)



	mux := http.NewServeMux()
	mux.HandleFunc("GET /categories", handler)
	mux.HandleFunc("POST /categories", handler)
	mux.HandleFunc("GET /categories/{id}", handler)
	mux.HandleFunc("PUT /categories/{id}", handler)
	mux.HandleFunc("DELETE /categories/{id}", handler)
//	log.Println("Server on http://localhost:8000")
//	log.Fatal(http.ListenAndServe(":8000", mux))

	// Setup routes
	http.HandleFunc("/api/produk", productHandler.HandleProducts)
	http.HandleFunc("/api/produk/", productHandler.HandleProductByID)

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("gagal running server", err)
	}

}
