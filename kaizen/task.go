package kaizen

import (
    "github.com/Shopify/go-lua"
    "os"
    "path/filepath"
    "math/rand"
    "encoding/hex"
)

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
    tempSrc := TempFileName(".", ".lua")
    if fp, oops := os.Create(tempSrc); oops != nil {
        return oops
    } else {
        fp.WriteString(task.Script)
        fp.Close()
    }

    l := lua.NewState()
    lua.OpenLibraries(l)
    if oops := lua.DoFile(l, tempSrc); oops != nil {
        return oops
    }

    if oops := os.Remove(tempSrc); oops != nil {
        return oops
    }

    return nil
}

// Generates a temporary filename for use in testing or whatever.
// From https://stackoverflow.com/questions/28005865/golang-generate-unique-filename-with-extension
func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return filepath.Join(os.TempDir(), prefix+hex.EncodeToString(randBytes) + suffix)
}
