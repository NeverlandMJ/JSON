package json

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Months struct {
	Abbreviation string
	Name         string
}

// parsing json to slice of Months struct through http request
func MonthsJson(url string) []Months {
	result := []Months{}

	monthClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	// get http request for raw months.json
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Sunbula", "Months")

	// getting response in 5 seconds
	res, getErr := monthClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	
	// parsing json to slice
	myJson := res.Body
	jsonErr := json.NewDecoder(myJson).Decode(&result)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	
	return result
}

type people struct{
	Name string
	Craft string
}

type Astros struct{
	Message string
	People []people
	Number int
}


// parsing nested json through http request
func AstrosJson(url string) Astros {
	result := Astros{}

	monthClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}

	// get http request for raw astros.json
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// getting response in 5 seconds
	res, getErr := monthClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	myJson := res.Body

	jsonErr := json.NewDecoder(myJson).Decode(&result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	
	return result
}

type AsianCountries struct{
	Name string `json: "name"`
	Code string `json: "code"`
}

// marshalling the json data
func Struct2Json(){
	countries := []AsianCountries{
		{
			Name: "Uzbekistan",
			Code: "UZ",
		},
		{
			Name: "Tajikistan",
			Code: "TJ",
		},
		{
			Name: "Turkmenistan",
			Code: "TM",
		},
		{
			Name: "Kazakistan",
			Code: "KZ",
		},
		{
			Name: "Kyrgyzstan",
			Code: "KG",
		},
	}

	data, err := json.Marshal(countries)

	if err != nil{
		fmt.Errorf("Error: %e", err)
		return
	}

	fmt.Println(string(data))

}