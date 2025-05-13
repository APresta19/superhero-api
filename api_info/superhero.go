package api_info

import(
	"encoding/json"
	"net/http"
)

type Superhero struct {
	ID int
}

type SuperheroResponse struct {
	Code int
	ID int
	Name string
	Ability string
}

type Error struct {
	Code int
	Message string
}

//write a successful json - can also use interface{} for a generic struct or obj
func WriteJSON(w http.ResponseWriter, superhero SuperheroResponse) {
	//set status code
	w.WriteHeader(superhero.Code)
	//set the header to of type json
	w.Header().Set("Content-Type", "application/json")
	//use NewEncoder to encode and send the response back
	json.NewEncoder(w).Encode(superhero)
}

func WriteError(write http.ResponseWriter, message string, code int) {
	//create error struct
	response := Error{
		Code: code,
		Message: message,
	}

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(code)
	json.NewEncoder(write).Encode(response)
}

var (
	SpecificError = func(write http.ResponseWriter, err error) {
		WriteError(write, err.Error(), http.StatusBadRequest)
	}
	UnexpectedError = func(write http.ResponseWriter) {
		WriteError(write, "An unexpected error ocurred", http.StatusInternalServerError)
	}
)
