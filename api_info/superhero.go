package api_info

type Superhero struct {
	ID int
}

type SuperheroResponse struct {
	Code int
	Name string
}

type Error struct {
	Code int
	Message string
}

func writeError(write http.ResponseWriter, message string, code int) {
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
		writeError(write, err.Error(), http.StatusBadRequest)
	}
	UnexpectedError = func(write http.ResponseWriter) {
		writeError(write, "An unexpected error ocurred", http.StatusInternalServerError)
	}
)
