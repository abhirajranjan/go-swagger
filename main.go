package main

import (
	"net/http"
	"os"
	"strconv"

	"goswag/models"

	"github.com/gin-gonic/gin"
)

const adminId = "404"

func main() {
	for _, arg := range os.Args {
		switch arg {
		case "swagger", "-s", "build":
			build()
			return
		}
	}

	router := gin.Default()
	defineRoute(router)

	router.Run()
}

func defineRoute(r *gin.Engine) {
	r.GET("/view", getAllTodos)
	r.GET("/view/:id", getTodoByID)
	r.POST("/admin/add", addTodo)
}

var todoList = []models.Todo{
	{Id: 1, Task: "Learn Go"},
	{Id: 2, Task: "Build an API with Go"},
	{Id: 3, Task: "Document the API with swag"},
}

// GetAllTodos godoc
// @Summary      Get all todos
// @Description  get all todos
// @Tags         view
// @Produce      json
// @Success      200  {object}  models.Todo
// @Router       /view [get]
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// GetTodoByID godoc
// @Summary      get todo by id
// @Description  get todo by id
// @Tags         view
// @Accept       json
// @Produce      json
// @Param        id    path       int  true  "todo id"
// @Success      200  {object}    models.Todo
// @Failure      404  {object}    models.Message
// @Failure      400  {object}    models.Message
// @Router       /view/{id}   [get]
func getTodoByID(c *gin.Context) {
	paramId := c.Param("id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			Message: "malformed id",
		})

		return
	}

	for _, todo := range todoList {
		if todo.Id == int(id) {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, models.Message{
		Message: "todo not found",
	})
}

// AddTodo godoc
// @Summary      Add todo
// @Description  add todo
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        id     header   string              true  "Account ID"
// @Param        body   body     models.NewMessage   true  "new message"
// @Success      200   {object} models.Message
// @Failure      401  {object}  models.Message
// @Failure      400  {object}  models.Message
// @Router       /admin/add [POST]
func addTodo(c *gin.Context) {
	accountId := c.GetHeader("id")

	if accountId != adminId {
		c.JSON(http.StatusUnauthorized, models.Message{
			Message: "user id unauthorized",
		})

		return
	}

	newMessage := &models.NewMessage{}
	if err := c.ShouldBind(newMessage); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			Message: err.Error(),
		})

		return
	}

	if newMessage.Task == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			Message: "task cannot be empty",
		})

		return
	}

	todoList = append(todoList, models.Todo{Id: newMessage.TodoId, Task: newMessage.Task})
	c.JSON(http.StatusOK, models.Message{
		Message: "Ok",
	})
}
