package main

import (
	"fmt"
	"lesson/pkg/JSON"
)

type Months struct {
	Abbreviation string 
  	Name string 
}



func main() {

	MonthsUrl := "https://gist.github.com/enriqueornelasjr/57cc380aa96a3e041e4886202c30e40c/raw/bccfbee870c0bff1ba04b35eb84f515923b8236b/Months.json"

	output := json.MonthsJson(MonthsUrl)
	for _, v := range output{
		fmt.Printf("Abbrivation: %s  Name: %s\n", v.Abbreviation, v.Name)
	}



	AstrosUrl := "http://api.open-notify.org/astros.json"
	output1 := json.AstrosJson(AstrosUrl)

	fmt.Printf("\n\nMessage: %s\n", output1.Message)
	for _, v := range output1.People{
		fmt.Printf("Name: %s  Craft: %s\n", v.Name, v.Craft)
	}
	fmt.Printf("Number: %v\n\n", output1.Number)


	json.Struct2Json()


}




