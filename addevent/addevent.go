/**
micromon::: addevent.go adds event to a database
Shajulin Benedict
*/

package addevent

import(
  _"fmt"
  _"log"
  _"gopkg.in/mgo.v2"
  _"net/http"
)

//var localAddEventInfo EventInfo

/// Output addevent info here.
var OutputAddEventInfo = `
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
<p> Please enter the AppID, QueryID, EventName, TimeStamp, and AdditionalInfo
values inorder to add event data from MicromonDB..!</p>
<p> QueryStatus are i) startmon or ii) stopmon </p>

</body>
</html>
`
/*
/// insertEntry
func InsertEntry(Entryval string){
   fmt.Println("Entering string to micromonDB \n", Entryval)

   session, err := mgo.Dial("localhost:27017")
   if err != nil {
           panic(err)
   }
   defer session.Close()

   // Keep the session monotonic.
   session.SetMode(mgo.Monotonic, true)

   //localEvent := github.micromon.Event{"ID","sdfsdfd","sdfsdf"}

   type Event struct {
       ID        string   `json:"id,omitempty"`
       Currenttime string   `json:"currenttime,omitempty"`
       Eventtoken  string  `json:"eventtoken,omitempty"`
   }


 ///Insert dataset to the database (MicormonDB:entries)
   c := session.DB("MicromonDB").C("entries")

   ///TODO:: split the string and put it in db or pass the struct to InsertEntry
   err = c.Insert(&Event{"ID-1", "2:00", "student"},
                   &Event{"ID-2", "3:00", "employee"})
   if err != nil {
          log.Fatal(err)
   }
   entryCount, err := c.Count()
   fmt.Println(" Total Entry in MicromonDB database :", entryCount)

   // Display query result (multiple entry)
   result := Event{}
   iter := c.Find(nil).Limit(entryCount).Iter()
   fmt.Println()
   fmt.Println("Display query result for all entries ")
   for iter.Next(&result) {
       fmt.Printf("Event: %v\n", result.ID + " , " + result.Currenttime + " , " + result.Eventtoken )
   }
   err = iter.Close()
   if err != nil {
       log.Fatal(err)
   }


}
*/
