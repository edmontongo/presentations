package main
import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

type TodoApp struct {
	todos   []Todo
	nextID  int
	mutex   sync.RWMutex
}

func NewTodoApp() *TodoApp {
	return &TodoApp{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

func (app *TodoApp) AddTodo(task string) {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	todo := Todo{
		ID:        app.nextID,
		Task:      task,
		Completed: false,
	}
	app.todos = append(app.todos, todo)
	app.nextID++
}

func (app *TodoApp) ToggleTodo(id int) {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	for i := range app.todos {
		if app.todos[i].ID == id {
			app.todos[i].Completed = !app.todos[i].Completed
			break
		}
	}
}

func (app *TodoApp) DeleteTodo(id int) {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	for i, todo := range app.todos {
		if todo.ID == id {
			app.todos = append(app.todos[:i], app.todos[i+1:]...)
			break
		}
	}
}

func (app *TodoApp) GetTodos() []Todo {
	app.mutex.RLock()
	defer app.mutex.RUnlock()

	todos := make([]Todo, len(app.todos))
	copy(todos, app.todos)
	return todos
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>Go Todo App</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px; }
		.todo-item { display: flex; align-items: center; margin: 10px 0; padding: 10px; border: 1px solid #ddd; }
		.completed { text-decoration: line-through; color: #888; }
		button { margin-left: 10px; padding: 5px 10px; }
		input[type="text"] { flex: 1; padding: 8px; margin-right: 10px; }
		.add-form { display: flex; margin-bottom: 20px; }
	</style>
</head>
<body>
	<h1>Todo App</h1>

	<form class="add-form" method="POST" action="/add">
		<input type="text" name="task" placeholder="Enter a new task" required>
		<button type="submit">Add Todo</button>
	</form>

	<div>
		{{range .}}
		<div class="todo-item">
			<span class="{{if .Completed}}completed{{end}}">{{.Task}}</span>
			<form method="POST" action="/toggle" style="display: inline;">
				<input type="hidden" name="id" value="{{.ID}}">
				<button type="submit">{{if .Completed}}Undo{{else}}Complete{{end}}</button>
			</form>
			<form method="POST" action="/delete" style="display: inline;">
				<input type="hidden" name="id" value="{{.ID}}">
				<button type="submit">Delete</button>
			</form>
		</div>
		{{end}}
	</div>
</body>
</html>
`

func main() {
	app := NewTodoApp()
	tmpl := template.Must(template.New("todo").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		todos := app.GetTodos()
		tmpl.Execute(w, todos)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			task := r.FormValue("task")
			if task != "" {
				app.AddTodo(task)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			idStr := r.FormValue("id")
			if id, err := strconv.Atoi(idStr); err == nil {
				app.ToggleTodo(id)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			idStr := r.FormValue("id")
			if id, err := strconv.Atoi(idStr); err == nil {
				app.DeleteTodo(id)
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}