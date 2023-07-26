package main

import (
	"context"
	"fmt"
	"html/template"
	"myapp/connection"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	Id int
	Name string
	Email string
	HashedPassword string
}

type UserLoginSession struct {
	IsLogin bool
	Name string
}

var userLoginSession = UserLoginSession{}

var formData = []Form {

}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("loginUser"))))

	e.Static("/assets", "assets")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/form-blog", formBlog)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/testimoni", testimoni)
	e.GET("/form-update/:id", FormUpdate)
	e.GET("/login", formLogin)
	e.GET("/register", formRegister)
	

	e.POST("/logout", logout)
	e.POST("/auth/login", login)
	e.POST("/auth/register", register)
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
	
	dataQuery, errData := connection.Conn.Query(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, image FROM tb_project")

	if errData != nil {
		return c.JSON(http.StatusInternalServerError, errData.Error())
	}

	formData = []Form{}
	for dataQuery.Next() {
		var each = Form{}

		err := dataQuery.Scan(&each.Id, &each.ProjectName, &each.Start, &each.End, &each.Description, &each.Technologies, &each.Image)
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
		if checkValue(each.Technologies, "typescript") {
			each.Typescript = true
		}

		formData = append(formData, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"forms": formData,
		"UserLogin": userLoginSession,
	}

	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	project_name := c.FormValue("project")
	start_date := c.FormValue("start")
	end_date := c.FormValue("end")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	reactjs := c.FormValue("reactjs")
	javascript := c.FormValue("javascript")
	typescript := c.FormValue("typescript")
	image := c.FormValue("image")

	date1, _ := time.Parse("2006-01-02", start_date)
	date2, _ := time.Parse("2006-01-02", end_date)

	technologies := []string{nodejs, reactjs, javascript, typescript}


	add, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_project (project_name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)",
	project_name, date1, date2, description, technologies, "project.jpeg")

	fmt.Println("rowaffected:", add.RowsAffected())

	if err != nil {
		fmt.Println("error")
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("project:", project_name)
	fmt.Println("Startdate:", start_date)
	fmt.Println("endDate:", end_date)
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

	blogDetail := Form{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, image FROM tb_project WHERE id=$1", id).Scan(&blogDetail.Id, &blogDetail.ProjectName,
	&blogDetail.Start, &blogDetail.End, &blogDetail.Description, &blogDetail.Technologies, &blogDetail.Image)

	fmt.Println("data query:", errQuery)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogDetail.Duration = timeDuration(blogDetail.Start, blogDetail.End)

	if checkValue(blogDetail.Technologies, "nodejs") {
		blogDetail.Nodejs = true
	}
	if checkValue(blogDetail.Technologies, "reactjs") {
		blogDetail.Reactjs = true
	}
	if checkValue(blogDetail.Technologies, "javascript" ) {
		blogDetail.Javascript = true
	}
	if checkValue(blogDetail.Technologies, "typescript") {
		blogDetail.Typescript = true
	}

		data := map[string] interface{}{
			"id" : id,
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

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", idToInt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}


// update
func FormUpdate(c echo.Context)error{
	id, _ := strconv.Atoi(c.Param("id"))

	blogDetail := Form{}

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, project_name, start_date, end_date, description, technologies, image FROM tb_project WHERE id=$1", id).Scan(&blogDetail.Id, &blogDetail.ProjectName,
	&blogDetail.Start, &blogDetail.End, &blogDetail.Description, &blogDetail.Technologies, &blogDetail.Image)

	fmt.Println("data query:", errQuery)
	
	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	blogDetail.Duration = timeDuration(blogDetail.Start, blogDetail.End)

	if checkValue(blogDetail.Technologies, "nodejs") {
		blogDetail.Nodejs = true
	}
	if checkValue(blogDetail.Technologies, "reactjs") {
		blogDetail.Reactjs = true
	}
	if checkValue(blogDetail.Technologies, "javascript" ) {
		blogDetail.Javascript = true
	}
	if checkValue(blogDetail.Technologies, "typescript") {
		blogDetail.Typescript = true
	}

		data := map[string] interface{}{
			"id" : id,
			"ProjectDetail" : blogDetail,
			"startDate" 	: blogDetail.Start.Format("2006-01-02"),
			"endDate"		: blogDetail.End.Format("2006-01-02"),
		}
	

	tmpl, err := template.ParseFiles("views/form-update.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c. Response(), data)
}

func updatedBlog(c echo.Context) error{
	id, _ := strconv.Atoi(c.FormValue("id"))
	project_name := c.FormValue("project")
	start_date := c.FormValue("start")
	end_date := c.FormValue("end")
	description := c.FormValue("description")
	nodejs := c.FormValue("nodejs")
	reactjs := c.FormValue("reactjs")
	javascript := c.FormValue("javascript")
	typescript := c.FormValue("typescript")

	date1, _ := time.Parse("2006-01-02", start_date)
	date2, _ := time.Parse("2006-01-02", end_date)
	technologies := []string{nodejs, reactjs, javascript, typescript}
	

	dataUpdate, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET project_name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5 WHERE id=$6", 
	project_name, date1, date2, description, technologies, id)
	
	fmt.Println("ini update", dataUpdate.RowsAffected())

	if err != nil {
		fmt.Println("ini error:", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("project:", project_name)
	fmt.Println("Startdate:", start_date)
	fmt.Println("endDate:", end_date)
	fmt.Println("description:", description)
	fmt.Println("nodejs:", nodejs)
	fmt.Println("reactjs:", reactjs)
	fmt.Println("javascript:", javascript)
	fmt.Println("typescript:", typescript)

	return  c.Redirect(http.StatusMovedPermanently, "/blog")
}

func timeDuration(a time.Time, b time.Time) string {

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

// redirectmessage
func redirectMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
} 

// Register 
func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-register.html")

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"flashMessage": sess.Values["message"],
		"flashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), flash)
}

func register(c echo.Context) error {
	name := c.FormValue("input_name")
	email := c.FormValue("input_email")
	password := c.FormValue("input_password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	
	query, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (username, email, password) VALUES ($1, $2, $3)", name, email, hashedPassword)
	
	fmt.Println("iniapa:", name, email, hashedPassword)
	fmt.Println("affected row:", query.RowsAffected())

	if err != nil {
		return redirectMessage(c, "Register berhasil", false, "/register")
	}

	return redirectMessage(c, "Register berhasil", true, "/login")
}


// Login
func formLogin(c echo.Context)error {
	tmpl, err := template.ParseFiles("views/form-login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"flashMessage": sess.Values["message"],
		"flashStatus" : sess.Values["status"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	email := c.FormValue("input-email")
	password := c.FormValue("input-password")

	user := User{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, username, email, password FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectMessage(c, "Login Failed", false, "/login")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))

	if errPassword != nil {
		return redirectMessage(c, "Login Gagal", false, "/login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 
	sess.Values["message"] = "Login Success"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/home")

}

// Logout
func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectMessage(c, "Logout Berhasil", true, "/home")
}






