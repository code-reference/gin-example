package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	endpointport int
)

func init() {
	endpointport = 2222
}

func main() {
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)
	// router := gin.Default()
	router := gin.New()

	v1 := router.Group("/shentu/autotshlogin/v1")

	//Here is example list
	v1.GET("/user/:name", v1UserName)

	v1.GET("/user/:name/*action", v1UserNameAction)
	v1.GET("/test", v1Test)

	/*
		curl -i -X POST \
		   -H "Content-Type:application/json" \
		   -d \
		'{
		  "username": "hello",
		  "password": "world"
		}' \
		 'http://localhost:3678//shentu/autotshlogin/v1/testpost'
	*/
	v1.POST("/testpost", v1TestPOST)

	// curl http://127.0.0.1:8080/bastion/v1/welcome?loginname=Lisa\&password=123456
	v1.GET("/welcome", func(c *gin.Context) {
		loginname := c.DefaultQuery("loginname", "null")
		password := c.Query("password")

		c.String(http.StatusOK, "hello, %s %s \n", loginname, password)
	})
	// curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217" | python -m json.tool
	// curl -X POST http://127.0.0.1:8080/repmgreventnotifier/v1/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217" | python -m json.tool
	v1.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "OK",
			},
			"message": message,
			"nick":    nick,
		})
	})

	/*
		if config.EnableTLS {
			router.RunTLS(":5778", config.Appcontext.AppRootDir+config.Appcontext.AppConfDirRelativeroot+config.Appcontext.AppServerCert, config.Appcontext.AppRootDir+config.Appcontext.AppConfDirRelativeroot+config.Appcontext.AppServerCertKey)
		} else {
			router.Run(":5778")
		}

	*/

	router.Run(":3678")
}
