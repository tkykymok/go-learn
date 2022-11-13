package requests

type AddTodo struct {
	Title string `json:"title" validate:"required"`
}

type UpdateTodo struct {
	ID        int
	Title     string
	Completed bool
}
