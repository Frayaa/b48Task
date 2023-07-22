package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type Form struct {
	ProjectName string
	Start string
	End string
	Description string
	Nodejs bool
	Reactjs bool
	Javascript bool
	Typescript bool
}

var formData = []Form {
	// {
	// ProjectName: "lala",
	// Start: "2022-02-02",
	// End: "2022-02-04",
	// Description: "blabla",
	// Nodejs: true,
	// Reactjs: true,
	// Javascript: true,
	// Typescript: true,
	// },
	{  
	ProjectName: "blabalablabla",
	Start: "2023-08-06",
	End: "2023-08-10",
	Description: "blablablalala",
	Nodejs: true,
	Reactjs: false,
	Javascript: true,
	Typescript: false,
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
		"forms": formData,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	projectName := c.FormValue("project")
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
		ProjectName: projectName,
		Start: start,
		End: end,
		Description: description,
		Reactjs: (reactjs == "reactjs"),
		Nodejs: (nodejs == "nodejs"),
		Javascript: (javascript == "javascript"),
		Typescript: (typescript == "typescript"),

	}
	formData = append(formData, newBlog)

	fmt.Println("project:", projectName)
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
				ProjectName: data.ProjectName,
				Start: data.Start,
				End: data.End,
				Description: data.Description,
				Reactjs: data.Reactjs,
				Nodejs: data.Nodejs,
				Javascript: data.Javascript,
				Typescript: data.Typescript,
			}
		}
	}

	data := map[string] interface{}{
		"id": id,
		"ProjectDetail": blogDetail,
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

