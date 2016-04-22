package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/darshanmaiya/ds-blog/config"
	"bitbucket.org/darshanmaiya/ds-blog/server"
)

var allServers map[int]string
var client *rpc.Client

func main() {
	fmt.Println("Welcome to DS-Blog by Syncinators.\nType 'help' for a list of supported commands.\n")

	fmt.Println("Initializing available servers. Please wait...")

	var err error
	allServers, err = config.GetServersFromConfig()

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Servers initialized successfully.")
	listServers(allServers)

	for {
		consoleReader := bufio.NewReader(os.Stdin)

		fmt.Println("\nEnter command: ")
		command, _ := consoleReader.ReadString('\n')

		command = command[0 : len(command)-1]

		//fmt.Println("\nCommand: " + command)

		input := strings.Split(command, " ")

		/*client, err := rpc.DialHTTP("tcp", "127.0.0.1:50000")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		args := &server.Args{7, 8}
		var reply int
		err = client.Call("Server.Sync", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)*/

		switch strings.ToLower(input[0]) {
		case "connect":
			var serverId int
			serverId, err = strconv.Atoi(input[1])
			fmt.Printf("Connecting to server %d at %s, please wait...\n", serverId, allServers[serverId])
			err = connectToServer(serverId)

			if err != nil {
				fmt.Println("Connecting to server failed")
			} else {
				fmt.Println("Connected")
			}
		case "post":
			message := command[5:len(command)]
			args := &server.PostArgs{
				Message: message,
			}
			reply := server.PostReply{}
			err = client.Call("Server.PostMessage", args, &reply)
			if err != nil {
				log.Fatal("Server error:", err)
			}
			fmt.Printf("Server replied: %s\n", reply.Reply)

		case "lookup":
			args := &server.LookupArgs{}
			reply := server.LookupReply{}
			err = client.Call("Server.Lookup", args, &reply)
			if err != nil {
				log.Fatal("Server error:", err)
			}
			fmt.Println("Total number of messages in server: ", len(reply.Messages))
			printLogMessages(reply.Messages)

		case "sync":
		/*	fromServerID, err := strconv.Atoi(input[1])
			if err != nil {
				runes_array := []rune(input[1])
				fromServerID = int(runes_array[0]) - 'A' + 1
			}
			sync(fromServerID)*/

		case "help":
			fmt.Println("Here's a list of commands supported by the application:\n")

			fmt.Println("\nConnect commands\n--------------------------------")
			fmt.Println("connect <ID> - Connect to server with given ID")
			fmt.Println("disconnect - Disconnect from the server currently connected to")
			fmt.Println("list - List available servers to connect")
			fmt.Println("which - List details of server currently connected to")

			fmt.Println("\nDS-Blog commands\n--------------------------------")
			fmt.Println("post <Message> - Post a message in DS-Blog")
			fmt.Println("lookup - Display the posts in DS-Blog in causal order")
			fmt.Println("sync <Datacenter_ID> - Synchronize with Datacenter having the specified ID")

			fmt.Println("\nMisc. commands\n--------------------------------")
			fmt.Println("help - Display all supported commands")
			fmt.Println("exit - Exit application")

		case "exit":
			fmt.Println("\nDS-Blog shutting down...\nBye :)")
			os.Exit(0)
		}
	}
}

func listServers(serversList map[int]string) {
	fmt.Println("Available servers are:\n")

	for i, value := range serversList {
		fmt.Printf("Server %d @ %s\n", i, value)
	}
	fmt.Println()
}

func connectToServer(serverID int) error {
	var err error
	client, err = rpc.DialHTTP("tcp", allServers[serverID])

	if err != nil {
		log.Fatal("dialing: ", err)
		return err
	}

	return nil
}

func printLogMessages(messages []server.LogMsg) {
	for _, value := range messages {
		fmt.Printf("ID: %d, Message: \"%s\"\n", value.ID, value.Message)
	}
}
