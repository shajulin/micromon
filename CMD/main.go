// ProjectName:: Micromon ver1.0
//
// ProjectTitle:: micromon services
//
// ProjectDescription:: See doc folder for more information.
//
// Author:: Shajulin Benedict benedict@in.tum.de; shajulin@sxcce.edu.in
//
// PackageName: micromon: main.go - main file for MicromonService
// Supported by: Prof. Dr. Michael Gerndt, TUM, Germany
package main

import (
    "encoding/json"
    "log"
    "fmt"
    "os"
    "strconv"
    "io/ioutil"
    "net/http"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "github.com/gorilla/mux"
    "github.com/micromon/addevent"
    "github.com/micromon/getproctime"
    "github.com/micromon/getallevent"
)
//"github.com/micromon/get_event"

type Event struct {
    ID        string   `json:"id,omitempty"`
    Currenttime string   `json:"currenttime,omitempty"`
    Eventtoken  string  `json:"eventtoken,omitempty"`
}

type EventInfo struct {
    AppID   string `json:"appid,omitempty"`
    QueryID   string `json:"queryid,omitempty"`
    EventName string  `json:eventname,omitempty`
    TimeStamp string `json:timestamp,omitempty`
    AdditionalInfo string `json:additionalinfo,omitempty`
}

var eventlist []Event
var AddEventInfo EventInfo

var outputHtmlCode = `
<!DOCTYPE html>
<html>
<head>
  <title>MicromonDB</title>
  <style>
    body {background-color: powderblue;}
    h1 {color: red;}
    p {color: blue;}
  </style>
</head>
<body>

<h1>Congratulations!!</h1>
<img src="http://www.clipartsmania.com/gif/flower_gif/yellow_flower_animation.gif"
alt="http://www.clipartsmania.com/gif/flower_gif/yellow_flower_animation.gif" style="width:48px;height:48px;">
<p> You have successfully deleted MicromonDB..!</p>

</body>
</html>
`
// dropalleventhandler service drops mrimonDB database
// Invoke them from url:8172/dropall
func dropalleventhandler(w http.ResponseWriter, req *http.Request) {
  switch req.Method {
	case "GET":
		fmt.Fprintf(w, outputHtmlCode)
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
            panic(err)
    }
    defer session.Close()

    // drop database
    IsDrop := true
    if IsDrop {
        err = session.DB("MicromonDB").DropDatabase()
        if err != nil {
            panic(err)
        }
    }
  default:
    fmt.Fprintf(w, "Only dropDB option permitted in this interface !!")
  }
}


/// getproctimehandler interface : Definition
func getproctimehandler(w http.ResponseWriter, req *http.Request) {
  switch req.Method {
	case "GET":
		fmt.Fprintf(w, getproctime.OutputGetProcTimeInfo)

  case "POST":
    //Collect from the other programs
    var eInfo EventInfo
    collectData, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(collectData, &eInfo)

    fmt.Println("Received EventInfo: ", eInfo.AppID, eInfo.QueryID, eInfo.EventName, eInfo.TimeStamp, eInfo.AdditionalInfo)

    ///Now, Identify the values from micromonDB
    session, err := mgo.Dial("localhost:27017")
    c := session.DB("MicromonDB").C("entries")

    if err != nil {
            panic(err)
    }
    defer session.Close()


    type EventInfo struct {
        AppID   string `json:"appid,omitempty"`
        QueryID   string `json:"queryid,omitempty"`
        EventName string  `json:eventname,omitempty`
        TimeStamp string `json:timestamp,omitempty`
        AdditionalInfo string `json:additionalinfo,omitempty`
    }
    // Display query result (single entry)
    result := EventInfo{}
    err = c.Find(bson.M{"appid": eInfo.AppID,"queryid": eInfo.QueryID}).One(&result)

    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("QueryID and TimeStamp are:", result.QueryID, result.TimeStamp)

    ///Now, calculate the difference in the timeframe
    convert_dbTimeStamp, err := strconv.ParseFloat(result.TimeStamp, 64)
    convert_userTimeStamp, err := strconv.ParseFloat(eInfo.TimeStamp, 64)
    diff_val := convert_userTimeStamp - convert_dbTimeStamp
    fmt.Printf("db %f user %f diff_val is %f", convert_dbTimeStamp, convert_userTimeStamp, diff_val)

    ///Now, describe the value in the web
    diffstr := strconv.FormatFloat(diff_val, 'g', -1, 64)
    var OutputGetProcTimeValueInfo = `
    <!DOCTYPE html>
    <html>
    <head>
      <title>MicromonDB</title>
      <style>
        body {background-color: powderblue;}
        h1 {color: red;}
        p {color: blue;}
      </style>
    </head>
    <body>

    <h1>ProcessingTime for QueryID:</h1>

    <img src="http://www.clipartsmania.com/gif/star/animation-red-star.GIF"
    alt="http://www.clipartsmania.com/gif/star/animation-red-star.GIF" style="width:48px;height:48px;">
    <p> !</p>

    </body>
    </html>
    `
    fmt.Fprintf(w, OutputGetProcTimeValueInfo)
    fmt.Fprintf(w, diffstr)

    default:
      fmt.Fprintf(w, "Either set GET or POST operations! Resting now!!")
  }

}

