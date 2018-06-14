package kaizen

import (
    "testing"
    "fmt"
)


func TestEmptyTaskCanBeExecuted(t *testing.T) {
    emptyScript := ""
    task := NewTaskFromScript(emptyScript)
    oops := task.Run()
    if oops != nil {
        t.Error("Empty task could not be run")
    }
}

func TestTaskCanBeExecuted(t *testing.T) {
    textScript := `
    x = 12
    y = 10
    print(x+y)
    `
    task := NewTaskFromScript(textScript)
    oops := task.Run()
    if oops != nil {
        t.Error("Task from text could not be run")
    }
}

func TestWrongTaskShouldNotBeExecuted(t *testing.T) {
    wrongScript := `
    function f(x)
        return x*x
    end
    print(g(10))
    `
    task := NewTaskFromScript(wrongScript)
    oops := task.Run()
    if oops == nil {
        t.Error("Wrong task could be run")
    } else {
        fmt.Printf("%v", oops)
    }
}
