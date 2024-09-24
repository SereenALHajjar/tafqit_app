package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SereenALHajjar/tafqit"
)

type convertRequest struct {
	Num      int  `json:"Num"`
	Feminine bool `json:"Feminine"`
	Miah     bool `json:"Miah"`
	Billions bool `json:"Billions"`
	AG       bool `json:"AG"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var req convertRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	cnv := tafqit.NumberConverter{
		Num: req.Num,
		Opt: tafqit.Options{
			AG:       req.AG,
			Billions: req.Billions,
			Feminine: req.Feminine,
			Miah:     req.Miah,
		},
	}
	result := cnv.MakeNumber()
	fmt.Fprintf(w, result)
}

func main() {
	http.HandleFunc("/tafqit", handler)

	// Serve static files from ./public directory
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Redirect root to /static/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})

	// fmt.Println("Server is running on http://localhost:8009")
	address := os.Getenv("SERVER_ADDR")
	if address == "" {
		address = ":8009"
	}
	fmt.Println("Server is running on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
