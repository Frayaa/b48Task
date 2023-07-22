package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type Form struct {
	Id int
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
	{
	Id: 0,
	ProjectName: "lalalalal",
	Start: "2022-02-02",
	End: "2022-02-04",
	Description: "blabla",
	Nodejs: true,
	Reactjs: true,
	Javascript: true,
	Typescript: true,
	},
	{  
	Id:1,
	ProjectName: "blabalablabla",
	Start: "06-08-2020",
	End: "06-10-2020",
	Description: "blablablalala",
	Nodejs: true,
	Reactjs: false,
	Javascript: true,
	Typescript: false,
	},
	{  
	Id:2,
	ProjectName: "test",
	Start: "06-08-2020",
	End: "06-10-2020",
	Description: "blablablalala",
	Nodejs: true,
	Reactjs: true,
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
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/testimoni", testimoni)
	e.GET("/form-update/:id", FormUpdate)
	
	e.POST("/add-blog", addBlog)
	e.POST("/update-blog", updatedBlog)
	e.POST("/delete/:id", deleteBlog)
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
	var newBlog = Form{
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
				Id:index,
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
		"id" : idToInt,
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

// delete
func deleteBlog(c echo.Context) error {
	id := c.Param("id")
	idToInt, _ := strconv.Atoi(id)

	formData = append(formData[:idToInt], formData[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")


}


// update
func FormUpdate(c echo.Context)error{

	id, _ := strconv.Atoi(c.Param("id"))
	blogUpdate := Form{}

	for index, data := range formData{
		if id == index{
			blogUpdate = Form{
				Id: index,
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
		"forms": blogUpdate,
	}

	tmpl, err := template.ParseFiles("views/form-update.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	return tmpl.Execute(c. Response(), data)
}

func updatedBlog(c echo.Context) error{
	id, _ := strconv.Atoi(c.FormValue("id"))
	projectName := c.FormValue("project")
	start := c.FormValue("start")
	end := c.FormValue("end")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	reactjs := c.FormValue("reactjs")
	javascript := c.FormValue("javascript")
	typescript := c.FormValue("typescript")
	

	formData[id].ProjectName = projectName
	formData[id].Start = start
	formData[id].End = end
	formData[id].Description = description
	formData[id].Nodejs = (nodejs == "nodejs")
	formData[id].Reactjs = (reactjs == "reactjs")
	formData[id].Javascript = (javascript == "javascript")
	formData[id].Typescript = (typescript == "typescript")
	// append
 	// newBlog := Form{
	// 	ProjectName: projectName,
	// 	Start: start,
	// 	End: end,
	// 	Description: description,
	// 	Reactjs: (reactjs == "reactjs"),
	// 	Nodejs: (nodejs == "nodejs"),
	// 	Javascript: (javascript == "javascript"),
	// 	Typescript: (typescript == "typescript"),

	// }
	
	// formData[id] = newBlog

	return  c.Redirect(http.StatusMovedPermanently, "/blog")
}

// func formatDuration(duration time.Duration) string {
// 	months := duration / (time.Hour * 24 * 30)
// 	duration %= time.Hour * 24 * 30
// 	weeks := duration / (time.Hour * 24 * 7)
// 	duration %= time.Hour * 24 * 7
// 	days := duration / (time.Hour * 24)
// 	duration %= time.Hour * 24

// 	return fmt.Sprintf("%d months, %d weeks, %d days", months, weeks, days)
// }

