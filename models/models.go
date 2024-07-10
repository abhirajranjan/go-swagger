package models

// todo represents data about a task in the todo list
type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
}

// message represents request response with a message
type Message struct {
	Message string `json:"message"`
}

type NewMessage struct {
	TodoId int
	Task   string
}
