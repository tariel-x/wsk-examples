package main

import (
	"log"
)

// Main selects assignee for MR
func Main(obj map[string]interface{}) map[string]interface{} {
	assignee, ok := obj["assignee"].(string)
	if !ok {
		log.Printf("No assignee")
		return nil
	}
	log.Printf("New assignee %s", assignee)

	msg := make(map[string]interface{})
	msg["msg"] = "Please check MR, " + assignee
	return msg
}
