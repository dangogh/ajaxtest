package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/time", timeData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start.")
	}
} //main

func timeData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I can't tell time yet.")
}

func HomePage(writer http.ResponseWriter, request *http.Request) {
	output := `
<html>
<body>
<script type="text/javascript">
 
function getTime()
{
var ajaxObject ;
  
// code for IE7+, Firefox, Chrome, Opera, Safari
if (window.XMLHttpRequest)
{
ajaxObject = new XMLHttpRequest();
  
ajaxObject.open("GET","/time",true);
ajaxObject.send();
  
ajaxObject.onreadystatechange = function()
{
if (ajaxObject.readyState == 4 && ajaxObject.status == 200)
{
document.getElementById("time").innerHTML = "works"// ajaxObject.responseText;
  
} } } } </script>

<div id="time"> Click to display current time. </div>
<input id="timeButton" type="button" value="Click" onclick="getTime()" />
`
	fmt.Fprint(writer, output)
}
