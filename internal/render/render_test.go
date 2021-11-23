package render

import (
	"net/http"
	"rarebnb/internal/models"
	"testing"
)

func TestAddDefaultData(t *testing.T){
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Failed to add default data(123).")
	}
}

func TestRenderTemplate(t *testing.T){
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error writing template to browser.")
	}

	err = RenderTemplate(&ww, r, "nonce.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Retrieved nonexistent template.")
	}
} 

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
} 

func TestNewTemplates(t *testing.T){
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T){
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}