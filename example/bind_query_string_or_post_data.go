// bind by query
// curl -X GET "localhost:8085/testing?name=appleboy&address=xyz"
// bind by json
// curl -X GET localhost:8085/testing --data '{"name":"JJ", "address":"xyz"}' -H "Content-Type:application/json"
// curl -X GET "localhost:8085/testing?name=appleboy&address=xyz&birthday=1992-03-15&createTime=1562400033000000123&unixTime=1562400033"
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person2 struct {
	Name       string    `form:"name" json:"name"`
	Address    string    `form:"address" json:"address"`
	Birthday   time.Time `form:"birthday" json:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" json:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" json:"unixTime" time_format:"unix"`
}

func main() {
	engine := gin.Default()
	engine.GET("/testing", startPage2)
	engine.Run(":8085")

}

func startPage2(c *gin.Context) {
	var person Person2
	// 通用
	if c.ShouldBind(&person) == nil {
		log.Println("====== Bind By Query and JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}
	// only query
	if c.Bind(&person) == nil {
		log.Println("====== Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	// only json
	if c.BindJSON(&person) == nil {
		log.Println("====== Bind By JSON ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "Success")
}
