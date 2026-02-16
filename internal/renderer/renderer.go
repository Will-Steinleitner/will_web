package renderer

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const rendererTAG = "Renderer"

type Renderer struct {
	templateCache map[string]*template.Template
}

func NewRenderer() *Renderer {
	log.Println(rendererTAG, ": building renderer..")
	cache, err := newTemplateCache()
	if err != nil {
		log.Fatal(rendererTAG, err)
	}
	log.Println(rendererTAG, ": renderer built")

	return &Renderer{
		cache,
	}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("./ui/templates/html/*")
	if err != nil {
		log.Fatal(rendererTAG, err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		log.Println(rendererTAG, ": caching template -", name)

		templateSet, err := template.ParseFiles("./ui/templates/html/base.gohtml", page)
		if err != nil {
			log.Fatal(rendererTAG, err)
		}

		cache[name] = templateSet
	}

	return cache, nil
}
func (renderer *Renderer) RenderTemplate(writer http.ResponseWriter, tmpl string, data interface{}) {
	t, exists := renderer.templateCache[tmpl]
	if !exists || t == nil {
		http.Error(writer, rendererTAG, http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(writer, "base", data); err != nil {
		log.Println(rendererTAG+": template is missing", err)
		log.Println(t.Name())
		http.Error(writer, rendererTAG+": render error", http.StatusInternalServerError)
	}
}
