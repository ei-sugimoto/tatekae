package cmd

import (
	"github.com/ei-sugimoto/tatekae/api/web"
)

func ServeHTTP() {
	w := web.NewRouter()
	w.Run()
}
