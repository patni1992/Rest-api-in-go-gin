package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns all users
// @Summary		Returns all users
// @Description	Returns all users
// @Tags			users
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]database.User
// @Router			/api/v1/users [get]
// @Security		BearerAuth
func (app *application) getAllUsers(c *gin.Context) {
	users, err := app.models.Users.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive users"})
	}

	c.JSON(http.StatusOK, users)
}
