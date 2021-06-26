package main

import (
	"ExpertSystem/AnalyserFuzzy"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var host string
var config Config

func init() {
	// load config
	f, err := os.Open("config.json")
	configBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("load config failed")
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Host + ":" + config.Port)
}

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
	r.Run(config.Host + ":" + config.Port)
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

	jsonBytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(jsonBytes))
	var as AnswerSheet
	err = json.Unmarshal(jsonBytes, &as)
	if err != nil {
		log.Fatal(err)
	}

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
