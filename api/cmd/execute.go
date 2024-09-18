package cmd

import (
	"github.com/ei-sugimoto/tatekae/api/infrastructure"
	"github.com/ei-sugimoto/tatekae/api/web"
)

func ServeHTTP() {
	db := infrastructure.NewDB()
	db.Migrate()
	w := web.NewRouter()
	w.Run()
}
