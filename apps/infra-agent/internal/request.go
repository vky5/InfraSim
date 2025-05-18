package loadgen

import (
	"fmt"
	"net/http"
)
func sendHTTPGetRequest(url string) {
	
	// this sends the request and if there is a response it stores it in resp and if error in err
	resp, err := http.Get(url)

	// if err is not nil 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)

}