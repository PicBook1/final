package service

import (
	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery"
)

// GalleryService implements gallery.GalleryService interface
type GalleryService struct {
	galleryRepo gallery.GalleryRepository
}

// NewGalleryService will create new GalleryService object
func NewGalleryService(CatRepo gallery.GalleryRepository) gallery.GalleryService {
	
	return &GalleryService{galleryRepo: CatRepo}
}

// Gallerys returns list of gallerys
func (cs *GalleryService) Gallerys() ([]entity.Gallery, []error) {

	gallerys, errs := cs.galleryRepo.Gallerys()

	if len(errs) > 0 {
		return nil, errs
	}

	return gallerys, nil
}

// StoreGallery persists new gallery information
func (cs *GalleryService) StoreGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {

	cat, errs := cs.galleryRepo.StoreGallery(gallery)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// Gallery returns a gallery object with a given id
func (cs *GalleryService) Gallery(id uint) (*entity.Gallery, []error) {

	c, errs := cs.galleryRepo.Gallery(id)

	if len(errs) > 0 {
		return c, errs
	}

	return c, nil
}

// UpdateGallery updates a cateogory with new data
func (cs *GalleryService) UpdateGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {

	cat, errs := cs.galleryRepo.UpdateGallery(gallery)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// DeleteGallery delete a gallery by its id
func (cs *GalleryService) DeleteGallery(id uint) (*entity.Gallery, []error) {

	cat, errs := cs.galleryRepo.DeleteGallery(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return cat, nil
}

// PicsInGallery returns list of gallery items in a given gallery
func (cs *GalleryService) PicsInGallery(gallery *entity.Gallery) ([]entity.Pic, []error) {

	cts, errs := cs.galleryRepo.PicsInGallery(gallery)

	if len(errs) > 0 {
		return nil, errs
	}

	return cts, nil

}
