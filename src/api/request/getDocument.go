package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
)

//GetDocument - struct model for search requests. Edit only first fields group equal to your "PostDocument"
//All field are strings, because come like query params
type GetDocument struct {
	ID          string `json:"id,omitempty" form:"id"`
	Data        string `json:"data,omitempty" form:"data"`
	Enabled     string `json:"enabled,omitempty" form:"enabled"`
	LastIndexed string `json:"last_indexed,omitempty" form:"last_indexed" binding:"isValidDateFormat"`

	Sort      string `json:"sort" form:"sort"`
	Criteria  string `json:"criteria" form:"criteria" default:"desc"`
	Offset    string `json:"offset" form:"offset"`
	Limit     string `json:"limit" form:"limit"`
	Range     string `json:"range" form:"range"`
	EndDate   string `json:"end_date" form:"end_date" binding:"isValidDateFormat"`
	BeginDate string `json:"begin_date" form:"begin_date" binding:"isValidDateFormat"`
}

//ValidateAndBind - Bind request params to GetDocument struct
func (r *GetDocument) ValidateAndBind(c *gin.Context, request *GetDocument) error {
	var bindErrors error = nil
	
	defaults.SetDefaults(request)
	if bindErrors = c.ShouldBindQuery(request); bindErrors == nil {
		fmt.Println("bind result: ", request)
		request.normalizeParams()
	}

	return bindErrors
}

func (r *GetDocument) normalizeParams() error {
	//normalize the data to search accord to data saved in database
	//Example: r.Data = strings.ToUpper(r.Data)
	return nil
}
