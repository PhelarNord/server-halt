package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main()  {
	http.Handle("/",http.FileServer(http.Dir("./public")))

	haltMeHandler := func(w http.ResponseWriter, rec *http.Request) {
		io.WriteString(w, "Server is shutting down...")

		halt := "sudo halt"
		cmd := exec.Command(halt)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
	}

	http.HandleFunc("/halt", haltMeHandler)
	log.Println("Listening for halt at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
	
}