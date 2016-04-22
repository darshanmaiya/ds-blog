package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"github.com/darshanmaiya/ds-blog/server"
)

func parseServers() {

}

func main() {
	fmt.Println("Welcome to DS-Blog by Syncinators.\nType 'help' for a list of supported commands.\n")

	for {
		consoleReader := bufio.NewReader(os.Stdin)

		fmt.Println("\nEnter command: ")
		command, _ := consoleReader.ReadString('\n')

		command = command[0 : len(command)-1]

		//fmt.Println("\nCommand: " + command)

		input := strings.Split(command, " ")

		client, err := rpc.DialHTTP("tcp", "127.0.0.1:50000")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		args := &server.Args{7, 8}
		var reply int
		err = client.Call("Server.Sync", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

		switch strings.ToLower(input[0]) {
		case "post":
			message := command[5:len(command)]
			//postMessage(message)
			fmt.Printf("Message posted: %s\n", message)

		case "lookup":
			//lookup()

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
