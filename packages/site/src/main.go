package main

// sd
import (
	"html/template"

	"site/src/env"
	"site/src/libs/mongodb"
	"site/src/routes/auth/user"
	"site/src/routes/auth/client"
	"site/src/routes/auth/client/start"
	"site/src/routes/auth/client/cb"

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

	r.Static("/static", "./packages/site/src/static")
	r.LoadHTMLFiles("./packages/site/src/routes/auth/client/client.html")


	r.GET("/auth/:client", client.CreateHandler(&globalEnv))
	r.GET("/auth/:client/:provider", start.CreateHandler(&globalEnv))
	r.GET("/auth/:client/:provider/cb", cb.CreateHandler(&globalEnv))
	
	r.GET("/auth/user", user.CreateHandler(&globalEnv))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
