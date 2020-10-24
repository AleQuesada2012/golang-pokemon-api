package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"pokemon-api/database"
	"testing"
)
type testPokemon struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}
func TestPokemonLoaded(t *testing.T)  {
	if len(database.PokemonDB) != 2 {
		t.Error()
	}
}

func TestGetPokemon(t *testing.T) {
	request, error := http.NewRequest("GET", "/pokemon", nil)
	if error != nil {
		t.Fatal(error)
	}
	reqResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllPokemon)

	handler.ServeHTTP(reqResponse, request)
	if status := reqResponse.Code; status != http.StatusOK {
		t.Errorf("wrong status returned: got %v want %v", status,
		http.StatusOK)
	}

	// check the response body and compare it to expected
	expected := `[{"Id":"1","Name":"Pikachu","Type":"Electric"},{"Id":"2","Name":"Charmeleon","Type":"Fire"}]
` // must have the end line character at the end of the expected value to pass the test
	if reqResponse.Body.String() != expected {
		t.Errorf("handler returned unwanted body. got %v wanted %v",
			reqResponse.Body.String(), expected)
	}
}

func TestAddPokemon(t *testing.T) {
	// we have to create a 'body' for the request since we're testing a post method
	var jsonTestStr = []byte (`{"Id":3,"Name":"Piplup","Type":"Water"}`)
	request, error := http.NewRequest("POST","pokemon/add", bytes.NewBuffer(jsonTestStr))
	if error != nil {
		t.Fatal(error)
	}
	request.Header.Set("Content-Type", "application/json")
	reqResponse := httptest.NewRecorder()
	handler := http.HandlerFunc(addNewPokemon)
	handler.ServeHTTP(reqResponse, request)
	if status := reqResponse.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong code. got %v wanted %v", status, http.StatusOK)
	}
	expected := reqResponse.Body.String()

	if ! (len(expected) == 0) {
		t.Errorf("returned unwanted body. got %v expected none",expected)
	}
}