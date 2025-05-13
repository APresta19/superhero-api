package handlers

import(
	"fmt"
	"net/http"
	"superheroapi/api_info"
	"strconv"
	"io/ioutil"
)

func GetSuperheroAbility(w http.ResponseWriter, r *http.Request) {
	db, err := GetDatabase()

	//this code is also written in database.go but maybe it should be added if there is an error in between
	if(err != nil) {
		return
	}
	fmt.Println(db)

	request := r.URL.Query()
	fmt.Println("Request: ", request)

	s_id_str := request.Get("id")
	s_name := request.Get("name")
	s_ability, err_ab := http.Get("https://superheroapi.com/api/52826c755f50a2f10dfde3cfc418ea08/" + s_id_str + "/powerstats")

	if(err_ab != nil) {
		panic(err_ab)
	}

	//type conversions and cleaning
	s_id, conv_err := strconv.Atoi(s_id_str)
	if conv_err != nil {
		fmt.Println("String conversion error with superhero ID.")
		fmt.Printf("S_ID: %v", s_id_str)
		return
	}
	s_ability_body, err := ioutil.ReadAll(s_ability.Body)
	fmt.Println(string(s_ability_body))
	superhero := api_info.SuperheroResponse{
		Code: http.StatusOK,
		ID: s_id,
		Name: s_name,
		Ability: string(s_ability_body),
	}

	api_info.WriteJSON(w, superhero)
}