package main

import (
	"log"
	"math/rand"
)

var (
	humans = map[string][]string{
		"ivan":     {"go", "php"},
		"nikita":   {"go", "php"},
		"alexandr": {"go", "js"},
		"valentin": {"go", "js", "php"},
		"maxim":    {"js"},
		"alice":    {"java"},
	}
)

// Main selects assignee for MR
func Main(obj map[string]interface{}) map[string]interface{} {
	language, ok := obj["language"].(string)

	candidates := make([]string, 0, len(humans))

	if ok {
		for human := range humans {
			candidates = append(candidates, human)
		}
	} else {

		for human, langs := range humans {
			if in(language, langs) {
				candidates = append(candidates, human)
			}
		}
	}

	num := rand.Intn(len(candidates))

	log.Printf("New assegnee %s for language %s", candidates[num], language)

	msg := make(map[string]interface{})
	msg["assignee"] = candidates[num]
	return msg
}

func in(search string, scope []string) bool {
	for _, occurrence := range scope {
		if search == occurrence {
			return true
		}
	}
	return false
}
