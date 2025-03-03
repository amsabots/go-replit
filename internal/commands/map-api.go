package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)


const locationURL = "https://pokeapi.co/api/v2/location-area/"

type MapResponse struct{
 Count int `json:"count"`
 Next string `json:"next"`
 Previous string `json:"previous"`
 Results []LocationInfo `json:"results"`
}

type LocationInfo struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func pokeAPI(cc *CommandConfig) error{

var mapData MapResponse
 
req, err := http.NewRequest("GET", locationURL, nil)
if err != nil {return fmt.Errorf("error creating request: %w", err)}

resp, err := cc.HtttpClient.HTTP.Do(req)
if err != nil {return fmt.Errorf("error making request: %w", err)}
defer resp.Body.Close()

if resp.StatusCode > 299 {
	return fmt.Errorf("response failed with status code: %d", resp.StatusCode)
}

err = json.NewDecoder(resp.Body).Decode(&mapData)
if err != nil {return err}

cc.Next = mapData.Next
cc.Previous = mapData.Previous

for _, value := range mapData.Results {
fmt.Println(value.Name)
}
return nil
}


func  pokeAPIBack(cc *CommandConfig) error{
    if cc.Previous == ""{
		fmt.Println("you're on the first page")
		return nil
	}
	var mapData MapResponse
	res, err := http.Get(cc.Previous)
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {return err}
	defer res.Body.Close()
	
	err = json.NewDecoder(res.Body).Decode(&mapData)
	if err != nil {return err}
	for _, value := range mapData.Results {
	fmt.Println(value.Name)
	}
	return nil
	}