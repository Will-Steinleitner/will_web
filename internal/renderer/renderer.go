package renderer

import (
	"html/template"
	"log"
	"net/http"
)

const rendererTAG = "renderer"

type Renderer struct {
	templateCache map[string]*template.Template
}

func NewRenderer(fullCache map[string]*template.Template) *Renderer {
	return &Renderer{
		fullCache,
	}
}

func (renderer *Renderer) RenderTemplate(writer http.ResponseWriter, tmpl string, data interface{}) {
	t, exists := renderer.templateCache[tmpl]
	if !exists || t == nil {
		http.Error(writer, "Template fehlt", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(writer, "base", data); err != nil {
		log.Println(rendererTAG, err)
		log.Println(t.Name())
		http.Error(writer, "Render error", http.StatusInternalServerError)
	}
}
