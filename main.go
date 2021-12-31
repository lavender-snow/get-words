package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lavender-snow/get-words/config"
	"github.com/lavender-snow/get-words/routes"
	"github.com/lavender-snow/get-words/utils"
)

func main() {
	// ログ設定
	utils.LoggingSettings("musicScore")

	// gin初期化
	engine := gin.Default()
	store := sessions.NewCookieStore([]byte(config.Config.SecretKey))
	store.Options(sessions.Options{MaxAge: 3600})
	engine.Use(sessions.Sessions("getwords", store))
	engine.Static("/assets", "./assets")
	routes.Routes(engine)

	// 開始
	if config.Config.RunMode == "https" {
		engine.RunTLS(config.Config.PortNo, config.Config.CertFile, config.Config.KeyFile)
	} else {
		engine.Run(config.Config.PortNo)
	}

}
