package requests

type AddTodo struct {
	Title string `validate:"required,title-custom"`
}

type UpdateTodo struct {
	ID        int    `validate:"required"`
	Title     string `validate:"required,title-custom"`
	Completed bool
}

type GetTodosWithRelated struct {
	ID     int
	UserId int
}
