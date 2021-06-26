package main

import (
	"ExpertSystem/AnalyserFuzzy"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Analyser interface {
	Analyse(ans []int) (rateOfIntimacy, rateOfAnxiety int, err error)
	GetResult() (attachmentStyle, description string)
}

type AnswerSheet struct {
	Ans []int `json:"ans"`
}

func main() {
	r := gin.Default()
	r.Delims("{!{", "}!}")
	r.LoadHTMLFiles("./static/index.html")
	r.GET("/", Index)
	r.StaticFS("/css", http.Dir("./static/css"))
	r.StaticFS("/js", http.Dir("./static/js"))
	r.StaticFS("/fonts", http.Dir("./static/fonts"))
	r.POST("/analyse", Analyse)
	r.Run("127.0.0.1:80")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Analyse(c *gin.Context) {
	//get AnsSheet
	// f, err := os.Open("res.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jsonBytes, err := ioutil.ReadAll(f)
	ansStr := c.Request.FormValue("ans")
	ansStr = `{"ans":` + ansStr + `}`
	// fmt.Println(ansStr)
	var as AnswerSheet
	err := json.Unmarshal([]byte(ansStr), &as)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(as)

	// analyse the answer sheet
	var analyser Analyser
	analyser = AnalyserFuzzy.New()
	rateOfIntimacy, rateOfAnxiety, err := analyser.Analyse(as.Ans)
	if err != nil {
		log.Fatal(err)
	}
	attachmentStyle, description := analyser.GetResult()
	// write response
	// fmt.Println(attachmentStyle, description, rateOfIntimacy, rateOfAnxiety)
	c.JSON(http.StatusOK, gin.H{
		"rateOfIntimacy":  rateOfIntimacy,
		"rateOfAnxiety":   rateOfAnxiety,
		"attachmentStyle": attachmentStyle,
		"description":     description,
	})
}
