package serverDS

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
