package main

import (
	"context"
	"fmt"
	"gkgkgkgk/internal/repository"
	"gkgkgkgk/internal/repository/memory"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: NewRouter().engine,
	}

	log.Fatal(srv.ListenAndServe())
}

type Router struct {
	engine     *gin.Engine
	repository repository.ArticleRepository
}

func NewRouter() *Router {
	router := &Router{}

	r := gin.New()

	r.Use(gin.Logger())
	r.GET("/article/:id", router.homeHandler)

	repo := memory.NewAritcleRepository()
	router.repository = repo
	router.engine = r
	return router
}

func (r *Router) homeHandler(c *gin.Context) {
	ctx := context.Background()
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	log.Println("Error parsing param id", err)

	article, err := r.repository.ArticleById(ctx, id)
	if err != nil {
		log.Println("Error finding article: ", err)
		c.JSON(500, gin.H{"500": err.Error()})
	}
	c.JSON(200, article)
	return
}

func Aricle(c *gin.Context) {
	a := c.Param("id")
	fmt.Println(a)

}
