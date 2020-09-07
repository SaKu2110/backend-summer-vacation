package model

// SQL QUERY
const(
	QUERY_FORMAT_GET_USER = "SELECT id, password FROM `users` WHERE id = ?"
	QUERY_FORMAT_SET_USER = "INSERT INTO `users` (id, password) VALUES (?, ?)"
)

type User struct {
	Name	string `json:"name"`
}

// TASK1 RESPONSE BODY
type TimeStamp struct {
	TimeStamp	string	`json:"timestamp"`
	Detail		Detail	`json:"datail"`
}

type Detail struct {
	Date	string	`json:"date"`
	Time	string	`json:"time"`
}

// TASK2 REQUEST BODY
type ZellerElements struct {
	Year	int	`json:"year"`
	Month	int	`json:"month"`
	Day		int	`json:"day"`
}

// TASK3 REQUEST BODY
type DBUser struct {
	ID			string
	Password	string
}
