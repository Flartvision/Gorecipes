package main

import (
	"fmt"
	"gorecipes/src/recipes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gorecipes/src/file"
	"gorecipes/src/storage"
	"gorecipes/src/tabs"
	//"time"
)

var content = []recipes.Recipes{}
var tab2Con fyne.Container

func main() {
	a := app.NewWithID("gorecipes")
	wMain := a.NewWindow("Лисьи рецептики")
	data := []byte{}
	st := storage.NewStorage(&data)

	var err error
	// st = file.Rfile(wMain, st)

	rootURI := a.Storage().RootURI().String() + "/recipes.json"
	data, err = file.Rfile(rootURI)
	if err != nil {
		dialog.ShowError(err, wMain)

	}

	st = storage.NewStorage(&data)
	// if err != nil {
	// 	dialog.ShowError(err, wMain)
	//
	// }

	dialog.ShowInformation("Корневой путь программы: ", rootURI, wMain)
	fmt.Println("Storage после добавления рецепта: ", st)
	//
	loadStor := widget.NewButton("Загрузить данные", func() {
		data, err = file.Rfile(rootURI)
		st = storage.NewStorage(&data)
		tabs.SpavnList(st)
		fmt.Println(st)
		tab2Con = *tabs.SpavnList(st)
		tab2Con.Show()
	})

	btnFold := widget.NewButton("Выбрать папку", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}

			if uri == nil {
				fmt.Println("Пользователь отменил выбор")
				return
			}

			dialog.ShowInformation("Выбранная папка:", uri.Path(), wMain)

		}, wMain).Show()
	})

	saveButton := widget.NewButton("Сохранить файл", func() {
		data, err := st.ToBytes()
		if err != nil {
			dialog.ShowError(err, wMain)
		}
		file.Wfile(*data, rootURI)
		dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				return
			}
			if uc == nil {
				return
			}
			defer uc.Close()
			_, err = uc.Write(*data)
			if err != nil {
				dialog.ShowError(err, wMain)
			}
		}, wMain).Show()
	})

	tab1 := container.NewVBox(btnFold, saveButton, loadStor)
	tab21 := container.NewVBox(
		widget.NewButton("Добавить рецепт", func() { AddRecipe(wMain, a, st) }),
		&tab2Con,
	)
	tab2 := container.NewVScroll(tab21)

	tabes := container.NewAppTabs(
		container.NewTabItem("Расписание", tab1),
		container.NewTabItem("Рецепты", tab2),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabes.SetTabLocation(container.TabLocationBottom)

	wMain.SetContent(tabes)
	// go func() {
	// 	for range time.Tick(time.Minute / 10) {
	// 		for range st.Content {
	// 			tab21.Add(container.NewVBox(tabs.SpavnList(st)))
	// 			tab21.Refresh()
	// 		}
	// 	}
	// }()

	wMain.ShowAndRun()
}

func AddRecipe(w fyne.Window, a fyne.App, st *storage.Storage) {
	name := widget.NewEntry()
	description := widget.NewMultiLineEntry()

	confButton := widget.NewButton("Добавить", func() {
		MyREcipe := recipes.NewRecipe(name.Text, description.Text, "")
		fmt.Println("MyRecipe: ", MyREcipe)

		st, err := st.AddRecipe(*MyREcipe)
		if err != nil {
			dialog.ShowError(err, w)
		}
		fmt.Println("Storage после добавления рецепта: ", st)

	})

	content := container.NewVScroll(container.NewVBox(name, description, confButton))

	w2 := a.NewWindow("Добавление рецепта")
	w2.SetContent(content)
	w2.Show()

}
