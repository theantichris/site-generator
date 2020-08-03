package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// TODO: create struct or map, return in loadEnv()
var (
	port          string
	assetsDir     string
	markdownDir   string
	templateGlob  string
	indexTemplate string
	postTemplate  string
	errorTemplate string
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}") // sets template tags

	r.Use(static.Serve(assetsDir, static.LocalFile(assetsDir, false)))
	r.LoadHTMLGlob("." + templateGlob)

	r.GET("/", func(c *gin.Context) {
		var posts []string

		files, err := ioutil.ReadDir("." + markdownDir)
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

		mdFile, err := ioutil.ReadFile("." + markdownDir + "/" + postTitle)
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

	r.Run(port) // TODO: look for interrupt
}

func loadEnv() error {
	err := godotenv.Load()
	if err == nil {
		port = os.Getenv("PORT")
		assetsDir = os.Getenv("ASSETS_DIR")
		markdownDir = os.Getenv("MARKDOWN_DIR")
		templateGlob = os.Getenv("TEMPLATE_GLOB")
		indexTemplate = os.Getenv("INDEX_TEMPLATE")
		postTemplate = os.Getenv("POST_TEMPLATE")
		errorTemplate = os.Getenv("ERROR_TEMPLATE")
	}

	return err
}
