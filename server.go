package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
)

type JsonReponse struct {
	Stdout  string `json:"stdout"`
	Stderr  string `json:"stderr"`
	Retcode int64  `json:"retcode"`
}

var scriptDir = flag.String("dir", "bin", "The directory where scripts will be located")
var token = flag.String("token", "ABC123", "A token used to obtain API access")
var port = flag.String("port", "8099", "Set the http port to listen to")

func scriptsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.URL.Query().Get("token") == *token {

		jsonRes := JsonReponse{}

		cmdName := r.URL.Path[len("/scripts/"):]
		command := *scriptDir + "/" + cmdName + ".sh"
		cmd := exec.Command(command)
		r.ParseForm()

		for k, v := range r.Form {
			cmd.Env = append(cmd.Env, k+"="+v[0])
		}
		out, err := cmd.Output()

		if err != nil {
			jsonRes.Stderr = err.Error()
			jsonRes.Retcode = 1
		}

		jsonRes.Stdout = string(out)

		b, err := json.Marshal(jsonRes)

		fmt.Fprintln(w, string(b))
	}

	if r.Method == "GET" {
		// TODO Read files in script dir
		// parse comment section of the script
		// return an array of available scripts,
		// description, and required params
	}

	if r.URL.Query().Get("token") != *token {
		error := "Invalid token"
		http.Error(w, error, 401)
	}

}

func main() {
	flag.Parse()
	http.HandleFunc("/scripts/", scriptsHandler)
	http.ListenAndServe(":"+*port, nil)
}
