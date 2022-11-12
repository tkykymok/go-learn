package requests

import "github.com/volatiletech/null/v8"

type AddTodo struct {
	Title null.String
}

type UpdateTodo struct {
	ID        int
	Title     null.String
	Completed bool
}
