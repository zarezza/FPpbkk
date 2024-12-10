package bookcontroller

import (
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Model *models.BookModel
}

func getUserID(c *gin.Context) (uint, error) {
	userID, exists := c.Get("userID")
	if !exists {
		return 0, fmt.Errorf("user not authenticated")
	}
	return userID.(uint), nil
}

func (ctrl *BookController) Index(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	books, err := ctrl.Model.FindByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "index.html", books)
}

func (ctrl *BookController) Add(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "add.html", nil)
		return
	}

	var book models.Book
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.UserID = userID
	if err := ctrl.Model.Create(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/books")
}

func (ctrl *BookController) Edit(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if c.Request.Method == http.MethodGet {
		book, err := ctrl.Model.FindByIDAndUser(uint(id), userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "edit.html", book)
		return
	}

	var book models.Book
	if err := c.ShouldBind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)
	book.UserID = userID

	if err := ctrl.Model.Update(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/books")
}

func (ctrl *BookController) Delete(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Model.Delete(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/books")
}
