// 0.1
package main

//0.5
import (
	"fmt"
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

// 3.1
func handleBookRequest2(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to})
}

// 5.1
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 0.2
func main() {
	//0.3
	r := gin.Default()

	//0.6
	//case 1. GET root(/) path
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("1234"))
	})
	//Request => http://localhost:85/
	//Response => 1234

	//case 2. GET /profile
	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("1234"))
	})
	//Request => http://localhost:85/profile
	//Reponse => profile

	//case 3. GET /book/thai/japan
	//3.2
	r.GET("/book/:from/:to", handleBookRequest2)
	//Request => http://localhost:85/book/thai/japan
	//Response =>
	// {
	// 	"from": "thai",
	// 	"result": "ok",
	// 	"to": "japan"
	// }

	//case 4 GET /login
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})
	//Request => http://localhost:85/login?username=admin&password=1234
	//Response =>
	// 	{
	//   "password": "1234",
	//   "result": "ok",
	//   "username": "admin"
	// }

	//Case 5. POST /login
	//5.2
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.Username == "admin" && form.Password == "12345" {
				msg := fmt.Sprintf("you are logged in with %s,%s", form.Username, form.Password)
				c.JSON(http.StatusOK, gin.H{"status": msg})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		}
	})

	//0.4
	//new port from http://localhost:8080/ to http://localhost:85/
	r.Run(":85")
}