/// Infohandler service - micromon
func infohandler(w http.ResponseWriter, req *http.Request){
   w.Write([]byte("Welcome to Microservice Monitoring/Logging Infrastructure!\n Options are i) /getproctime and ii) /addevent\n"))
}

/// addeventhandler service - micromon
func addeventhandler(w http.ResponseWriter, req *http.Request){
  switch req.Method {
  case "GET":
    fmt.Fprintf(w, addevent.OutputAddEventInfo)

  case "POST":
    collectData, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(collectData, &AddEventInfo)

    fmt.Println("AppID, QueryID, EventName, TimeStamp, and AdditionalInfo are", AddEventInfo.AppID, AddEventInfo.QueryID, AddEventInfo.EventName, AddEventInfo.TimeStamp, AddEventInfo.AdditionalInfo)

    params := mux.Vars(req)
    for _, entry := range eventlist {
        if entry.ID == params["id"] {
            json.NewEncoder(w).Encode(entry)
            return
        }
    }
    json.NewEncoder(w).Encode(&Event{})

    ///CHANGE THIS:::TODO::Directly entering the values to mrimonDB here.
    fmt.Println("Entering JSON values to micromonDB \n")

    session, err := mgo.Dial("localhost:27017")
    if err != nil {
            panic(err)
    }
    defer session.Close()

    // Keep the session monotonic.
    session.SetMode(mgo.Monotonic, true)

    type Event struct {
        ID        string   `json:"id,omitempty"`
        Currenttime string   `json:"currenttime,omitempty"`
        Eventtoken  string  `json:"eventtoken,omitempty"`
    }

  ///Insert dataset to the database (MicormonDB:entries)
    c := session.DB("MicromonDB").C("entries")

    ///TODO:: split the string and put it in db or pass the struct to InsertEntry
    err = c.Insert(&EventInfo{AddEventInfo.AppID, AddEventInfo.QueryID, AddEventInfo.EventName, AddEventInfo.TimeStamp, AddEventInfo.AdditionalInfo})
    if err != nil {
           log.Fatal(err)
    }

    // Display query result (multiple entry)
    entryCount, err := c.Count()
    fmt.Println("Total Entry in MicromonDB database :", entryCount)

    result := EventInfo{}
    iter := c.Find(nil).Limit(entryCount).Iter()
    fmt.Println()
    fmt.Println("AllEntries at addeventhandler: ")
    for iter.Next(&result) {
        fmt.Printf("EventInfo: %v\n", result.AppID + " , " + result.QueryID + " , " + result.EventName + " , " + result.TimeStamp + " , " + result.AdditionalInfo )
    }
    err = iter.Close()
    if err != nil {
        log.Fatal(err)
    }
    default:
      fmt.Fprintf(w, "Set POST operations! Resting now!!")
  }


}


