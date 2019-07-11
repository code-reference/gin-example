package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func v1UserName(c *gin.Context) {
	name := c.Param("name")

	c.String(http.StatusOK, "Hello, %s %s \n", name, name)
}

// v1UserNameAction v1UserNameAction
func v1UserNameAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action + "\n"
	c.String(http.StatusOK, message)
}

//v1Test v1Test
func v1Test(c *gin.Context) {
	c.String(http.StatusOK, "This is a test.")
}

func v1TestPOST(c *gin.Context) {
	var account Account

	c.BindJSON(&account)

	log.Println(account)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"name":   account.Username,
		"passwd": account.Password,
	})
}
