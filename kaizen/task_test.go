package kaizen

import (
    "testing"
)


func TestEmptyTaskCanBeExecuted(t *testing.T) {
    emptyScript := ""
    task := NewTaskFromScript(emptyScript)
    oops := task.Run()
    if oops != nil {
        t.Error("Task could not be run")
    }
}
