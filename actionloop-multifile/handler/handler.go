package handler

import (
	"fmt"
)

// Handle function for the action
func Handle(obj map[string]interface{}) map[string]interface{} {
	name := obj["name"].(string)
	msg := map[string]interface{}{
		"msg": make(name),
	}
	return msg
}

func make(name string) string {
	if name == "" {
		name = "stranger"
	}
	fmt.Printf("name=%s\n", name)
	return "This is the action with actionloop and subpackage, " + name + "!"
}
