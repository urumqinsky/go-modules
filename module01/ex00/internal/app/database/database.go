package database

// Аннотирование структур `json:"" xml:""`

type Recipes struct {
	Cake []Cake `json:"cake" xml:"cake"`
}

type Cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []Ingredients `json:"ingredients" xml:"ingredients>item"`
}

type Ingredients struct {
	IngredientName  string `json:"ingredient_name" xml:"itemname"`
	IngredientCount string `json:"ingredient_count" xml:"itemcount"`
	IngredientUnit  string `json:"ingredient_unit" xml:"itemunit,omitempty"`
}
