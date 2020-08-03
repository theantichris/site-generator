package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// TODO: load from env
const (
	assetsFolder   = "/assets"
	markdownFolder = "/markdown"
	templateGlob   = "/templates/*.tmpl.html"
	indexTemplate  = "index.tmpl.html"
	postTemplate   = "post.tmpl.html"
	errorTemplate  = "error.tmpl.html"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}") // sets template tags

	r.Use(static.Serve(assetsFolder, static.LocalFile("/assets", false)))
	r.LoadHTMLGlob("." + templateGlob)

	r.GET("/", func(c *gin.Context) {
		var posts []string

		files, err := ioutil.ReadDir("." + markdownFolder)
		if err != nil {
			log.Fatalln(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}

		c.HTML(http.StatusOK, indexTemplate, gin.H{"posts": posts})
	})

	r.GET("/:postTitle", func(c *gin.Context) {
		postTitle := c.Param("postTitle")

		mdFile, err := ioutil.ReadFile("." + markdownFolder + "/" + postTitle)
		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusNotFound, errorTemplate, nil)
			c.Abort()
			return
		}

		postContent := template.HTML(blackfriday.MarkdownCommon([]byte(mdFile)))

		c.HTML(http.StatusOK, postTemplate, gin.H{
			"Title":   postTitle,
			"Content": postContent,
		})
	})

	r.Run(":8080")
}
