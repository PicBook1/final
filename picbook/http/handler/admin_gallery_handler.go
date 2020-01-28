package handler

import (
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/avavirus3/test2/entity"
	
"github.com/avavirus3/test2/form"
	"github.com/avavirus3/test2/gallery"
	"github.com/avavirus3/test2/rtoken"
)

// AdminGalleryHandler handles gallery handler admin requests
type AdminGalleryHandler struct {
	tmpl        *template.Template
	gallerySrv gallery.GalleryService
	csrfSignKey []byte
}

// NewAdminGalleryHandler initializes and returns new AdminCateogryHandler
func NewAdminGalleryHandler(t *template.Template, cs gallery.GalleryService, csKey []byte) *AdminGalleryHandler {
	return &AdminGalleryHandler{tmpl: t, gallerySrv: cs, csrfSignKey: csKey}
}

// AdminGallerys handle requests on route /admin/gallerys
func (ach *AdminGalleryHandler) AdminGallerys(w http.ResponseWriter, r *http.Request) {
	gallerys, errs := ach.gallerySrv.Gallerys()
	if errs != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
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
	ach.tmpl.ExecuteTemplate(w, "admin.categ.layout", tmplData)
}

// AdminGallerysNew hanlde requests on route /admin/gallerys/new
func (ach *AdminGalleryHandler) AdminGallerysNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		newCatForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		newCatForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		newCatForm.Required("catname", "catdesc")
		newCatForm.MinLength("catdesc", 10)
		newCatForm.CSRF = token
		// If there are any errors, redisplay the signup form.
		if !newCatForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
			return
		}
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			newCatForm.VErrors.Add("catimg", "File error")
			ach.tmpl.ExecuteTemplate(w, "admin.categ.new.layout", newCatForm)
			return
		}
		defer mf.Close()
		ctg := &entity.Gallery{
			Name:        r.FormValue("catname"),
			Description: r.FormValue("catdesc"),
			Image:       fh.Filename,
		}
		writeFile(&mf, fh.Filename)
		_, errs := ach.gallerySrv.StoreGallery(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/admin/gallerys", http.StatusSeeOther)
	}
}

// AdminGallerysUpdate handle requests on /admin/gallerys/update
func (ach *AdminGalleryHandler) AdminGallerysUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		cat, errs := ach.gallerySrv.Gallery(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		values := url.Values{}
		values.Add("catid", idRaw)
		values.Add("catname", cat.Name)
		values.Add("catdesc", cat.Description)
		values.Add("catimg", cat.Image)
		upCatForm := struct {
			Values   url.Values
			VErrors  form.ValidationErrors
			Gallery *entity.Gallery
			CSRF     string
		}{
			Values:   values,
			VErrors:  form.ValidationErrors{},
			Gallery: cat,
			CSRF:     token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.categ.update.layout", upCatForm)
		return
	}
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		updateCatForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		updateCatForm.Required("catname", "catdesc")
		updateCatForm.MinLength("catdesc", 10)
		updateCatForm.CSRF = token

		catID, err := strconv.Atoi(r.FormValue("catid"))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		ctg := &entity.Gallery{
			ID:          uint(catID),
			Name:        r.FormValue("catname"),
			Description: r.FormValue("catdesc"),
			Image:       r.FormValue("imgname"),
		}
		mf, fh, err := r.FormFile("catimg")
		if err == nil {
			ctg.Image = fh.Filename
			err = writeFile(&mf, ctg.Image)
		}
		if mf != nil {
			defer mf.Close()
		}
		_, errs := ach.gallerySrv.UpdateGallery(ctg)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/gallerys", http.StatusSeeOther)
		return
	}
}

// AdminGallerysDelete handle requests on route /admin/gallerys/delete
func (ach *AdminGalleryHandler) AdminGallerysDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := ach.gallerySrv.DeleteGallery(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/admin/gallerys", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, "ui", "assets", "img", fname)
	image, err := os.Create(path)
	if err != nil {
		return err
	}
	defer image.Close()
	io.Copy(image, *mf)
	return nil
}
