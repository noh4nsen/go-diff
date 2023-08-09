package helper

import (
	"encoding/json"
	"go-diff/model"
	"log"
)

func EncodeToJson(changes model.Changes) string {
	response, err := json.Marshal(changes)
	CheckAndLogError("Error encoding to JSON: %v", err)
	prettyPrintOutput(changes)

	return string(response)
}

func prettyPrintOutput(changes model.Changes) {
	prettyJson, err := json.MarshalIndent(changes, "", "    ")
	CheckAndLogError("Error encoding to JSON: %v", err)
	log.Println("Output: ", (string(prettyJson)))
}
