/**
micromon::: getallevent.go --> gets all event from micromonDB
Shajulin Benedict
*/
package getallevent

import(
  _"fmt"
  _"log"
  _"gopkg.in/mgo.v2"
  _"gopkg.in/mgo.v2/bson"
)

/// Output getallevent info here.
var OutputGetAllEventInfo = `
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
<p> Please enter the AppID, QueryID, EventName, TimeStamp, and AdditionalInfo inorder to get all event data from MicromonDB..!</p>

</body>
</html>
`
