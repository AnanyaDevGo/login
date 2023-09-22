package main

import (
	"github.com/gin-gonic/gin"
)

type UserSt struct {
	Username string
	Password string
}

var users = make(map[string]UserSt)
var session = make(map[string]string)

type errors struct {
	UsernameError string
	PasswordError string
}

func init() {

	users["anu@gmail.com"] = UserSt{"Anu", "3456"}
	users["vandu@gmail.com"] = UserSt{"Vandu", "3456"}
	users["achu@gmail.com"] = UserSt{"Achu", "4561"}

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", loginPage)
	router.POST("/", loginPage)
	router.GET("/home", home)
	router.POST("/logout", logout)

	router.Run()
}
