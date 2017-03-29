/**
micromon::: getproctime.go --> gets processingtime from micromonDB
Shajulin Benedict
*/
package getproctime

import(
  _"fmt"
  _"log"
  _"gopkg.in/mgo.v2"
  _"gopkg.in/mgo.v2/bson"
)

/// Output getproctime info here.
var OutputGetProcTimeInfo = `
<!DOCTYPE html>
<html>
<head>
  <title>EventInfo</title>
  <style>
    body {background-color: powderblue;}
    h1 {color: red;}
    p {color: blue;}
  </style>
</head>
<body>

<h1>Ooopps!!! Do correctly!!</h1>

<img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
<p> Please enter the AppID, QueryID, EventName, TimeStamp, and AdditionalInfo inorder to get event data from MicromonDB..!</p>

</body>
</html>
`

/*
///QueryEntry function queries data from micromonDB
func QueryEntry(Queryval string){
  fmt.Println("Quering string from micromonDB \n", Queryval)

  session, err := mgo.Dial("localhost:27017")
  c := session.DB("MicromonDB").C("entries")

  if err != nil {
          panic(err)
  }
  defer session.Close()

  // Keep the session monotonic.
  //session.SetMode(mgo.Monotonic, true)

  type Event struct {
      ID        string   `json:"id,omitempty"`
      Currenttime string   `json:"currenttime,omitempty"`
      Eventtoken  string  `json:"eventtoken,omitempty"`
  }

  type EventInfo struct {
      AppID   string `json:"appid,omitempty"`
      QueryID   string `json:"queryid,omitempty"`
      QueryStatus string  `json:querystatus,omitempty`
      QueryTime string `json:querytime,omitempty`
  }

  // Display query result (single entry)
  result := EventInfo{}
  err = c.Find(bson.M{"AppID": "ID-1"}).One(&result)
  if err != nil {
          log.Fatal(err)
  }
  fmt.Println("QueryID and QueryTime are:", result.QueryID, result.QueryTime)

}
*/
