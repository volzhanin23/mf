package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Nick struct {
	Nickname string `form:"nick"`
}

type Setting struct {
	Setting []string
}

var players []string

var setting Setting

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		log.Print(setting)
		if len(players) > 9 && setting.Setting == nil {
			setting.Setting = Shuffle(players)
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Users":   players,
			"Setting": setting.Setting,
		})
	})

	r.GET("/reset", func(c *gin.Context) {
		players = nil
		setting.Setting = nil

		c.Redirect(302, "/")
	})

	r.POST("/", func(c *gin.Context) {
		var nick Nick
		c.Bind(&nick)
		if len(players) > 9 {
			c.HTML(http.StatusOK, "error.tmpl", gin.H{
				"message": "Набрано нужное количество игроков",
			})

			return
		}

		players = append(players, nick.Nickname)

		c.Redirect(302, "/")
	})

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Shuffle(players []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(players))
	perm := r.Perm(len(players))
	for i, randIndex := range perm {
		ret[i] = players[randIndex]
	}
	return ret
}
