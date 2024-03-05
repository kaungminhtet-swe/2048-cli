package ui

import (
	"io"
	"log"
	"text/template"

	"github.com/mr-kmh/2048-cli/pkg/board"
)

func SetUI(uiFilename string) board.Render {
	t, err := template.ParseFiles(uiFilename)
	if err != nil {
		log.Fatal("err in template", err)
	}
	return func(w io.Writer, s [4][4]int) error {
		return t.Execute(w, s)
	}
}

func New(uiFileName string) board.Render {
	return SetUI(uiFileName)
}
