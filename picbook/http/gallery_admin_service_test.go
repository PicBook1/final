package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/avavirus3/test2/picbook/http/handler"
	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery/repository"
	"github.com/avavirus3/test2/gallery/service"
)

func TestAdminGallerys(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	galleryRepo := repository.NewMockGalleryRepo(nil)
	galleryServ := service.NewGalleryService(galleryRepo)

	adminCatgHandler := handler.NewAdminGalleryHandler(tmpl, galleryServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/gallerys", adminCatgHandler.AdminGallerys)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/gallerys")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Gallery 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminGallerysNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	categoryRepo := repository.NewMockGalleryRepo(nil)
	categoryServ := service.NewGalleryService(categoryRepo)

	adminCatgHandler := handler.NewAdminGalleryHandler(tmpl, categoryServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/gallerys/new", adminCatgHandler.AdminGallerys)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("name", entity.GalleryMock.Name)
	form.Add("Description", entity.GalleryMock.Description)
	form.Add("Image", entity.GalleryMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/gallerys/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Gallery 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminCategoresUpdate(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	galleryRepo := repository.NewMockGalleryRepo(nil)
	galleryServ := service.NewGalleryService(galleryRepo)

	adminCatgHandler := handler.NewAdminGalleryHandler(tmpl, galleryServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/gallerys/update", adminCatgHandler.AdminGallerys)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("ID", string(entity.GalleryMock.ID))
	form.Add("Name", entity.GalleryMock.Name)
	form.Add("kescription", entity.GalleryMock.Description)
	form.Add("Image", entity.GalleryMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/gallerys/update?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Gallery 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminCategoresDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	galleryRepo := repository.NewMockGalleryRepo(nil)
	galleryServ := service.NewGalleryService(galleryRepo)

	adminCatgHandler := handler.NewAdminGalleryHandler(tmpl, galleryServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/gallerys/delete", adminCatgHandler.AdminGallerys)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("ID", string(entity.GalleryMock.ID))
	form.Add("Name", entity.GalleryMock.Name)
	form.Add("kescription", entity.GalleryMock.Description)
	form.Add("Image", entity.GalleryMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/gallerys/delete?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock Gallery 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
