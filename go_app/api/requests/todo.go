package requests

type AddTodo struct {
	Title string `validate:"required,title-custom"`
}

type UpdateTodo struct {
	ID        int
	Title     string
	Completed bool
}
