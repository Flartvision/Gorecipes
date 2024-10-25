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

func fileURI(URI string) fyne.URI {
	return storage.NewFileURI(URI)
}

func Wfile(data []byte, URI string) error {
	fileURI, err := storage.ParseURI(URI)
	if err != nil {
		return err
	}
	file, err := storage.Writer(fileURI)
	if err != nil {
		return err
	}
	fmt.Println("WfileURI: ", fileURI)
	defer file.Close()
	fmt.Println("Wfile file:", file)

	file.Write(data)
	return err
}

func Rfile(URI string) ([]byte, error) {
	var err error
	fileURI, err := storage.ParseURI(URI)
	if err != nil {
		return nil, err
	}
	file, err := storage.Reader(fileURI)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GuiReader(w fyne.Window, stor *st.Storage) {
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
		stor = st.NewStorage(&data)
		fmt.Println("База данных после загрузки:", stor)

	}, w).Show()

}
