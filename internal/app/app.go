package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"training/internal/api"
	"training/internal/store"
	"training/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	workoutStore := store.NewPostgresWorkoutStor(nil)

	wh := api.NewWorkoutHandler(workoutStore)

	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	app := &Application{
		Logger:         logger,
		WorkoutHandler: wh,
		DB:             pgDB,
	}
	return app, nil
}

func (a *Application) Healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status is available\n")
}
