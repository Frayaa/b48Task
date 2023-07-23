package main

import (
	"context"
	"fmt"
	"html/template"
	"myapp/connection"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)


type Form struct {
	Id int
	ProjectName string
	Start time.Time
	End time.Time
	Duration string
	Description string
	Technologies []string
	Nodejs bool
	Reactjs bool
	Javascript bool
	Typescript bool
	Image string
}

var formData = []Form {
	// {
	// Id: 0,
	// ProjectName: "lalalalal",
	// Start: "2022-02-02",
	// End: "2022-02-04",
	// Duration: timeDuration("2022-02-02", "2022-02-04"),
	// Description: "blabla",
	// Nodejs: true,
	// Reactjs: true,
	// Javascript: true,
	// Typescript: true,
	// },
	// {  
	// Id:1,
	// ProjectName: "blabalablabla",
	// Start: "2022-02-02",
	// End: "2020-04-02",
	// Duration: timeDuration("2022-02-02", "2020-04-02"),
	// Description: "blablablalala",
	// Nodejs: true,
	// Reactjs: false,
	// Javascript: true,
	// Typescript: false,
	// },
	// {  
	// Id:2,
	// ProjectName: "test",
	// Start: "2020-08-01",
	// End: "2020-10-11",
	// Duration: timeDuration("2020-08-01", "2020-10-11"),
	// Description: "blablablalala",
	// Nodejs: true,
	// Reactjs: true,
	// Javascript: true,
	// Typescript: false,
	// },

}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

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
	
	data1, errData := connection.Conn.Query(context.Background(), "SELECT * FROM tb_project")

	if errData != nil {
		return c.JSON(http.StatusInternalServerError, errData.Error())
	}

	formData = []Form{}
	for data1.Next() {
		var each = Form{}

		err := data1.Scan(&each.Id, &each.ProjectName, &each.Start, &each.End, &each.Description, &each.Technologies, &each.Image)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())

		}

		each.Duration = timeDuration(each.Start, each.End )
		if checkValue(each.Technologies, "nodejs") {
			each.Nodejs = true
		}
		if checkValue(each.Technologies, "reactjs") {
			each.Reactjs = true
		}
		if checkValue(each.Technologies, "javascript" ) {
			each.Javascript = true
		}
		if checkValue(each.Technologies, "typeScript") {
			each.Typescript = true
		}

		formData = append(formData, each)
	}

	data := map[string]interface{}{
		"forms": formData,
	}

	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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

	date1, _ := time.Parse("2006-01-02", start)
	date2, _ := time.Parse("2006-01-02", end)

	// append
	var newBlog = Form{
		ProjectName: projectName,
		Start: date1,
		End: date2,
		Duration: timeDuration(date1, date2),
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
				Duration: data.Duration,
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
		"ProjectDetail" : blogDetail,
		"startDate" 	: blogDetail.Start.Format("2006-01-02"),
		"endDate"		: blogDetail.End.Format("2006-01-02"),
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
				Duration: data.Duration,
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
	
	date1, _ := time.Parse("2006-01-02", start)
	date2, _ := time.Parse("2006-01-02", end)

	formData[id].ProjectName = projectName
	formData[id].Start = date1
	formData[id].End = date2
	formData[id].Duration = timeDuration(date1, date2 )
	formData[id].Description = description
	formData[id].Nodejs = (nodejs == "nodejs")
	formData[id].Reactjs = (reactjs == "reactjs")
	formData[id].Javascript = (javascript == "javascript")
	formData[id].Typescript = (typescript == "typescript")

	return  c.Redirect(http.StatusMovedPermanently, "/blog")
}

func timeDuration(a time.Time, b time.Time) string {
	// date1, _ := time.Parse("2006-01-02", a)
	// date2, _ := time.Parse("2006-01-02", b)

	difference := b.Sub(a)
	days := int(difference.Hours() / 24)
	weeks := days / 7
	months := days / 30

	if months > 12 {
		return strconv.Itoa(months/12) + " tahun"
	}
	if months > 0 {
		return strconv.Itoa(months) + " bulan"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " minggu"
	}
	return strconv.Itoa(days) + " hari"
}

func checkValue(tech []string, obj string) bool {
	for _, data := range tech {
		if data == obj {
			return true
		}

	}
	return false

}

