package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"errors"
)

var correctPort = regexp.MustCompile("^\\:[0-9]+$")

func main() {
	port, err := getPort(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		return
	}
	http.HandleFunc("/", makeStaticPageHandler("home"))
	http.ListenAndServe(port, nil)
}

func getPort(args []string) (string, error) {
	if len(args) > 2 {
		return "", errors.New("Too many arguments!")
	}
	var port string
	if len(args) == 1 {
		port = ":80"
	} else if len(args) == 2 {
		port = args[1]
	}
	if !(correctPort.Match([]byte(port))) {
		return "", errors.New("That is not a port!")
	}
	return port, nil
}

func printUsage() {
	fmt.Println("Usage: TicTacToeWebsite [:port]")
}