package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pooria1/cloud-services/backend/database"
)

func RunServer() {

	e := echo.New()

	e.GET("/", showFirstPage)
	e.POST("/newcomment1", insertNewComment1)
	e.POST("/newcomment2", insertNewComment2)
	e.POST("/newcomment3", insertNewComment3)

	e.Logger.Fatal(e.Start(":8080"))
}

func insertNewComment1(c echo.Context) error {
	return insertNewCommnet(1, c)
}

func insertNewComment2(c echo.Context) error {
	return insertNewCommnet(2, c)
}

func insertNewComment3(c echo.Context) error {
	return insertNewCommnet(3, c)
}

func insertNewCommnet(postNumber int, c echo.Context) error {
	res, err := addCommentToHTML(c.FormValue("comments"), postNumber)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
	}
	return c.HTML(http.StatusOK, res)
}

func showFirstPage(c echo.Context) error {
	commnets1, err := database.GetFromDB("1")
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
	}
	fmt.Println("len of c1 is ", len(commnets1))
	for _, comment := range commnets1 {
		if _, err := addCommentToHTML(comment, 1); err != nil {
			return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
		}
	}

	commnets2, err := database.GetFromDB("2")
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
	}
	fmt.Println("len of c2 is ", len(commnets2))
	for _, comment := range commnets2 {
		if _, err := addCommentToHTML(comment, 2); err != nil {
			return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
		}
	}

	commnets3, err := database.GetFromDB("3")
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
	}
	fmt.Println("len of c3 is ", len(commnets3))
	for _, comment := range commnets3 {
		if _, err := addCommentToHTML(comment, 3); err != nil {
			return c.HTML(http.StatusInternalServerError, "an error has been occured: "+err.Error())
		}
	}

	return c.HTML(http.StatusOK, htmlTemplateVariable)
}
