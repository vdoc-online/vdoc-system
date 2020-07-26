package web

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//Start web service
func Start(port int) {
	r := gin.Default()
	r.LoadHTMLGlob("assets/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/vdoc", func(c *gin.Context) {
		url := c.Query("url")
		url = strings.Trim(url, " ")
		if url == "" {
			c.JSON(http.StatusOK, Result{Code: 1, Msg: "empty url param"})
			return
		}
		dataTyppe := c.Query("type")
		vdocs := loadVdocs(c)
		switch dataTyppe {
		case "json":
			c.JSON(http.StatusOK, OK(vdocs))
		default:
			c.HTML(http.StatusOK, "vdoc.tmpl", gin.H{
				"vdocs": vdocs,
			})
		}
	})
	r.POST("/submit", func(c *gin.Context) {
		url := c.PostForm("url")
		vurl := c.PostForm("vurl")
		if url == "" || vurl == "" {
			c.JSON(http.StatusOK, Result{Code: 1, Msg: "url and vurl is empty"})
			return
		}

	})
	r.Run(":" + strconv.Itoa(port))

}

func loadVdocs(c *gin.Context) []*Vdoc {
	vdocs := make([]*Vdoc, 0)
	return vdocs
}

func submitVdoc(c *gin.Context) {

}
