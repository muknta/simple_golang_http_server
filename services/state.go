package services

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"simple_http_server/models"
)

var (
	previousID  int
	previousRes string
)

func init() {
	previousID, previousRes = loadState()
}

func HandleRequest(req models.Request) models.Response {
	var res string
	if req.ID == previousID {
		res = previousRes
	} else {
		res = randomResult()
		previousID = req.ID
		previousRes = res

		// Save the state to a file
		saveState(previousID, previousRes)
	}

	return models.Response{
		Cmd:    req.Cmd,
		ID:     req.ID,
		Result: res,
	}
}

func randomResult() string {
	results := []string{"red", "blue", "green", "yellow", "orange", "purple", "black"}
	return results[rand.Intn(len(results))]
}

func saveState(id int, result string) {
	file, err := os.Create("state.txt")
	if err != nil {
		log.Println("Error saving state:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d\n%s", id, result))
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

func loadState() (int, string) {
	file, err := os.Open("state.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return 0, ""
		}
		log.Println("Error loading state:", err)
		return 0, ""
	}
	defer file.Close()

	var id int
	var result string
	fmt.Fscanf(file, "%d\n%s", &id, &result)

	return id, result
}
