package repository

import (
	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery"
	"github.com/jinzhu/gorm"
)

// PicGormRepo implements the gallery.PicRepository interface
type PicGormRepo struct {
	conn *gorm.DB
}

// NewPicGormRepo will create a new object of PicGormRepo
func NewPicGormRepo(db *gorm.DB) gallery.PicRepository {
	return &PicGormRepo{conn: db}
}

// Pics returns all food gallerys stored in the database
func (picRepo *PicGormRepo) Pics() ([]entity.Pic, []error) {
	pics := []entity.Pic{}
	errs := picRepo.conn.Find(&pics).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pics, errs
}

// Pic retrieves a food gallery by its id from the database
func (picRepo *PicGormRepo) Pic(id uint) (*entity.Pic, []error) {
	pic := entity.Pic{}
	errs := picRepo.conn.First(&pic, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &pic, errs
}

// UpdatePic updates a given food gallery pic in the database
func (picRepo *PicGormRepo) UpdatePic(pic *entity.Pic) (*entity.Pic, []error) {
	itm := pic
	errs := picRepo.conn.Save(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// DeletePic deletes a given food gallery pic from the database
func (picRepo *PicGormRepo) DeletePic(id uint) (*entity.Pic, []error) {
	itm, errs := picRepo.Pic(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = picRepo.conn.Delete(itm, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}

// StorePic stores a given food gallery pic in the database
func (picRepo *PicGormRepo) StorePic(pic *entity.Pic) (*entity.Pic, []error) {
	itm := pic
	errs := picRepo.conn.Create(itm).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return itm, errs
}
