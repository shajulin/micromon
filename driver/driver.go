/**
micromon -- driver.go to test the micromon examples
Shajulin Benedict
*/

package main

import(
  "fmt"
  _"os"
  _"log"
  "io/ioutil"
  _"strings"
  "bytes"
  "net/http"
  "net/url"
  "encoding/json"
)

const (
	connectURL string = "http://localhost:8172/getallevent"
)

func main() {

  //Initialize the basic data structure for micromon
  type EventInfo struct {
      AppID   string `json:"appid,omitempty"`
      QueryID   string `json:"queryid,omitempty"`
      QueryStatus string  `json:querystatus,omitempty`
      QueryTime string `json:querytime,omitempty`
  }

  //Specify the url to be connected and encode the url
	restVal := url.Values{}
  encodeURL := restVal.Encode()
  fmt.Printf("encodeURL.Encode(): %v\n", encodeURL)

  //set the values to be submitted to the live service at getallevent
  toSentInfo := EventInfo{"", "q2", "", ""}
  marshalledInfo, err := json.Marshal(toSentInfo)

  //Driver description here
  //POSTing it with GET_VAL http method after converting to the JSON format
  req, err := http.NewRequest("GET_VAL", connectURL, bytes.NewReader(marshalledInfo))
  if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  //Initializing the client driver and sending the marshalled info
  clientDriver := &http.Client{}
	driv, err := clientDriver.Do(req)
	if err != nil {
		fmt.Printf("Error: Client Driver failed to initialize %v\n", err)
		return
	}
	defer driv.Body.Close()

  //Collecting the values received from the service
	driverInfo, err := ioutil.ReadAll(driv.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

  //Display the collected value in the terminal
  fmt.Printf("Retrived QueryID information from micromonDB successfully:\n%v\n", string(driverInfo))

}
