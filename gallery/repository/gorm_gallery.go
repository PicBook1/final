package repository

import (
	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery"
	"github.com/jinzhu/gorm"
)

// GalleryGormRepo implements the gallery.GalleryRepository interface
type GalleryGormRepo struct {
	conn *gorm.DB
}

// NewGalleryGormRepo will create a new object of GalleryGormRepo
func NewGalleryGormRepo(db *gorm.DB) gallery.GalleryRepository {
	
	return &GalleryGormRepo{conn: db}
}

// Gallerys returns all gallerys stored in the database
func (cRepo *GalleryGormRepo) Gallerys() ([]entity.Gallery, []error) {
	ctgs := []entity.Gallery{}
	errs := cRepo.conn.Find(&ctgs).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ctgs, errs
}

// Gallery retrieve a gallery from the database by its id
func (cRepo *GalleryGormRepo) Gallery(id uint) (*entity.Gallery, []error) {
	ctg := entity.Gallery{}
	errs := cRepo.conn.First(&ctg, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &ctg, errs
}

// UpdateGallery updates a given gallery in the database
func (cRepo *GalleryGormRepo) UpdateGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {
	cat := gallery
	errs := cRepo.conn.Save(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// DeleteGallery deletes a given gallery from the database
func (cRepo *GalleryGormRepo) DeleteGallery(id uint) (*entity.Gallery, []error) {
	cat, errs := cRepo.Gallery(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = cRepo.conn.Delete(cat, cat.ID).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// StoreGallery stores a given gallery in the database
func (cRepo *GalleryGormRepo) StoreGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {
	cat := gallery
	errs := cRepo.conn.Create(cat).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cat, errs
}

// PicsInGallery retrive from a database a list of food pic gallerys from a given gallery
func (cRepo *GalleryGormRepo) PicsInGallery(gallery *entity.Gallery) ([]entity.Pic, []error) {
	pics := []entity.Pic{}
	cat, errs := cRepo.Gallery(gallery.ID)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cRepo.conn.Model(cat).Related(&pics, "Pics").GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pics, errs
}
