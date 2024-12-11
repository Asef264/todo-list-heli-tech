package ports

const (
	CreateTodoItemQuery = `INSERT INTO todo_items (id, description, due_date, file_id) 
			  VALUES ($1, $2, $3, $4) RETURNING id, description, due_date, file_id`
)
