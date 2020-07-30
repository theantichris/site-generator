# site-generator

A static site generator written in Go.

## Dependencies

* [Gin](https://github.com/gin-gonic/gin)

## Use

### Setup

1. Remove `.gitkeep` file from `markdown`
1. Customize the following templates in `templates`
    * `index.tmpl.html` - front page and list of all posts
    * `post.tmpl.html`  - individual post page
    * `error.tmpl.html` - error page
