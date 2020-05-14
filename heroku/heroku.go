package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/PAWA-cl/goradbot"
)

//This process is only for heroku.
//The local function can be found in goradbot.go

func herokuLaunch() {
	//Heroku Stuff.
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)

	//Launch the radio server.
	goradbot.Start(token)
}
