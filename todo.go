package todo

type TodoList struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"desctiptions"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriptions string `json:"desctiptions"`
	Done         bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
