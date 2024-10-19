package handler

import (
	"net/http"

	"github.com/furmedia/bookings/config"
	"github.com/furmedia/bookings/pkg/models"
	"github.com/furmedia/bookings/pkg/render"
)

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandler sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.template.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.template.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
