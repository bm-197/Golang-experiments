package main

import (
    "flag"
    "os"
    "strings"
    "fmt"
    "strconv"
)

type CMDFlags struct {
    Add string
    Delete int
    Edit string
    Toggle int
    List bool
}

func NewCMDFlags() *CMDFlags {
    cf := CMDFlags{}

    flag.StringVar(&cf.Add, "add", "", "Add a new todo specify Task")
    flag.StringVar(&cf.Edit, "edit", "", "Edit a task by index & specify a new title, id:new_title") 
    flag.IntVar(&cf.Delete, "del", -1, "Delete a task by index")
    flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a task by index to toggle")
    flag.BoolVar(&cf.List, "list", false, "List all tasks")

    flag.Parse()

    return &cf
}

func (cf *CMDFlags) excute(todos *Todos) {
    switch {
    case cf.List:
        todos.print()
    case cf.Add != "":
        todos.add(cf.Add)
    case cf.Edit != "":
        parts := strings.SplitN(cf.Edit, ":", 2)

        if len(parts) != 2 {
            fmt.Println("Error, Invalid format for edit, please use id:new_title")
        }
        index, err := strconv.Atoi(parts[0])

        if err != nil {
            fmt.Println("Error: Invalid index for edit")
            os.Exit(1)
        }
        
        todos.edit(index, parts[1])

    case cf.Toggle != -1:
        todos.toggle(cf.Toggle)
    case cf.Delete != -1:
        todos.delete(cf.Delete)
    default:
        fmt.Println("Invalid Command")

    }
}
