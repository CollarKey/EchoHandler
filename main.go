package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type RequestBody struct {
	Task string `json:"task"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var task string

func getHandler(c echo.Context) error {
	answer := "Hello " + task // самостоятельно записал, как task := "Hello" + task, при проверке выводилось несколько раз Hello Hello... ГПТ помог
	return c.JSON(http.StatusOK, answer)
}

func postHandler(c echo.Context) error {
	var requestBody RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, Response{"Error", "Could not add the task"})
	}
	task = requestBody.Task
	return c.JSON(http.StatusOK, Response{"Status Success", "Task added"})
}

func main() {

	e := echo.New()

	e.GET("/task", getHandler)
	e.POST("/task", postHandler)

	log.Println("Server is running on port 8080")
	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(e.Start(":8080"))
	}
}
