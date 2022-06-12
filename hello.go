package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

// import "rsc.io/quote"



type Response struct {
	Timezone	string	`json:"timezone"`
	Flag		bool	`json:"shouldideploy"`
	Message		string 	`json:"message"`
}


func main() {

	response, err := http.Get("https://shouldideploy.today/api?tz=UTC")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    //fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	content := responseObject.Message
	//fmt.Println(">> " + responseObject.Message)











	// content := quote.Go()
	size := len(content) + 2

	fmt.Print("  ")
	for i := 1; i <= size; i++ {
		fmt.Print("-")
	}
	fmt.Println(" ");

	fmt.Println(" | " + content + " | ")


	fmt.Print("  ")
	for i := 1; i <= size; i++ {
		fmt.Print("-")
	}
	fmt.Println(" ")




	fmt.Println("  \\   ^__^")
	fmt.Println("   \\  (oo)\\_______")
	fmt.Println("      (__)\\       )\\/\\")
	fmt.Println("          ||----w |")
	fmt.Println("          ||     ||")
	fmt.Println(" ")

}

