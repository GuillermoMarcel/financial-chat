package models

type Config struct {
	BrokerAddress      string
	CmdQueue           string
	ResponsesQueue     string
	DatabaseLocation   string
	Port               int
	InitializeDatabase bool
}
