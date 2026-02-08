package main

import (
	"encoding/json"
	"log"
	"os"
)

func writeResponseBody(data []byte) {
	prettyJSON, err := json.MarshalIndent(json.RawMessage(data), "", "   ")

	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile("response_body.json", prettyJSON, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Response written to response_body.json file")
}
