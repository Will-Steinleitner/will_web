package renderer

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const rendererTAG = "Renderer"

type Renderer struct {
	templateCache map[string]*template.Template
	htmlCache     map[string][]byte
}

func NewRenderer() *Renderer {
	log.Println(rendererTAG, ": building renderer..")
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(rendererTAG, err)
	}

	htmlCache, err := newHTMLCache()
	if err != nil {
		log.Fatal(rendererTAG, err)
	}
	log.Println(rendererTAG, ": renderer built")

	return &Renderer{
		templateCache: templateCache,
		htmlCache:     htmlCache,
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
		var templateSet *template.Template
		var err error

		ext := filepath.Ext(name) //end of a file e.g .gohtml or .html
		if ext == ".gohtml" {
			log.Println(rendererTAG, ": caching template -", name)
			templateSet, err = template.ParseFiles(
				"./ui/templates/html/base.gohtml",
				page,
			)
		} else {
			log.Println(rendererTAG, ": caching html -", name)
			templateSet, err = template.ParseFiles(page)
		}

		if err != nil {
			log.Fatal(rendererTAG, err)
		}

		cache[name] = templateSet
	}

	return cache, nil
}
func newHTMLCache() (map[string][]byte, error) {
	cache := make(map[string][]byte)

	pages, err := filepath.Glob("./ui/static/html/*")
	if err != nil {
		log.Fatal(rendererTAG, err)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		log.Println(rendererTAG, ": caching html -", name)
		content, err := os.ReadFile(page)
		if err != nil {
			log.Fatal(rendererTAG, err)
		}
		log.Println(rendererTAG, ": caching html -", content)

		cache[name] = content
	}
	return cache, err
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
func (renderer *Renderer) RenderHTML(writer http.ResponseWriter, html string) {
	content, exists := renderer.htmlCache[html]
	if !exists {
		http.Error(writer, "html fehlt", http.StatusInternalServerError)
		return
	}

	if _, err := writer.Write(content); err != nil {
		log.Println(rendererTAG, "write:", err)
	}
}
