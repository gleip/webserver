package webserver

import "fmt"

// Hi Возвращает приветственное сообение
func Hi(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
