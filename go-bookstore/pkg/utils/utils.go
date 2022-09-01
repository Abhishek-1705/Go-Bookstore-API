package utils //include the package

import ( //Importing the libary from the package that is inhertines  the main
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Function to pass the body of request for the json Unmarshlling
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { //Reading the all tha data for Request
		if err := json.Unmarshal([]byte(body), x); err != nil { //unmarshalling the Body of the Request
			return
		}
	}
}
