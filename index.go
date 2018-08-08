package main

import (
	"database/sql"
	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)
import f "landing/sql"

type user struct {
	id       int
	name     string
	password string
}

func main() {
	app := martini.Classic()
	connStr := "host=sharko21.db user=meuser password=0929598d5f0b2faea7b85a06a366d9aa dbname=medb sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {
		if err != nil {
			return "Db is not connected"
		}

		req.ParseForm()
		if req.Form["login"] != nil && req.Form["password"] != nil {
			user := user{}

			DbUser := db.QueryRow(f.Sql["user"]["getUser"], req.Form["login"][0])
			DbUser.Scan(&user.id, &user.name, &user.password)

			if bcrypt.CompareHashAndPassword([]byte(user.password), []byte(req.Form["password"][0])) == nil {
				return "Loggined"
			} else {
				return "CHECK YOUR LOGIN AND PASSWORD!"
			}
		}
		return "Enter login and password in GET params!"
	})

	app.RunOnAddr(":8000")
}
