package save

import (
	"Calculator-Go/src/connection"
	"Calculator-Go/src/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"
)

func SaveHistory(example string, result float64) error {
	apiCLient := connection.NewApiCLient("http://localhost:8080")

	postBody := &types.PostBody{
		Calculation: fmt.Sprint(example, " = ", result),
		CreatedAt:   fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")),
	}
	postBodyJson, err := json.Marshal(postBody)
	if err != nil {
		log.Fatalf("error convert struct to json. error: %v", err)
	}

	response, statusCode, err := apiCLient.Post("/save", postBodyJson)
	if err != nil {
		log.Fatalf("error: %v", err)
		return fmt.Errorf("error: %v", err.Error())
	}

	var record *types.ResponsStruct
	err = json.Unmarshal(response, &record)
	if err != nil {
		return fmt.Errorf("invalid request format: %v", err)
	}

	if statusCode == 201 {
		return nil
	} else {
		log.Fatalf("error creating record: %v", statusCode)
		return fmt.Errorf("error save")
	}

}

func GetHistory(key string) ([]*types.ResponseHistory, error) {
	apiCLient := connection.NewApiCLient("http://localhost:8080")

	var requestURL string = "/history"
	if key != "" {
		params := url.Values{}
		params.Add("id", key)

		encodedParams := params.Encode()

		requestURL = requestURL + "?" + encodedParams
	}

	response, err := apiCLient.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	var data []*types.ResponseHistory
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding: %v", err)
	}

	return data, nil
}

// Save .txt
func writeToFile(text string) {
	fileName := "history.txt"
	// Get current date and time
	now := time.Now()

	data := now.Format("2006-01-02 15:04:05")

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	_, err = file.WriteString(data + " | " + text)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Done.")
}

func readingFromFile(index uint8) {
	const lengthSlice int8 = 64

	file, err := os.Open("history.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			data := make([]byte, lengthSlice)
			var i uint8
			for i = 0; i <= index; i++ {
				n, err := file.Read(data)
				if err == io.EOF {
					break
				}
				fmt.Print(string(data[:n]))
			}
		} else {
			fmt.Println("Error reading file")
		}
	}(file)
}

func hasNonZeroDecimalPart(num float64) bool {
	return num != float64(int(num))
}
