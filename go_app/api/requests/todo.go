package requests

type AddTodo struct {
	Title string `validate:"required"`
}

type UpdateTodo struct {
	ID        int
	Title     string
	Completed bool
}
