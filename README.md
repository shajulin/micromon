# micromon
micromon: A microservice oriented performance monitoring tool for Cloud Applications.


Overall Description: micromon is a golang based logging service.
Detailed Description: Users could submit the event information to micromon
service. This information will be stored in mongodb based database named
micromonDB. Later, they could collect the processing time of events.

The micromon package consists of the
following folders (FOLDER/Source information):

micromon
    -addevent (addevent interface)
      This interface adds events to the micromon service.
      Usage: (from terminal)
      curl -d '{"AppID":"ID-1","QueryID":"q1","EventName":"RegisterStudent","TimeStamp":"23.1","AdditionalInfo":"TUMStudent"}' -i http://localhost:8172/addevent

    -CMD (main)
      This folder contains the main package of micromon service.

    -config (future)
    -doc (documentation)

    -driver (driver interface)
      This folder contains driver.go file. This is an example written in order to
      demonstrate how the micromon service could be accessed and the processing time
      for events could be retrieved from micromonDB.

    -getproctime (driver interface)
      THis interface is responsible for getting the processing time of the events.
      To utilize this service via. terminal,
      curl -d '{"AppID":"ID-1","QueryID":"q2","EventName":"RegisterStudent","TimeStamp":"29.1","AdditionalInfo":"HPCCLoudStudent"}' -i http://localhost:8172/getproctime
      If invoked via. browser, the information about utilizing this service will be displayed.
      In this service, the user has to specify the current timestamp in seconds or milliseconds
      along with the queryid (Mandatory).
      It searches for the first entry of this queryid and calculates the processing time for
      the event and reports the same to the user.

    -getallevent (getallevent interface)
      This interface reports all entries of the events that were added to the mrimondb earlier.
      THe output is reported to the user in the json format.

    -micromonDB (database folder)
      This is the folder which stores the event information.

    -servicediscovery (future)
      This folder is kept for the future use.

How to use micromon service?
    Users could utilize in three different ways: i) via. web ii) via. client terminal iii) via. programs
      i) via. web:
        The services could be contacted using the respective urls. FOr eg. http://localhost:port/addevent
        Only limited services are provided with this option.
      ii) via. client terminal
        The services could be accessed via. the client terminal using curl operations. Pls. see
        above for more information.
      iii) via. programs.
        Users could write a program to utilize micromon service. To do so, the users have to do the
        following in their programs (although the programs could be written in any of the languages):
          a) connect to the service
          b) submit the json encoded EventInfo datastructure in bytes to the services.
          c) collect the output and parse the output as per their needs.
      A sample example is written in golang which is available in the driver folder.
