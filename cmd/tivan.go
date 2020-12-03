package main

import (
	"os/signal"
	"os"
	"flag"
	"log"
	"fmt"
	"github.com/YangShuting/tivan"
)

var fConfig = flag.String("config", "", "config file to load")

func main(){
	flag.Parse()

	var config *Config
	var err error

	if *fConfig != ""{
		config, err = LoadConfig(*fConfig)
		if err != nil{
			log.Fatal(err)
		}
	} else {
		config = DefaultConfig()
	}

	fmt.Printf("config: %+v\n", config)

	signals := make(chan os.Signal, 1)

	shutdown := make(chan struct{})

	signal.Notify(signals, os.Interrupt)

	go func(){
		<-signals
		log.Println("Tivan shutdowns.")
		close(shutdown)
	}()

	a := &Agent{HTTP: ""}
	
	a.Run(shutdown)

}