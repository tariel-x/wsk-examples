package main

// Main function for the action
func Main(obj map[string]interface{}) map[string]interface{} {
	name := obj["name"].(string)
	msg := map[string]interface{}{
		"msg": handle(name),
	}
	return msg
}
