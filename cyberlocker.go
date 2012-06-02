package main

import (
    "github.com/hoisie/web"
    "github.com/hoisie/mustache"
)

func gotUploadRequest(ctx *web.Context, val string) string { 
    _, hdr, _ := ctx.Request.FormFile("file")
    return hdr.Filename
}

func form(val string) string { 
    return mustache.RenderFile("views/index.mu")
}

func main() {
    web.Get("/", form)
    web.Post("/blah/upload", gotUploadRequest)
    web.Run("0.0.0.0:8080")
}
