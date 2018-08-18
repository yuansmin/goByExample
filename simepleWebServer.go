/* a simple web server
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := &database{"shoes": 50, "T-shirt": 70}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range *db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}
	priceString := req.URL.Query().Get("price")
	price, err := strconv.Atoi(priceString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error: %s\n", err)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "update %s success\n%s: %s\n", item, item, db[item])
}
