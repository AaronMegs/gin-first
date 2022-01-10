// Only Bind Query String
//ShouldBindQuery function only binds the query params and not the post data
// try
// only bind query
//$ curl -X GET "localhost:8085/testing?name=eason&address=xyz"
// only bind query string, ignore form data
//$ curl -X POST "localhost:8085/testing?name=eason&address=xyz" --data 'name=ignore&address=ignore' -H "Content-Type:application/x-www-form-urlencoded"
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Person define struct
type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	engine := gin.Default()
	engine.Any("/testing", startPage)

	engine.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	// must bind
	//if c.BindQuery(&person) == nil {}
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "Success")

}
