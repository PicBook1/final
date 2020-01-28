package handler

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/form"
	"github.com/avavirus3/test2/gallery"
	"github.com/avavirus3/test2/rtoken"
)

// GalleryHandler handles gallery related requests
type GalleryHandler struct {
	tmpl        *template.Template
	gallerySrv gallery.GalleryService
	csrfSignKey []byte
}

// NewgalleryHandler initializes and returns new galleryHandler
func NewGalleryHandler(T *template.Template, CS gallery.GalleryService, csKey []byte) *GalleryHandler {
	return &GalleryHandler{tmpl: T, gallerySrv: CS, csrfSignKey: csKey}
}

// Index handles request on route /
func (mh *GalleryHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	gallerys, errs := mh.gallerySrv.Gallerys()
	if len(errs) > 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values     url.Values
		VErrors    form.ValidationErrors
		Gallerys []entity.Gallery
		CSRF       string
	}{
		Values:     nil,
		VErrors:    nil,
		Gallerys: gallerys,
		CSRF:       token,
	}

	mh.tmpl.ExecuteTemplate(w, "index.layout", tmplData)
}

// About handles requests on route /about
func (mh *GalleryHandler) About(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "about.layout", tmplData)
}

// Gallery handle request on route /gallery
func (mh *GalleryHandler) Gallery(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "gallery.layout", tmplData)
}

// Contact handle request on route /Contact
func (mh *GalleryHandler) Contact(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "contact.layout", tmplData)
}

// Admin handle request on route /admin
func (mh *GalleryHandler) Admin(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(mh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	tmplData := struct {
		Values  url.Values
		VErrors form.ValidationErrors
		CSRF    string
	}{
		Values:  nil,
		VErrors: nil,
		CSRF:    token,
	}
	mh.tmpl.ExecuteTemplate(w, "admin.index.layout", tmplData)
}
