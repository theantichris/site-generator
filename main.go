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

type options struct {
	port          string
	assetsDir     string
	markdownDir   string
	templateGlob  string
	indexTemplate string
	postTemplate  string
	errorTemplate string
}

func main() {
	options, err := loadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}") // sets template tags

	r.Use(static.Serve(options.assetsDir, static.LocalFile(options.assetsDir, false)))
	r.LoadHTMLGlob("." + options.templateGlob)

	r.GET("/", func(c *gin.Context) {
		var posts []string

		files, err := ioutil.ReadDir("." + options.markdownDir)
		if err != nil {
			log.Fatalln(err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
			posts = append(posts, file.Name())
		}

		c.HTML(http.StatusOK, options.indexTemplate, gin.H{"posts": posts})
	})

	r.GET("/:postTitle", func(c *gin.Context) {
		postTitle := c.Param("postTitle")

		mdFile, err := ioutil.ReadFile("." + options.markdownDir + "/" + postTitle)
		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusNotFound, options.errorTemplate, nil)
			c.Abort()
			return
		}

		postContent := template.HTML(blackfriday.MarkdownCommon([]byte(mdFile)))

		c.HTML(http.StatusOK, options.postTemplate, gin.H{
			"Title":   postTitle,
			"Content": postContent,
		})
	})

	r.Run(options.port) // TODO: look for interrupt
}

func loadEnv() (options, error) {
	if err := godotenv.Load(); err != nil {
		return options{}, err
	}

	options := options{
		port:          os.Getenv("PORT"),
		assetsDir:     os.Getenv("ASSETS_DIR"),
		markdownDir:   os.Getenv("MARKDOWN_DIR"),
		templateGlob:  os.Getenv("TEMPLATE_GLOB"),
		indexTemplate: os.Getenv("INDEX_TEMPLATE"),
		postTemplate:  os.Getenv("POST_TEMPLATE"),
		errorTemplate: os.Getenv("ERROR_TEMPLATE"),
	}

	return options, nil
}
