package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"io/ioutil"
	
	"net/http"

	"github.com/monzo/typhon"
)

func ping(req typhon.Request) typhon.Response {
	return req.Response("pong")
	// log.Println("you made a request to ping")
}

func api(req typhon.Request) typhon.Response {
	var url string = "https://www.googleapis.com/books/v1/volumes?q=ruby+programming"

	// response, err := http.Get(url)
	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n, err")
	// 	return req.Response("There was a problem with google books api")
	// } 
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(data))
	// 	// log.Println(string(data))
	// 	return req.Response(string(data))

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	// json := string(body)

	return req.Response(body)
}

func main() {
	router := typhon.Router{}
	router.GET("/ping", ping)
	router.GET("/apiSearch", api)

	svc := router.Serve().
		Filter(typhon.ErrorFilter).
		Filter(typhon.H2cFilter)
	srv, err := typhon.Listen(svc, ":8000")
	if err != nil {
		panic(err)
	}
	log.Printf("👋  Listening on %v", srv.Listener().Addr())

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Printf("☠️  Shutting down")
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Stop(c)
}