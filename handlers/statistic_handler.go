package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Products struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Statistic struct {
	ProductCount int `json:"product_count"`
}

func GetStatistics(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://go-api-navy.vercel.app/product")
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("upstream error %d: %s", resp.StatusCode, string(body)), http.StatusBadGateway)
		return
	}

	var products []Products
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statistic := Statistic{
		ProductCount: len(products),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statistic)
}
