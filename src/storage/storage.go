package storage

import (
	"encoding/json"
	"fmt"
	"gorecipes/src/recipes"
	//"fyne.io/fyne/v2"
	//"gorecipes/src/file"
	//"fyne.io/fyne/v2/dialog"
)

type Storage struct {
	Content []recipes.Recipes `json:"recipes"`
}

func (s *Storage) ToBytes() (*[]byte, error) {
	file, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}
	fmt.Println(file)
	return &file, nil
}

func (s *Storage) AddRecipe(res recipes.Recipes) (*Storage, error) {
	s.Content = append(s.Content, res)

	data, err := s.ToBytes()
	if err != nil {
		return nil, err
	}
	fmt.Println("AddRecipe: ", data, s.Content)

	return s, nil

	//file.Wfile(*data, "recipes.json")

}

func NewStorage(d *[]byte) *Storage {
	var st Storage

	err := json.Unmarshal(*d, &st)

	if err != nil {
		fmt.Println(err)

	}
	return &st
}
