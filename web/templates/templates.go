package templates

import (
	"embed"
	_ "embed"
	"html/template"
	"io/fs"
)

var (
	//go:embed *.gohtml
	files     embed.FS
	Templates map[string]*template.Template
)

func init() {
	if Templates == nil {
		Templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, ".")
	if err != nil {
		panic(err)
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, err := template.ParseFS(files, tmpl.Name())
		if err != nil {
			panic(err)
		}

		Templates[tmpl.Name()] = pt
	}
}
