package main

import (
	"bytes"
    "encoding/json"
    "fmt"
    "net/http"
	"io/ioutil"
    "log"
)

var apibase string

//===========================================
type ConfigData struct 
{
  
    Api      string
}

//===========================================
func sendPostRequest(urlpath string,apikey string) string{

   
}

//===========================================
func sendGetRequest(urlpath string,apikey string) string{

   
   req, err := http.NewRequest("GET", apibase+urlpath, nil)
	if err != nil {

	}

	req.Header.Set("X-API-KEY", apikey)
	client := &http.Client{}
	resp, err := client.Do(req)
	
	
  body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     log.Fatalln(err)
   }
   sb := string(body)
   return sb;
   
   
}

//===========================================
func base(w http.ResponseWriter, req *http.Request){


 fmt.Printf("Req: %s %s\n", req.Host, req.URL.Path); 
 
apikey := req.Header.Get("X-API-KEY")
	
	 fmt.Printf(req.Method);
	
	if req.Method == http.MethodGet {
        	fmt.Fprintf(w, sendGetRequest(req.URL.Path,apikey));
    } else if req.Method == http.MethodPost {
        	fmt.Fprintf(w, sendPostRequest(req.URL.Path,apikey));
    } else {
       
    }
	
	
}


//===========================================
func loadConfig(){

   var data ConfigData
    file, err := ioutil.ReadFile("config.json")
    if err != nil {
        log.Fatal(err)
    }
    err = json.Unmarshal(file, &data)
    if err != nil {
        log.Fatal(err)
    }
    
	apibase = data.Api;
}


//===========================================
func main(){


	loadConfig();
    http.HandleFunc("/", base)	
    http.ListenAndServe(":8090", nil)
}