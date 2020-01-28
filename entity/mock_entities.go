package entity

import "time"

// GalleryMock mocks Food Menu Gallery
var GalleryMock = Gallery{
	ID:          1,
	Name:        "Mock Gallery 01",
	Description: "Mock Gallery 01 Description",
	Image:       "mock_cat.png",
	Pic:       []Pic{},
}

// RoleMock mocks user role entity
var RoleMock = Role{
	ID:    1,
	Name:  "Mock Role 01",
	Users: []User{},
}

// ItemMock mocks food menu pics
var PicMock = Pic{
	ID:          1,
	Name:        "Mock Item 01",
	Description: "Mock Item 01 Description",
	Gallerys:  []Gallery{},
	Image:       "mock_pic.png",
	}



// UserMock mocks application user
var UserMock = User{
	ID:       1,
	FullName: "Mock User 01",
	Email:    "mockuser@example.com",
	Password: "P@$$w0rd",
	RoleID:   1,

}

// SessionMock mocks sessions of loged in user
var SessionMock = Session{
	ID:         1,
	UUID:       "_session_one",
	SigningKey: []byte("RestaurantApp"),
	Expires:    0,
}

// CommentMock mocks comments forwarded by application users
var CommentMock = Comment{
	ID:        1,
	FullName:  "Mock User 01",
	Message:   "Mock message",
	Email:     "mockuser@example.com",
	CreatedAt: time.Time{},
}