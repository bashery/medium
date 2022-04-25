package main

import (
	"database/sql"
	"io"
	"text/template"

	"fmt"
	"os"
	"os/exec"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Phon   string
	Avatar string
}

func main() {
	db := setdb()
	defer db.Close()
	e := echo.New()
	e.Renderer = t
	e.GET("/", personInfo)
	fmt.Println(e.Start(":8080"))
}

func personInfo(c echo.Context) (err error) {
	user := getUserInfo(1)
	fmt.Println(user)

	err = c.Render(200, "index.html", user) // map[string]User{"User": user})
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// gets all user information for update this info
func getUserInfo(userid int) (user User) {
	err := db.QueryRow(
		"SELECT username, email,phon, linkavatar FROM stores.users WHERE userid = ?",
		userid).Scan(&user.Name, &user.Email, &user.Phon, &user.Avatar)
	if err != nil {
		fmt.Println("no result or", err.Error())
	}
	return user
}

var (
	db  *sql.DB
	err error
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("templates/*.html")),
}

func setdb() *sql.DB {
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil { // why no error when db is not runinig ??
		fmt.Println("error when open mysql server", err)
		// TODO report this error.
		os.Exit(1)

	}

	if err = db.Ping(); err != nil {

		fmt.Println("error when ping to database", err)
		switch {
		case strings.Contains(err.Error(), "connection refused"):
			// TODO handle errors by code of error not by strings.

			cmd := exec.Command("mysql.server", "restart")
			// for systemd linux : exec.Command("sudo", "service", "mariadb", "start")
			//cmd.Stdin = strings.NewReader(os.Getenv("JAWAD"))
			errc := cmd.Run()
			if errc != nil {
				fmt.Println("error when run database cmd ", errc)
			}
		default:
			fmt.Println("error at  setdb() func, db.Ping() func")
			fmt.Println("unknown this error", err)
			os.Exit(1)
		}
	}
	return db
}
