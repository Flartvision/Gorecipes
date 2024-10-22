package file

import (
	//"fyne.io/fyne/v2"
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	st "gorecipes/src/storage"
)

func fileURI() fyne.URI {
	return storage.NewFileURI("recipes.json")
}

func Wfile(data []byte, URI string) error {
	fileURI := storage.NewFileURI(URI)
	file, err := storage.Writer(fileURI)
	fmt.Println("Wfile: ", fileURI, file)
	defer file.Close()
	if err != nil {
		return err
	}
	file.Write(data)
	return err
}

func Rfile(w fyne.Window, stor *st.Storage) *st.Storage {

	dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
		}
		if reader == nil {
			fmt.Println("Пользователь отменил действие")
		}
		data, err := io.ReadAll(reader)
		if err != nil {
			dialog.ShowError(err, w)
		}
		stor = st.NewStorage(w, &data)
		fmt.Println("База данных после загрузки:", stor)

	}, w).Show()

	return stor
}

/*func callFile() fyne.URI {
	d, err := dialog.NewFileOpen()
	}, w).Show()
		dialog.ShowError(err)
	}
}*/
