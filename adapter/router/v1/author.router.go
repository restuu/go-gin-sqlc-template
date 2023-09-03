package v1

import (
	"go-gin-sqlc-template/domain"

	"github.com/gin-gonic/gin"
)

type authorHandler struct {
	uc domain.AuthorUsecase
}

func SetAuthorRoutes(r *gin.RouterGroup, uc domain.AuthorUsecase) {

	h := authorHandler{
		uc: uc,
	}

	authors := r.Group("/authors")

	authors.GET("", h.findAllAuthors)
}

func (h *authorHandler) findAllAuthors(c *gin.Context) {
	authors, err := h.uc.FindAllAuthors(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, authors)
}
