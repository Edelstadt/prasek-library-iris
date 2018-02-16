package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"time"
)

func main() {
	app := iris.New()
	app.StaticServe("./templates/images", "/images")
	app.StaticServe("./templates/", "/")
	app.RegisterView(iris.Django("./templates", ".html"))

	app.Get("/", cache.Handler(10*time.Second), writeMarkdown)
	app.Get("/o-nas", cache.Handler(10*time.Second), writeMarkdownOnas)
	app.Get("/akce", cache.Handler(10*time.Second), writeMarkdownAkce)
	app.Get("/o-knihovne", cache.Handler(10*time.Second), writeMarkdownOknihovne)
	app.Get("/sluzby", cache.Handler(10*time.Second), writeMarkdownSluzby)
	app.Get("/zajimavosti", cache.Handler(10*time.Second), writeMarkdownZajimavosti)
	app.Get("/podekovani", cache.Handler(10*time.Second), writeMarkdownPodekovani)
	app.Get("/kontakt", cache.Handler(10*time.Second), writeMarkdownKontakt)

	// Start the server using a network address.
	app.Run(iris.Addr(":8080"))
}

func writeMarkdownKontakt(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_kontakt.html")
}

func writeMarkdownPodekovani(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_podekovani.html")
}

func writeMarkdownZajimavosti(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_zajimavosti.html")
}

func writeMarkdown(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_index.html")
}

func writeMarkdownOnas(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_o_nas.html")
}

func writeMarkdownAkce(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_akce.html")
}

func writeMarkdownOknihovne(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_o_knihovne.html")
}

func writeMarkdownSluzby(ctx iris.Context) {
	println("Handler executed. Content refreshed.")
	ctx.View("test_sluzby.html")
}
