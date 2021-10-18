package routes

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
)
type takeStringData struct {
    StrData    string `json:"stringtext" binding:"required"`
}

func welcome(c *gin.Context) {
	var result bool = false
	var requestBody takeStringData
	stringTxt,bovar :=c.GetPostForm("stringtext")
log.Println(stringTxt,"\n",bovar)
	if err := c.BindJSON(&requestBody); err == nil {
		c.AbortWithStatus(400)
		log.Println(err)
		return
	}
	f, err := os.Open("./insultWords.csv")

    if err != nil {

        log.Fatal(err)
    }

    r := csv.NewReader(f)

    for {

        record, err := r.Read()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }

        for value := range record {
            if strings.Contains(stringTxt,record[value]){
				result = true
			} 
        }
    }
	if (result){	c.JSON(http.StatusOK,gin.H{
		"result": result,
	}) }else{
		c.JSON(http.StatusNotFound,gin.H{
			"result": result,
		})	
	}
}
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	
}