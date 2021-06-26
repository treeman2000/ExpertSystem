package main

import (
	"ExpertSystem/AnalyserFuzzy"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Analyser interface {
	Analyse(ans []int) (rateOfIntimacy, rateOfAnxiety int, err error)
	GetResult() (attachmentStyle, description string)
}

type AnswerSheet struct {
	Ans []int `json:"ans"`
}

func main() {
	//get AnsSheet
	f, err := os.Open("res.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadAll(f)
	var as AnswerSheet
	err = json.Unmarshal(jsonBytes, &as)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(as)

	// analyse the answer sheet
	var analyser Analyser
	analyser = AnalyserFuzzy.New()
	rateOfIntimacy, rateOfAnxiety, err := analyser.Analyse(as.Ans)
	if err != nil {
		log.Fatal(err)
	}
	attachmentStyle, description := analyser.GetResult()
	// write response
	fmt.Println(attachmentStyle, description, rateOfIntimacy, rateOfAnxiety)
}
