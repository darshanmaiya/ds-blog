package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"strconv"
	"time"

	"bitbucket.org/darshanmaiya/ds-blog/config"
	"bitbucket.org/darshanmaiya/ds-blog/server"
)

func main() {
	appServer := server.Server{}
	serverID, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("** Starting server with ID %d **\n\n", serverID)

	// Store ID
	appServer.ServerID = serverID
	var err error
	appServer.Participants, err = config.GetServersFromConfig()
	appServer.IpConfig = appServer.Participants[serverID]

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Server %d initialized @ %s\n", appServer.ServerID, appServer.IpConfig)

	fmt.Printf("\nParticipant details: \n")
	for i, value := range appServer.Participants {
		if i != appServer.ServerID {
			fmt.Printf("Server %d @ %s\n", i, value)
		}
	}
	fmt.Println()

	now := time.Now().UTC()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\nReceived %v signal, server %d shutting down...\n", sig, serverID)
			os.Exit(0)
		}
	}()

	//server := new(Server)
	rpc.Register(appServer)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", appServer.IpConfig)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(listener, nil)

	fmt.Println("** Started server ", serverID, " at: ", now.Format(time.UnixDate), "**")
	for {

	}
}
