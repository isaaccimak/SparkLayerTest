package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"os"
	"encoding/csv"
)

type Response struct {
	PeopleJson []PeopleJson `json:"people"`
}
type PeopleJson struct {
	Craft string `json:"craft"`
	Name string `json:"name"`
}
type People struct {
	Craft string
	Name string
}

func main() {
    response, err := http.Get("http://api.open-notify.org/astros.json")
    if err != nil {
        log.Fatal(err)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    // fmt.Println(string(responseData))
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	
	var PeopleList []People
	for i := 0; i < len(responseObject.PeopleJson); i++ {
		PeopleList = append(PeopleList, People{responseObject.PeopleJson[i].Craft, responseObject.PeopleJson[i].Name})
	}
	for i := 0; i < len(PeopleList); i++ {
		fmt.Println(PeopleList[i])
	}
	sort.Slice(PeopleList, func(i, j int) bool {
		if PeopleList[i].Craft == PeopleList[j].Craft {
			return PeopleList[i].Name < PeopleList[j].Name
		} else {
			return PeopleList[i].Craft < PeopleList[j].Craft
		}
	})

	fmt.Println("\nAfter sort")
	for i := 0; i < len(PeopleList); i++ {
		fmt.Println(PeopleList[i])
	}
	
	file, err := os.Create("test.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Craft"})
	for _, value := range PeopleList {
		writer.Write([]string{value.Name, value.Craft})
	}
}