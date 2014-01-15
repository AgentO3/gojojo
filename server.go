package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func scriptsHandler(w http.ResponseWriter, r *http.Request) {
	cmdName := r.URL.Path[len("/scripts/"):]
	args := r.FormValue("args")
	command := "bin/" + cmdName + ".sh"
	out, err := exec.Command(command, args).Output()

	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	fmt.Fprintln(w, string(out))

}

func main() {
	http.HandleFunc("/scripts/", scriptsHandler)
	http.ListenAndServe(":8099", nil)
}
