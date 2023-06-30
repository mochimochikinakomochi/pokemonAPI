package handleSQL

type Pokemon struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Types []string `json"Types"`
	Status Pokemon_status `json"Status"`
	Category string `json:"Category"`
	Height int `json"Height"`
	Weight int `json"Weight"`
	ImageURL string `json"ImageURL"`
}

type Pokemons struct {
	id    int    `json:"ID"`
	name  string `json:"Name"`
	category string `json:"Category"`
	height int `json"Height"`
	weight int `json"Weight"`
	imageURL string `json"ImageURL"`
}

type Pokemon_type struct {
	id int `json"pokemonID"`
	types []string `json"pokemonType"`
}

type Pokemon_status struct {
	H int `json"H"`
	A int `json"A"`
	B int `json"B"`
	C int `json"C"`
	D int `json"D"`
	S int `json"S"`
}

type Move struct {
	ID int `json"ID"`
	Name string `json"Name"`
	Type string `json"Type"`
	Class string `json"Class"`
	PP int `json"PP"`
	Accuracy int `json"Accuracy"`
	Priority int `json"Priority"`
	Power int `json"Power"`
	Description string `json"Description"`
}

