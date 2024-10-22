package storage

import (
	"encoding/json"
	"fmt"
	"gorecipes/src/recipes"

	"fyne.io/fyne/v2"
	//"gorecipes/src/file"

	"fyne.io/fyne/v2/dialog"
)

type Storage struct {
	Content []recipes.Recipes `json:"recipes"`
}

func (s *Storage) ToBytes(w fyne.Window) *[]byte {
	file, err := json.Marshal(s)

	if err != nil {
		dialog.ShowError(err, w)
	}
	fmt.Println(file)
	return &file
}

func (s *Storage) AddRecipe(res recipes.Recipes, w fyne.Window) *Storage {
	s.Content = append(s.Content, res)

	data := s.ToBytes(w)

	fmt.Println("AddRecipe: ", data, s.Content)

	return s

	//file.Wfile(*data, "recipes.json")

}

func NewStorage(w fyne.Window, d *[]byte) *Storage {

	var st Storage
	err := json.Unmarshal(*d, &st)

	if err != nil {
		fmt.Println(err)

	}
	return &st
}
