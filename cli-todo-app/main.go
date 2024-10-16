package main


func main() {
    todos := Todos{}
    storage := NewStorage[Todos]("todos.json")
    storage.Load(&todos)
    cmdFlags := NewCMDFlags()
    cmdFlags.excute(&todos)
    storage.Save(todos)
}

