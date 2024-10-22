package recipes

type Recipes struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	URI         string `json:"URI"`
}

type RecipesList struct {
	data []Recipes
}

func NewRecipe(name, description, URI string) *Recipes {
	return &Recipes{
		Name:        name,
		Description: description,
		URI:         URI,
	}
}

func (r *Recipes) AddRecipes(data []Recipes) {
	data = append(data, *r)
}
