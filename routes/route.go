package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/get-words/app/controllers"
)

// Routes ルート情報を設定
func Routes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("/get-words", indexHandler)
	}
}

func indexHandler(context *gin.Context) {
	word, _ := controllers.GetWord()
	context.String(http.StatusOK, word)
}
