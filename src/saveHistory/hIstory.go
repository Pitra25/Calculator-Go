package save

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type ApiCLient struct {
	BaseURL string
	Client  *http.Client
}

func NewApiCLient(baseUrl string) *ApiCLient {
	return &ApiCLient{
		BaseURL: baseUrl,
		Client:  &http.Client{},
	}
}

type PostBody struct {
	Calculation string `json:"calculation"`
	CreatedAt   string `json:"createdAt"`
}

type responsStruct struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SaveHistory(example string, result float64) error {
	apiCLient := NewApiCLient("http://localhost:8080")

	postBody := PostBody{
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

	var record responsStruct
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

	// if hasNonZeroDecimalPart(result) {
	// 	writeToFile(fmt.Sprintf("%s = %.2f\n", example, result))
	// } else {
	// 	writeToFile(fmt.Sprintf("%s = %d\n", example, int(result)))
	// }
}

type ResponseHistory struct {
	ID          int    `json:"id"`
	CreatedAt   string `json:"createdAt"`
	Calculation string `json:"calculation"`
}

func GetHistory(key string) ([]ResponseHistory, error) {
	apiCLient := NewApiCLient("http://localhost:8080")

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

	var data []ResponseHistory
	err = json.Unmarshal(response, &data)
	if err != nil {
		return nil, fmt.Errorf("error decoding: %v", err)
	}

	// reader := bytes.NewReader(response)
	// err = binary.Read(reader, binary.LittleEndian, &data)
	// if err != nil {
	// 	return nil, fmt.Errorf("error decoding: %v", err)
	// }

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
