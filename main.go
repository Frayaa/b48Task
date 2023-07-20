package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type Form struct {
	project string
	start string
	end string
	description string
}

var formData = []Form {
	{
	project: "lala",
	start: "12/10/2020",
	end: "12/11/2020",
	description: "blabla",
},
	{
	project: "lalalala",
	start: "12/10/2020",
	end: "12/11/2020",
	description: "blablablalala",
},

}
func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/form-blog", formBlog)
	e.GET("/blog", blog)
	e.POST("/add-blog", addBlog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/testimoni", testimoni)

	e.GET("/about", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string {
			"message": "Hello World",
		})
	})

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// Handler
func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func formBlog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	data := map[string]interface{}{
		"Form": formData,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	project := c.FormValue("project")
	start := c.FormValue("start")
	end := c.FormValue("end")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	reactjs := c.FormValue("reactjs")
	javascript := c.FormValue("javascript")
	typescript := c.FormValue("typescript")
	image := c.FormValue("image")

	// append
	newBlog := Form{
		project: project,
		start: "12/10/2020",
		end: "12/12/2020",
		description: description,
	}
	formData = append(formData, newBlog)

	fmt.Println("projectName:", project)
	fmt.Println("Startdate:", start)
	fmt.Println("endDate:", end)
	fmt.Println("description:", description)
	fmt.Println("nodejs:", nodejs)
	fmt.Println("reactjs:", reactjs)
	fmt.Println("javascript:", javascript)
	fmt.Println("typescript:", typescript)
	fmt.Println("image:", image)

	return  c.Redirect(http.StatusMovedPermanently, "/blog")
}

func blogDetail(c echo.Context) error {
id := c.Param("id")

	tmpl, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	blogDetail := Form{}

	for index, data := range formData{
		if index == idToInt{
			blogDetail = Form{
				project: data.project,
				start: data.start,
				end: data.end,
				description: data.description,
			}
		}
	}

	data := map[string] interface{}{
		"id": id,
		"project": blogDetail,
	}

	
	return tmpl.Execute(c. Response(), data)
}

func testimoni(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimoni.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

