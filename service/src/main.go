package main

import (
	"fmt"
	"html/template"
	"os"

	// "github.com/joho/godotenv"

	"service/src/env"
	"service/src/libs/mongodb"
	"service/src/routes/auth/stage0"
	"service/src/routes/auth/stage1"
	"service/src/routes/auth/stage2"
	"service/src/routes/auth/user"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var icons = map[string]string{
	"google": "/static/google.png",
	"do":     "/static/do.png",
	"github": "/static/github.png",
}

func main() {
	globalEnv := env.Env{
		Mongo: mongodb.Connect(),
	}

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"capitalize": func(str string) string {
			return cases.Title(language.English).String(str)
		},
		"getIcon": func(provider string) string {
			return icons[provider]
		},
	})

	// })
	r.Static("/static", "./service/src/static")
	r.LoadHTMLFiles("./service/src/routes/auth/stage0/stage0.html")

	r.GET("/auth/:app/stage0", stage0.CreateHandler(&globalEnv))
	r.GET("/auth/:app/stage1/:provider", stage1.CreateHandler(&globalEnv))
	r.GET("/auth/:app/stage2/:provider", stage2.CreateHandler(&globalEnv))
	r.GET("/auth/user", user.CreateHandler(&globalEnv))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
