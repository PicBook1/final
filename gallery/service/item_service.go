package service

import (
	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery"
)

// PicService implements gallery.PicService interface
type PicService struct {
	picRepo gallery.PicRepository
}

// NewPicService returns new PicService object
func NewPicService(picRepository gallery.PicRepository) gallery.PicService {
	return &PicService{picRepo: picRepository}
}

// Pics returns all stored food gallery pics
func (is *PicService) Pics() ([]entity.Pic, []error) {
	itms, errs := is.picRepo.Pics()
	if len(errs) > 0 {
		return nil, errs
	}
	return itms, errs
}

// Pic retrieves a food gallery pic by its id
func (is *PicService) Pic(id uint) (*entity.Pic, []error) {
	itm, errs := is.picRepo.Pic(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// UpdatePic updates a given food gallery pic
func (is *PicService) UpdatePic(pic *entity.Pic) (*entity.Pic, []error) {
	itm, errs := is.picRepo.UpdatePic(pic)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeletePic deletes a given food gallery pic
func (is *PicService) DeletePic(id uint) (*entity.Pic, []error) {
	itm, errs := is.picRepo.DeletePic(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StorePic stores a given food gallery pic
func (is *PicService) StorePic(pic *entity.Pic) (*entity.Pic, []error) {
	itm, errs := is.picRepo.StorePic(pic)
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
