package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Session(c *gin.Context) {
	if len(domain.GamesMap) == 0 {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}
