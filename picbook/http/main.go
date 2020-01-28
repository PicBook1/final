package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/avavirus3/test2/Picbook/http/handler"
	"github.com/avavirus3/test2/entity"
	mrepim "github.com/avavirus3/test2/gallery/repository"
	msrvim "github.com/avavirus3/test2/gallery/service"
	"github.com/avavirus3/test2/rtoken"

	urepimp "github.com/avavirus3/test2/user/repository"
	usrvimp "github.com/avavirus3/test2/user/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createTables(dbconn *gorm.DB) []error {
	errs := dbconn.CreateTable(&entity.User{}, &entity.Role{}, &entity.Session{}, &entity.Pic{}, &entity.Gallery{}, &entity.Comment{}).GetErrors()
	if errs != nil {
		return errs
	}
	return nil
}
const (
	// Name of the database.
	user = "postgres"
	DBName          = "webproject"
	password = "+"
)
func main() {
	//createTables(dbconn)

	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	tmpl := template.Must(template.ParseGlob("ui/templates/*.html"))

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, DBName)
	dbconn, err := gorm.Open("postgres", connStr)
	//createTables(dbconn)
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	sessionRepo := urepimp.NewSessionGormRepo(dbconn)
	sessionSrv := usrvimp.NewSessionService(sessionRepo)

	galleryRepo := mrepim.NewGalleryGormRepo(dbconn)
	galleryServ := msrvim.NewGalleryService(galleryRepo)

	userRepo := urepimp.NewUserGormRepo(dbconn)
	userServ := usrvimp.NewUserService(userRepo)

	roleRepo := urepimp.NewRoleGormRepo(dbconn)
	roleServ := usrvimp.NewRoleService(roleRepo)

	ach := handler.NewAdminGalleryHandler(tmpl, galleryServ, csrfSignKey)
	mh := handler.NewGalleryHandler(tmpl, galleryServ, csrfSignKey)

	sess := configSess()
	uh := handler.NewUserHandler(tmpl, userServ, sessionSrv, roleServ, sess, csrfSignKey)

	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", mh.Index)
	http.HandleFunc("/about", mh.About)
	http.HandleFunc("/contact", mh.Contact)
	http.HandleFunc("/gallery", mh.Gallery)
	http.Handle("/admin", uh.Authenticated(uh.Authorized(http.HandlerFunc(mh.Admin))))

	http.Handle("/admin/gallerys", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminGallerys))))
	http.Handle("/admin/gallerys/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminGallerysNew))))
	http.Handle("/admin/gallerys/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminGallerysUpdate))))
	http.Handle("/admin/gallerys/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(ach.AdminGallerysDelete))))

	http.Handle("/admin/users", uh.Authenticated(uh.Authorized(http.HandlerFunc(uh.AdminUsers))))
	http.Handle("/admin/users/new", uh.Authenticated(uh.Authorized(http.HandlerFunc(uh.AdminUsersNew))))
	http.Handle("/admin/users/update", uh.Authenticated(uh.Authorized(http.HandlerFunc(uh.AdminUsersUpdate))))
	http.Handle("/admin/users/delete", uh.Authenticated(uh.Authorized(http.HandlerFunc(uh.AdminUsersDelete))))

	http.HandleFunc("/login", uh.Login)
	http.Handle("/logout", uh.Authenticated(http.HandlerFunc(uh.Logout)))
	http.HandleFunc("/signup", uh.Signup)

	//port := fmt.Sprintf(":%s", os.Getenv("HPORT"))

	http.ListenAndServe(":8000", nil)
}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}