///getallevent handler interface - micromon
func getalleventhandler(w http.ResponseWriter, req *http.Request) {

  switch req.Method {
  case "GET_VAL":
    //fmt.Fprintf(w, getallevent.OutputGetAllEventInfo)

        var allGetInfo EventInfo
        collectAllGetData, _ := ioutil.ReadAll(req.Body)
        json.Unmarshal(collectAllGetData, &allGetInfo)

        fmt.Println("Received Request for SSS QueryID: ", allGetInfo.QueryID)

        ///Now, Identify the values from micromonDB
        session, err := mgo.Dial("localhost:27017")
        c := session.DB("MicromonDB").C("entries")

        if err != nil {
                panic(err)
        }
        defer session.Close()


        type EventInfo struct {
            AppID   string `json:"appid,omitempty"`
            QueryID   string `json:"queryid,omitempty"`
            EventName string  `json:eventname,omitempty`
            TimeStamp string `json:timestamp,omitempty`
            AdditionalInfo string `json:additionalinfo,omitempty`
        }

        entryAllCount, err := c.Count()
        fmt.Println("Total Entry in MicromonDB database :", entryAllCount)

        resultAll := EventInfo{}
        iter := c.Find(bson.M{"queryid": allGetInfo.QueryID}).Limit(entryAllCount).Iter()
        fmt.Println()
        fmt.Println("AllEntries at getalleventhandler: ")

        var jsonAllDisplay string
        for iter.Next(&resultAll) {
          //Marshall value to json format
          marshalAllData := EventInfo{
        		AppID:     resultAll.AppID,
        		QueryID:   resultAll.QueryID,
            EventName:  resultAll.EventName,
            TimeStamp:  resultAll.TimeStamp,
            AdditionalInfo: resultAll.AdditionalInfo,
        	}
        	MarshalAllFinal, err := json.MarshalIndent(marshalAllData, "", "   ")
        	if err != nil {
        		fmt.Println("error:", err)
        	}
        	os.Stdout.Write(MarshalAllFinal)
          jsonAllDisplay = jsonAllDisplay + string(MarshalAllFinal)
        }
        err = iter.Close()

        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("QueryID and TimeStamp at getalleventhandler are:", resultAll.QueryID, resultAll.TimeStamp)

        //Output in the website
        var OutputGetAllEventValuePrint string
        OutputGetAllEventValuePrint = "   <!DOCTYPE html> " +
        "<html> " +
        "<head> " +
        "  <title>MicromonDB</title> " +
        "  <style> " +
        "    body {background-color: powderblue;} " +
        "    h1 {color: red;} " +
        "    p {color: blue;} " +
        "  </style> " +
        "</head> " +
        "<body> " +
        "<h1>All Specific Events at MicromonDB:</h1> " + jsonAllDisplay +
        "<img src=\"http://www.clipartsmania.com/gif/star/animation-red-star.GIF\" " +
        "alt=\"http://www.clipartsmania.com/gif/star/animation-red-star.GIF\" style=\"width:48px;height:48px;\"> " +
        "<p> !</p> " +
        "</body> " +
        "</html> "

        fmt.Fprintf(w, OutputGetAllEventValuePrint)


	case "GET":
		fmt.Fprintf(w, getallevent.OutputGetAllEventInfo)

  case "POST":

    var allInfo EventInfo
    collectAllData, _ := ioutil.ReadAll(req.Body)
    json.Unmarshal(collectAllData, &allInfo)

    fmt.Println("Received Request for AllEvents--EventInfo: ", allInfo.AppID, allInfo.QueryID, allInfo.EventName, allInfo.TimeStamp)

    ///Now, Identify the values from micromonDB
    session, err := mgo.Dial("localhost:27017")
    c := session.DB("MicromonDB").C("entries")

    if err != nil {
            panic(err)
    }
    defer session.Close()


    type EventInfo struct {
        AppID   string `json:"appid,omitempty"`
        QueryID   string `json:"queryid,omitempty"`
        EventName string  `json:eventname,omitempty`
        TimeStamp string `json:timestamp,omitempty`
        AdditionalInfo string `json:additionalinfo,omitempty`
    }

    entryCount, err := c.Count()
    fmt.Println("Total Entry in MicromonDB database :", entryCount)

    result := EventInfo{}
    iter := c.Find(bson.M{"queryid": allInfo.QueryID}).Limit(entryCount).Iter()
    fmt.Println()
    fmt.Println("AllEntries at getalleventhandler: ")

    var jsonDisplay string
    for iter.Next(&result) {
      //Marshall value to json format
      marshalData := EventInfo{
    		AppID:     result.AppID,
    		QueryID:   result.QueryID,
        EventName:  result.EventName,
        TimeStamp:  result.TimeStamp,
        AdditionalInfo: result.AdditionalInfo,
    	}
    	MarshalFinal, err := json.MarshalIndent(marshalData, "", "   ")
    	if err != nil {
    		fmt.Println("error:", err)
    	}
    	os.Stdout.Write(MarshalFinal)
      jsonDisplay = jsonDisplay + string(MarshalFinal)
    }
    err = iter.Close()

    if err != nil {
            log.Fatal(err)
    }
    fmt.Println("QueryID and TimeStamp at getalleventhandler are:", result.QueryID, result.TimeStamp)

    //Output in the website
    var OutputGetAllEventValueInfo string
    OutputGetAllEventValueInfo = "   <!DOCTYPE html> " +
    "<html> " +
    "<head> " +
    "  <title>MicromonDB</title> " +
    "  <style> " +
    "    body {background-color: powderblue;} " +
    "    h1 {color: red;} " +
    "    p {color: blue;} " +
    "  </style> " +
    "</head> " +
    "<body> " +
    "<h1>All Specific Events at MicromonDB:</h1> " + jsonDisplay +
    "<img src=\"http://www.clipartsmania.com/gif/star/animation-red-star.GIF\" " +
    "alt=\"http://www.clipartsmania.com/gif/star/animation-red-star.GIF\" style=\"width:48px;height:48px;\"> " +
    "<p> !</p> " +
    "</body> " +
    "</html> "

    fmt.Fprintf(w, OutputGetAllEventValueInfo)

  default:
    fmt.Fprintf(w, "Either set GET or POST operations! Resting now!!")
  }

}

/**
  MicromonService - main function
*/
func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", infohandler)
    router.HandleFunc("/dropall", dropalleventhandler)
    router.HandleFunc("/addevent", addeventhandler)
    router.HandleFunc("/getproctime", getproctimehandler)
    router.HandleFunc("/getallevent", getalleventhandler)

    log.Fatal(http.ListenAndServe(":8172", router))
}
