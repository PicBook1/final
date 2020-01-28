package repository

import (
	"database/sql"
	"errors"

	"github.com/avavirus3/test2/entity"
)

// GalleryRepositoryImpl implements the gallery.GalleryRepository interface
type GalleryRepositoryImpl struct {
	conn *sql.DB
}

// NewGalleryRepositoryImpl will create an object of PsqlGalleryRepository
func NewGalleryRepositoryImpl(Conn *sql.DB) *GalleryRepositoryImpl {
	return &GalleryRepositoryImpl{conn: Conn}
}

// Gallerys returns all cateogories from the database
func (cri *GalleryRepositoryImpl) Gallerys() ([]entity.Gallery, error) {

	rows, err := cri.conn.Query("SELECT * FROM gallerys;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Gallery{}

	for rows.Next() {
		gallery := entity.Gallery{}
		err = rows.Scan(&gallery.ID, &gallery.Name, &gallery.Description, &gallery.Image)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, gallery)
	}

	return ctgs, nil
}

// Gallery returns a gallery with a given id
func (cri *GalleryRepositoryImpl) Gallery(id uint) (entity.Gallery, error) {

	row := cri.conn.QueryRow("SELECT * FROM gallerys WHERE id = $1", id)

	c := entity.Gallery{}

	err := row.Scan(&c.ID, &c.Name, &c.Description, &c.Image)
	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateGallery updates a given object with a new data
func (cri *GalleryRepositoryImpl) UpdateGallery(c entity.Gallery) error {

	_, err := cri.conn.Exec("UPDATE gallerys SET name=$1,description=$2, image=$3 WHERE id=$4", c.Name, c.Description, c.Image, c.ID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteGallery removes a gallery from a database by its id
func (cri *GalleryRepositoryImpl) DeleteGallery(id uint) error {

	_, err := cri.conn.Exec("DELETE FROM gallerys WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreGallery stores new gallery information to database
func (cri *GalleryRepositoryImpl) StoreGallery(c entity.Gallery) error {

	_, err := cri.conn.Exec("INSERT INTO gallerys (name,description,image) values($1, $2, $3)", c.Name, c.Description, c.Image)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
