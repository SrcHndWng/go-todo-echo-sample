package repository

type (
	Todo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	todos = make([]Todo, 0)
	seq   = 1
)

func NewTodo() *Todo {
	t := &Todo{
		ID: seq,
	}
	seq++
	return t
}

func AddTodo(t *Todo) {
	todos = append(todos, *t)
}

func GetTodos() []Todo {
	return todos
}

func GetTodo(id int) *Todo {
	return selectTodo(id)
}

func UpdateTodo(id int, t *Todo) *Todo {
	update := selectTodo(id)
	if update == nil {
		return nil
	}
	update.Name = t.Name
	return update
}

func DeleteTodo(id int) []Todo {
	return removeToto(id)
}

func removeToto(id int) []Todo {
	isExist := false
	result := make([]Todo, 0)
	for _, t := range todos {
		if id == t.ID {
			isExist = true
		} else {
			result = append(result, t)

		}
	}
	if isExist == false {
		return nil
	}
	todos = result
	return todos
}

func selectTodo(id int) *Todo {
	for _, t := range todos {
		if id == t.ID {
			return &t
		}
	}
	return nil
}
