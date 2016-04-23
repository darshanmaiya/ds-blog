package server

import (
	"fmt"
	"strconv"
	"time"
)

type LogMsg struct {
	ID           int
	Message      string
	ReplyTo      int
	InfluencedBy int
	Timestamp    int64
	ServerID     int
}

type Server struct {
	ServerID int
	IpConfig string

	Log          []LogMsg
	Timetable    [][]int
	Participants map[int]string
}

type PostArgs struct {
	Message string
}

type PostReply struct {
	Reply string
}

type LookupArgs struct {
}

type LookupReply struct {
	Messages []LogMsg
}

type SyncArgs struct {
	SyncFromServer int
}

type SyncReply struct {
	Reply string
}

func (server Server) PostMessage(args *PostArgs, reply *PostReply) error {
	logMsg := LogMsg{
		ID:           len(server.Log),
		Message:      args.Message,
		ReplyTo:      -1,
		InfluencedBy: -1,
		Timestamp:    time.Now().UTC().Unix(),
		ServerID:     server.ServerID,
	}

	server.Log = append(server.Log, logMsg)

	reply.Reply = "Success. Total number of messages is: " + strconv.Itoa(len(server.Log))

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
