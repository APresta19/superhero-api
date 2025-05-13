package handlers

import (
	"fmt"
	"net/http"
)

/*
	Grab the character_id from the api and setup a "database"
	Database is a struct returned from the api
	Perform any error checking
*/

//https://superheroapi.com/api/52826c755f50a2f10dfde3cfc418ea08
func GetDatabase() (*http.Response, error) {
	//*http.Response, err
	res, err := http.Get("https://superheroapi.com/api/52826c755f50a2f10dfde3cfc418ea08")

	if(err != nil) {
		fmt.Println("Error in DB: %v", err)
		return nil, err
	}

	//close response
	defer res.Body.Close()

	return res, nil
}