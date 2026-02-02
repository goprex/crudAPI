package models

// Menggunakan pointer (*) agar jika datanya kosong, 
// JSON-nya bisa jadi null (omitempty)
// Nested Model
type Product struct {
	ID 	int	`json:"id"`
	Name	string	`json:"name"`
	Price	int	`json:"price"`
	Stock	int	`json:"stock"`
	CategoryID int	`json:"category_id"`
	Category *Category	`json:"category,omitempty"`
}

