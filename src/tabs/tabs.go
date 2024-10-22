package tabs

import (
	"gorecipes/src/storage"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewCard(st *storage.Storage) (names []string, description []string) {
	names = make([]string, 0, 100)
	description = make([]string, 0, 100)
	for _, v := range st.Content {
		names = append(names, v.Name)
		description = append(description, v.Description)
	}
	return names, description

}

func SpavnList(st *storage.Storage) *fyne.Container {
	/*list := widget.NewList(
		func() int {
			return len(names) // Возвращаем количество элементов в списке
		},
		func() fyne.CanvasObject {
			// Создаем новый элемент списка (карточку)
			return widget.NewCard("", "", nil)
		},
		func(i int, o fyne.CanvasObject) {
			card := o.(*widget.Card)
			card.SetTitle(names[i])                                          // Устанавливаем заголовок карточки
			card.SetContent(widget.NewLabel("Содержимое " + description[i])) // Устанавливаем содержимое карточки
		},
	)

	return list*/
	c := container.NewVBox()
	for _, n := range st.Content {
		card := widget.NewCard(
			n.Name,
			n.Description,
			container.NewVBox(widget.NewButton("Показать больше информации", func() {}),
				widget.NewButton("Удалить рецепт", func() {}),
			),
		)
		c.Add(card)
	}

	return c
}
