# site-generator

A static site generator written in Go.

## Dependencies

* [Godotenv](https://github.com/joho/godotenv)
* [Gin](https://github.com/gin-gonic/gin)
* [static](https://github.com/gin-contrib/static)
* [Blackfriday](https://github.com/russross/blackfriday)

## Use

### Setup

1. Copy `.example.env` to `.env`
1. Customize `.env` variables if needed
    * `port`: server port number
    * `assetsDir`: directory contain CSS, images, etc.
    * `markdownDir`: directory containing site posts as markdown files
    * `templateGlob`: pattern to match for HTML templates
    * `indexTemplate`: file name of index template
    * `postTemplate`: file name of post template
    * `errorTemplate`: file name of error template
1. Customize the following templates in `templates` if needed
    * `index.tmpl.html`: front page and list of all posts
    * `post.tmpl.html`: individual post page
    * `error.tmpl.html`: error page
1. Run with `go run main.go`
