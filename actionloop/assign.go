package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	// debugging
	var debug = os.Getenv("OW_DEBUG") != ""

	if debug {
		filename := os.Getenv("OW_DEBUG")
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err == nil {
			log.SetOutput(f)
			defer f.Close()
		}
		log.Printf("ACTION ENV: %v", os.Environ())
	}

	// assign the main function
	type Action func(event map[string]interface{}) map[string]interface{}
	var action Action
	action = Main

	// input
	out := os.NewFile(3, "pipe")
	defer out.Close()
	reader := bufio.NewReader(os.Stdin)

	// read-eval-print loop
	if debug {
		log.Println("started")
	}
	for {
		// read one line
		inbuf, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}
		if debug {
			log.Printf(">>>'%s'>>>", inbuf)
		}
		// parse one line
		var input map[string]interface{}
		err = json.Unmarshal(inbuf, &input)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(out, "{ error: %q}\n", err.Error())
			continue
		}
		if debug {
			log.Printf("%v\n", input)
		}
		// set environment variables
		err = json.Unmarshal(inbuf, &input)
		for k, v := range input {
			if k == "value" {
				continue
			}
			if s, ok := v.(string); ok {
				os.Setenv("__OW_"+strings.ToUpper(k), s)
			}
		}
		// get payload if not empty
		var payload map[string]interface{}
		if value, ok := input["value"].(map[string]interface{}); ok {
			payload = value
		}
		// process the request
		result := action(payload)
		// encode the answer
		output, err := json.Marshal(&result)
		if err != nil {
			log.Println(err.Error())
			fmt.Fprintf(out, "{ error: %q}\n", err.Error())
			continue
		}
		output = bytes.Replace(output, []byte("\n"), []byte(""), -1)
		if debug {
			log.Printf("'<<<%s'<<<", output)
		}
		fmt.Fprintf(out, "%s\n", output)
	}
}

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
	log.Printf("language %s", language)

	candidates := make([]string, 0, len(humans))
	if !ok {
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
	log.Printf("Candidates num %d", len(candidates))

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
