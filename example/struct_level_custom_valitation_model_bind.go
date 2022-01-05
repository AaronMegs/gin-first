// 结构体级参数验证 参考： https://github.com/gin-gonic/examples/tree/master/struct-lvl-validations
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Users contains user information.
type Users struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:required,email`
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For example, this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't not just do this outside of validator. Doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
func UserStructLevelValidation(sl validator.StructLevel) {
	//user := structLevel.CurrentStruct.Interface().(User)
	users := sl.Current().Interface().(Users)

	if len(users.FirstName) == 0 && len(users.LastName) == 0 {
		sl.ReportError(users.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(users.LastName, "LastName", "lname", "fnameorlname", "")
	}

	// plus can to more, even with different tag than "fnameorlname"
}

func main() {
	router := gin.Default()
	// 使用了 接口 Comma-ok 断言 来确定类型
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, Users{})
	}
	router.POST("/user", validatorUser)
	router.Run(":8085")
}

func validatorUser(c *gin.Context) {
	var u Users
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User validation successful.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed.",
			"error":   err.Error(),
		})
	}

}
