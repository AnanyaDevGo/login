package main

import (
 "fmt"
 "net/http"

 "github.com/gin-gonic/gin"
 // uuid "github.com/satori/go.uuid"
)

var errorStruct errors

func loginPage(c *gin.Context) {
 c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
 c.Header("Expires", "0")
 cookie, err := c.Cookie("session")
 if err != nil {
  fmt.Println(err)

 }

 if err == nil && cookie != "" {
  c.Redirect(303, "/home")
  return
 }
 
 if c.Request.Method == "POST" {
  username := c.Request.FormValue("userName")
  password := c.Request.FormValue("password")

  
  if _, ok := users[username]; !ok {
   errorStruct.UsernameError = "Invalid Username"
   c.Redirect(http.StatusSeeOther, "/")
   return
  }

  if password != users[username].Password {
   errorStruct.PasswordError = "Invalid Password"
   c.Redirect(http.StatusSeeOther, "/")
   return
  }

  errorStruct.UsernameError = ""
  errorStruct.PasswordError = ""

  if password == users[username].Password {


   c.SetSameSite(http.SameSiteLaxMode)
   c.SetCookie("session", "123", 3000, "/", "", false, true)

   c.Redirect(http.StatusSeeOther, "/home")
   return
  }
 }

 c.HTML(200, "index.login.html", errorStruct)
}
func home(c *gin.Context) {
 c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
 c.Header("Expires", "0")

 cookie, err := c.Cookie("session")

 if err != nil || cookie == "" {
  fmt.Println(err)
  c.Redirect(http.StatusSeeOther, "/")
  return
 }

 username := session[cookie]
 us := users[username]

 c.HTML(http.StatusOK, "index.home.html", us)
}

func logout(c *gin.Context) {
 c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
 c.Header("Expires", "0")

 c.SetCookie("session", "", -1, "/", "", false, true)

 c.Redirect(http.StatusSeeOther, "/")

}