package request

import (
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
)

//PostDocument - struct model for save requests. Edit only first fields group equal to your documents to save
type PostDocument struct {
	Document    `json:"-"`
	ID          uint64 `json:"id,omitempty" form:"id" binding:"required"`
	Data        string `json:"data,omitempty" form:"data" binding:"required"`
	LastIndexed string `json:"last_indexed" form:"last_indexed"`
	Enabled     bool   `json:"enabled" form:"enabled"`
}

//ValidateAndBind - Bind request params to PostDocument struct
func (r *PostDocument) ValidateAndBind(c *gin.Context, request *PostDocument) error {
	var bindErrors error = nil
	if bindErrors = c.ShouldBindJSON(request); bindErrors == nil {
		fmt.Println("Binding POST result: ", request)
	}

	return bindErrors
}

//GetID - Implementation of interface Document for index documents
func (r PostDocument) GetID() string {
	return strconv.FormatUint(r.ID, 10)
}
