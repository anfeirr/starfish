package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

func main()  {
	http.HandleFunc("/",deploy)
	http.ListenAndServe("8001",nil)

	
}

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