TCP chat hub

this is a simple Tcp Chat Hub Implementation
it runs on port 8080
to start the hub just say 
go run Hub.go

Once Connected Client can send 3 types of JSON messages
{"messageType":"ID"}
For this message the Server returns a Unique Integer assigned by the server eg.  5577006791947779410

{"messageType":"LIST"}
For this message the Server returns a Comma Seperated list of Connected Clients ids 

{"messageType":"RELAY", "receiverIds":[5577006791947779410], "messageBody":"Hello Dude"}
For this message the server will pass on the message to the clients listed in the receiver ids list 