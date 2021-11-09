package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	// switch handler part:
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			fmt.Println("myip returns your ip adderess")
			fmt.Println("Usage: myip [OPTIONS: [-h, --help: for help]]")
			return
		default:
			log.Fatal("invalid input: ", os.Args[1:])
		}
	}

	// get ip part:
	res, err := exec.Command("bash", "-c", "curl icanhazip.com || curl -4 icanhazip.com || curl -6 icanhazip.com").Output()
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Print(string(res))

}
