// 0.1
package main

//0.5
import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 4.1
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 0.2
func main() {
	//0.3
	r := gin.Default()
	//0.6
	runningDir, _ := os.Getwd()
	//0.11 ทดสอบ
	count := 0

	//0.7
	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0000)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0000)

	//0.8 ผูกไฟล์ กับตัว logger
	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	//0.9.1 standard
	r.Use(gin.Logger())

	//0.9.2 custom
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))

	//0.9.1.0 ไม่ให้แสดง log ของ path นั้นๆ ใน ไฟล์
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/login"))

	//0.10
	//Case 1. "/" Root Path
	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Root"))
	})

	//Case 2.
	r.GET("/profile", func(c *gin.Context) {
		//0.12 print ทดสอบ
		count = count + 1
		//0.9.1.1 ใช้ log standard
		//accesslogfile.WriteString(fmt.Sprintf("Count : %d\n", count))
		//0.9.2.1 ใช้ log custom
		accesslogfile.WriteString(fmt.Sprintf("Count : %d", count))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("profile"))
	})

	//Case 3.
	r.GET("/error", func(c *gin.Context) {
		//0.12 print ทดสอบ
		count = count + 1
		//0.9.1.2 ใช้ log standard
		//errlogfile.WriteString(fmt.Sprintf("Error : %d\n", count))
		//0.9.2.2 ใช้ log custom
		errlogfile.WriteString(fmt.Sprintf("Error : %d\n", count))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("error"))
	})

	//Case 4. POST /login
	//4.2
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
	r.Run(":85")

}
