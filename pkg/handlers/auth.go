package handler

import (
	"net/http"

	"github.com/aleks-tim/todo-app"
	"github.com/gin-gonic/gin"
	// "github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	// logrus.Println("Handler -> signUp()")

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// logrus.Printf("	--> Name	: %s", input.Name)
	// logrus.Printf("	--> Username: %s", input.Username)
	// logrus.Printf("	--> Password: %s", input.Password)

	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
