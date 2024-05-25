package render

import (
	"log"
	"net/http"

	"github.com/mo0Oonnn/bookings/internal/config"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func SetConfig(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td jet.VarMap, r *http.Request) jet.VarMap {
	td.Set("success", app.Session.PopString(r.Context(), "success"))
	td.Set("error", app.Session.PopString(r.Context(), "error"))
	td.Set("warning", app.Session.PopString(r.Context(), "warning"))

	td.Set("csrf_token", nosurf.Token(r))
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td jet.VarMap) {
	var tc *jet.Set

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc = CreateTemplateCache()
	}

	t, err := tc.GetTemplate(tmpl)
	if err != nil {
		log.Fatal("cannot get template from template cache: ", err)
	}

	td = AddDefaultData(td, r)

	err = t.Execute(w, td, nil)
	if err != nil {
		log.Println("error writing template to browser:", err)
	}
}

func CreateTemplateCache() *jet.Set {
	views := jet.NewHTMLSet("./templates")
	return views
}
