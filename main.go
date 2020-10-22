package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"pokemon-api/database"
)
type pokeArr [] database.Pokemon

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database.PokemonDB)

}
func addNewPokemon(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var pokemon database.Pokemon

	json.Unmarshal(requestBody, &pokemon)
	found := false
	for i := 0; i < len(database.PokemonDB); i++ {
		if database.PokemonDB[i] == pokemon {
			found = true
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	if  !found {
		database.PokemonDB = append(database.PokemonDB, pokemon)
	}
	w.WriteHeader(http.StatusOK)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(commonMiddleware)
	myRouter.HandleFunc("/pokemon", getAllPokemon).Methods("GET")
	myRouter.HandleFunc("/pokemon/add", addNewPokemon).Methods("POST")
	log.Fatal(http.ListenAndServe(":8095", myRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Pokemon Rest API")
	handleRequests()
}
