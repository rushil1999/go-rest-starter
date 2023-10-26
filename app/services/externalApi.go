package services

import (
	"net/http"
	"fmt"
	"io"
	"reflect"
	"encoding/json"
)


func GetUniqueBankAccountNumber() map[string]interface{} {
	url := "https://random-data-api.com/api/v2/users"
	httpMehod := "GET"
	request, err := http.NewRequest(httpMehod, url, nil)
	if err != nil {
		fmt.Println("Could not establish http request")
		// Handle custom error here
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Could not execute http request")
		// Handle custom error here
	}
	defer response.Body.Close() // To prevent data leaks

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not get response body")
		// Handle custom error here
	}
	var parsedBody interface{}
	json.Unmarshal(body, &parsedBody)
	fmt.Println("External API call -------------------------------")
	fmt.Println( parsedBody, reflect.TypeOf(parsedBody))
	return parsedBody.(map[string]interface{})
}