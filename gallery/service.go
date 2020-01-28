package gallery

import "github.com/avavirus3/test2/entity"

// GalleryService specifies food menu gallery services
type GalleryService interface {
	Gallerys() ([]entity.Gallery, []error)
	Gallery(id uint) (*entity.Gallery, []error)
	UpdateGallery(gallery *entity.Gallery) (*entity.Gallery, []error)
	DeleteGallery(id uint) (*entity.Gallery, []error)
	StoreGallery(gallery *entity.Gallery) (*entity.Gallery, []error)
	PicsInGallery(gallery *entity.Gallery) ([]entity.Pic, []error)
}

// PicService specifies food menu pic related services
type PicService interface {
	Pics() ([]entity.Pic, []error)
	Pic(id uint) (*entity.Pic, []error)
	UpdatePic(menu *entity.Pic) (*entity.Pic, []error)
	DeletePic(id uint) (*entity.Pic, []error)
	StorePic(pic *entity.Pic) (*entity.Pic, []error)
}
