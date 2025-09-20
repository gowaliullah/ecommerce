package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productList)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my ecommerce......!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", welcome)

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running port on:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error from server", err)
	}
}

func init() {

	prd1 := Product{
		ID:          1,
		Title:       "Wireless Noise-Canceling Headphones",
		Description: "Experience crystal-clear audio with our top-of-the-line wireless headphones. Features include active noise-cancellation, a 20-hour battery life, and comfortable over-ear cushions.",
		Price:       199.99,
		ImgUrl:      "https://example.com/images/headphones.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Ergonomic Office Chair",
		Description: "Designed for all-day comfort and support. This chair features adjustable lumbar support, a breathable mesh back, and a smooth 360-degree swivel.",
		Price:       249.50,
		ImgUrl:      "https://example.com/images/chair.jpg",
	}
	prd3 := Product{
		ID:          3,
		Title:       "4K Ultra HD Smart TV",
		Description: "Immerse yourself in stunning visuals with this 55-inch smart TV. It comes with built-in streaming apps, HDR support, and a sleek, minimalist design.",
		Price:       799.00,
		ImgUrl:      "https://example.com/images/tv.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Portable Power Bank",
		Description: "Never run out of battery again. This compact power bank offers a 20,000mAh capacity, fast-charging ports, and a durable, non-slip finish.",
		Price:       45.75,
		ImgUrl:      "https://example.com/images/powerbank.jpg",
	}
	prd5 := Product{
		ID:          5,
		Title:       "Stainless Steel Coffee Maker",
		Description: "Brew the perfect cup of coffee every time. This maker features a 12-cup capacity, a programmable timer, and a sleek stainless steel design that complements any kitchen.",
		Price:       89.99,
		ImgUrl:      "https://example.com/images/coffeemaker.jpg",
	}
	productList = append(productList, prd1, prd2, prd3, prd4, prd5)
}
