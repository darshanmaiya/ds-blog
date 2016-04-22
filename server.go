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

	"github.com/darshanmaiya/ds-blog/serverDS"
	"github.com/spf13/viper"
)

func (server serverDS.Server) startServer(id int) {
	// Store ID
	server.ServerID = id
	server.Participants = make(map[int]string)

	viper.SetConfigName("conf")
	viper.AddConfigPath("../conf")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("No configuration file loaded - exiting...")
		os.Exit(1)
	}

	allServers := viper.GetStringMap("Servers")

	for i, value := range allServers {
		servId, _ := strconv.Atoi(i)
		if servId == server.ServerID {
			server.IpConfig = value.(string)

			continue
		}

		server.Participants[servId] = value.(string)
	}

	fmt.Printf("Server %d initialized @ %s\n", server.ServerID, server.IpConfig)

	fmt.Printf("\nParticipant details: \n")
	for i, value := range server.Participants {
		if i != server.ServerID {
			fmt.Printf("Server %d @ %s\n", i, value)
		}
	}
	fmt.Println()
}

func (server serverDS.Server) PostMessage(args *serverDS.PostArgs, reply *serverDS.PostReply) error {
	logMsg := LogMsg{
		ID:           len(server.Log),
		Message:      args.Message,
		ReplyTo:      -1,
		InfluencedBy: -1,
		Timestamp:    time.Now().UTC().Unix(),
		ServerID:     server.ServerID,
	}

	server.Log = append(server.Log, logMsg)

	reply.Reply = "Success"

	return nil
}

func (server Server) Lookup(args *LookupArgs, reply *LookupReply) error {
	for _, logMsg := range server.Log {
		reply.Messages = append(reply.Messages, logMsg)
	}

	return nil
}

func (server Server) Sync(args *SyncArgs, reply *SyncReply) error {
	fmt.Println("Sync from server ", args.SyncFromServer)
	reply.Reply = "Success"

	return nil
}

func main() {
	server := Server{}
	serverID, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("** Starting server with ID %d **\n\n", serverID)

	server.startServer(serverID)

	now := time.Now().UTC()
	fmt.Println("** Started server ", serverID, " at: ", now.Format(time.UnixDate), "**")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\nReceived %v signal, server %d shutting down...\n", sig, serverID)
			os.Exit(0)
		}
	}()

	//server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", server.IpConfig)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(listener, nil)

	for {

	}
}
