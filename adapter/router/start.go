package router

import (
	routerMiddleware "go-gin-sqlc-template/adapter/router/middleware"
	v1 "go-gin-sqlc-template/adapter/router/v1"
	"go-gin-sqlc-template/domain"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitHandler(usecases *domain.Usecases) http.Handler {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{
			http.MethodPost,
			http.MethodPut,
			http.MethodGet,
			http.MethodDelete,
		},
		AllowHeaders: []string{
			"X-Request-ID",
			"X-User-ID",
			"Authorization",
			"Origin",
			"User-Agent",
		},
	}))

	r.Use(routerMiddleware.ErrorHandler())

	v1Router := r.Group("/v1")
	{
		v1.SetAuthorRoutes(v1Router, usecases.AuthorUsecase)
	}

	return r.Handler()
}
