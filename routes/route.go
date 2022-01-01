package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/get-words/app/controllers"
)

// Routes ルート情報を設定
func Routes(engine *gin.Engine) {
	engine.GET("/get-words", indexHandler)
	engine.GET("/get-words/pokemon", pokemonHandler)
}

func indexHandler(context *gin.Context) {
	word, _ := controllers.GetWord()
	context.String(http.StatusOK, word)
}

func pokemonHandler(context *gin.Context) {
	pokemon, _ := controllers.GetPokemon()
	context.String(http.StatusOK, pokemon)
}
