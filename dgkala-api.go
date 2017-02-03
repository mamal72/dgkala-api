package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mamal72/dgkala"
	"gopkg.in/kataras/iris.v4"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func incredibleOffersHandler(ctx *iris.Context) {
	offers, err := dgkala.IncredibleOffers()
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}

	ctx.JSON(iris.StatusOK, offers)
}

func searchProductHandler(ctx *iris.Context) {
	keyword := ctx.Param("keyword")
	result, err := dgkala.Search(keyword)
	if err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
		return
	}
	if result.Count == 0 {
		ctx.EmitError(iris.StatusNotFound)
		return
	}

	ctx.JSON(iris.StatusOK, result)
}

func main() {
	godotenv.Load()
	iris.Get("/search/:keyword", searchProductHandler)
	iris.Get("/incredible-offers", incredibleOffersHandler)
	iris.Listen(fmt.Sprintf("%s:%s", getenv("HOST", "0.0.0.0"), getenv("PORT", "8080")))
}
