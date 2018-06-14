package kaizen

type Task struct {
    Script string
}

func NewTaskFromScript(script string) *Task {
    task := Task {
        Script: script,
    }

    return &task
}

func (task *Task) Run() error {
    return nil
}
