package main

import(

	"fmt"
	"time"
	"net/http"
)

func say(s string, done chan string){

	for i := 0;i < 5;i++{
		time.Sleep( 100 * time.Millisecond)
		fmt.Println(s)

	}
	done <- "Terminei"
}


func main(){

	url := "https://medium.com/trainingcenter/goroutines-e-go-channels-f019784d6855"
	resp, err := http.Get(url)
	if err != nil{
		fmt.Println("Erro na URL")
	} 
	body := resp.Body
	fmt.Println(body)
	


}