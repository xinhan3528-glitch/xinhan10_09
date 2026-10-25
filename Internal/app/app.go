package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"xinhan10_09/Internal/api"
	"xinhan10_09/Internal/store"
	"xinhan10_09/migrations"
)

type Application struct {
	Logger      *log.Logger
	PostHandler *api.PostHandler
	DB          *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	postHandler := api.NewPostHandler()
	postgresDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(postgresDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	app := &Application{
		Logger:      logger,
		PostHandler: postHandler,
		DB:          postgresDB,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "状况正常\n")
}
