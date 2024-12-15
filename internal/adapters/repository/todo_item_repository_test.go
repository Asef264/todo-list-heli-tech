package repository

import (
	"context"
	"testing"
	"time"
	"todo-list/internal/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoItem(t *testing.T) {
	// mock database and mock object
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error opening mock database: %v", err)
	}
	defer db.Close()

	todo := domain.TodoItem{
		ID:          "new-id",
		Description: "Test Task",
		DueDate:     time.Now().Add(24 * time.Hour),
		FileID:      "test-file-id",
		CreatedAt:   time.Date(2024, 12, 0, 0, 0, 0, 0, time.Local),
	}
	//query on mocked database
	mock.ExpectQuery(`INSERT INTO todo_items \(id, description, due_date, file_id\) VALUES \(\$1, \$2, \$3, \$4\) 
	RETURNING id, description, due_date, file_id, created_at`).
		WithArgs(todo.ID, todo.Description, todo.DueDate, todo.FileID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "description", "due_date", "file_id", "created_at"}).
			AddRow(todo.ID, todo.Description, todo.DueDate, todo.FileID, todo.CreatedAt))

	todoItemRepo := NewTodoItem(db)
	// Call the CreateTodoItem method
	result, err := todoItemRepo.CreateTodoItem(context.Background(), todo)
	assert.NoError(t, err)

	// Assert that the returned TodoItem matches the expected values
	assert.NotNil(t, result)
	assert.Equal(t, todo.ID, result.ID)
	assert.Equal(t, todo.Description, result.Description)
	assert.Equal(t, todo.DueDate, result.DueDate)
	assert.Equal(t, todo.FileID, result.FileID)
	assert.Equal(t, todo.CreatedAt, result.CreatedAt)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}
