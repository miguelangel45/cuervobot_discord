package structures

type Pokemon struct {
	Species PokemonDefault `json:"species"`
	Sprites PokemonSprites `json:"sprites"`
}
type PokemonDefault struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonSprites struct {
	Back_default       string `json:"back_default"`
	Back_female        string `json:"back_female"`
	Back_shiny         string `json:"back_shiny"`
	back_shiny_female  string `json:"back_shiny_female"`
	Front_default      string `json:"front_default"`
	Front_female       string `json:"front_female"`
	Front_shiny        string `json:"front_shiny"`
	Front_shiny_female string `json:"front_shiny_female"`
}
