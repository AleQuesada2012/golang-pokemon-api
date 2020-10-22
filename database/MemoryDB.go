package database

type Pokemon struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}

var PokemonDB = []Pokemon{
	Pokemon{ID: "1", Name: "Pikachu", Type: "Electric"},
	Pokemon{ID: "2", Name: "Charmeleon", Type: "Fire"},
}
