package handler

import (
	"fmt"
	"github.com/rylans/getlang"
)

// Handle function for the action
func Handle(obj map[string]interface{}) map[string]interface{} {
	text := obj["text"].(string)
	msg := map[string]interface{}{
		"msg": detect(text),
	}
	return msg
}

func detect(text string) string {
	info := getlang.FromString(text)
	return fmt.Sprintf("Language of %q is %s", text, info.LanguageName())
}
