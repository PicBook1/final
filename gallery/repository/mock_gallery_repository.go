package repository

import (
	"errors"

	"github.com/avavirus3/test2/entity"
	"github.com/avavirus3/test2/gallery"
	"github.com/jinzhu/gorm"
)

// MockGalleryRepo implements the gallery.GalleryRepository interface
type MockGalleryRepo struct {
	conn *gorm.DB
}

// NewMockGalleryRepo will create a new object of MockGalleryRepo
func NewMockGalleryRepo(db *gorm.DB) gallery.GalleryRepository {
	return &MockGalleryRepo{conn: db}
}

// Gallery returns all fake gallery
func (mCatRepo *MockGalleryRepo) Gallerys() ([]entity.Gallery, []error) {
	ctgs := []entity.Gallery{entity.GalleryMock}
	return ctgs, nil
}

// Gallery retrieve a fake gallery with id 1
func (mCatRepo *MockGalleryRepo) Gallery(id uint) (*entity.Gallery, []error) {
	ctg := entity.GalleryMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateGallery updates a given fake gallery
func (mCatRepo *MockGalleryRepo) UpdateGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {
	cat := entity.GalleryMock
	return &cat, nil
}

// DeleteGallery deletes a given gallery from the database
func (mCatRepo *MockGalleryRepo) DeleteGallery(id uint) (*entity.Gallery, []error) {
	cat := entity.GalleryMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

// StoreGallery stores a given mock gallery
func (mCatRepo *MockGalleryRepo) StoreGallery(gallery *entity.Gallery) (*entity.Gallery, []error) {
	cat := gallery
	return cat, nil
}

// PicsInGallery returns mock food gallery pics
func (mCatRepo *MockGalleryRepo) PicsInGallery(gallery *entity.Gallery) ([]entity.Pic, []error) {
	pics := []entity.Pic{entity.PicMock}
	return pics, nil
}
