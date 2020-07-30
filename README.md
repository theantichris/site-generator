# site-generator

A static site generator written in Go.

## Dependencies

* [Gin](https://github.com/gin-gonic/gin)

## Use

### Setup

1. Remove `.gitkeep` files from `markdown` and `templates`
1. Add the following templates to `templates`
    * `index.tmpl.html` - front page and list of all posts
    * `post.tmpl.html`  - individual post page
    * `error.tmpl.html` - error page
