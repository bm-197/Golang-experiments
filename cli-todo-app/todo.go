package main

import (
    "fmt"
    "time"
    "errors"
    "os"
    "github.com/aquasecurity/table"
    "strconv"
)

type Todo struct {
    Title string
    Completed bool
    CreatedAt time.Time
    CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
    todo := Todo{
        Title: title,
        Completed: false,
        CreatedAt: time.Now(),
        CompletedAt: nil,
    }
    
    *todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
    if index < 0 || index >= len(*todos) {
        err := errors.New("Invalid Index")
        fmt.Println(err)
        return err
    }
    return nil
}

func (todos *Todos) delete(index int) error {
    t := *todos

    if err := t.validateIndex(index); err != nil {
        return err
    }

    *todos = append(t[:index], t[index+1:]...)

    return nil
}

func (todos *Todos) toggle(index int) error {
    t := *todos
    
    if err := t.validateIndex(index); err != nil {
        return err
    }
    
    isCompleted := t[index].Completed

    if !isCompleted {
        completionTime := time.Now()
        t[index].CompletedAt = &completionTime
    }

    t[index].Completed = !isCompleted

    return nil
}

func (todos *Todos) edit(index int, newTitle string) error {
    t := *todos
    
    if err := t.validateIndex(index); err != nil {
        return err
    }
    
    t[index].Title = newTitle
    
    return nil
}

func (todos *Todos) print() {
    table := table.New(os.Stdout)
    table.SetRowLines(false)
    table.SetHeaders("#", "Title", "Done", "Created At", "Completed At")

    for i, t := range *todos {
        done := "❌"
        completedAt := ""

        if t.Completed {
            done = "✅"
            if t.CompletedAt != nil {
                completedAt = t.CompletedAt.Format(time.RFC1123)
            }
        }

        table.AddRow(strconv.Itoa(i), t.Title, done, t.CreatedAt.Format(time.RFC1123), completedAt)
    }
    table.Render()
}












