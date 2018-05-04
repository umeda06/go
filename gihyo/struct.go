package main

import (
	"fmt"
)

// 構造体
type Task struct {
	ID     int
	Detail string
	done   bool
}

// コンストラクタ
func NewTask(id int, detail string) *Task {
	task := &Task{
		ID:     id,
		Detail: detail,
		done:   false,
	}
	return task
}

// メソッド
func (task Task) String1() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func (task *Task) Finish() {
	task.done = true
}

func main() {
	task := NewTask(1, "buy the milk")
	// &{ID:1 Detail:buy the milk done:false}
	fmt.Printf("%s\n", task.String1())

	task.Finish()
	// &{ID:1 Detail:buy the milk done:true}
	fmt.Printf("%+v", task)
}
