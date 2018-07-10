package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"

)
func deploy(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"<h1> Deploy is starting </h1>")
	reLaunch()

}

func reLaunch(){
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err !=nil{
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func main()  {
	http.HandleFunc("/",deploy)
	http.ListenAndServe(":5000",nil)

	
}

