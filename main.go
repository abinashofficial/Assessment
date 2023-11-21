package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func splitNumericAndNonNumeric(input string) (numericPart, nonNumericPart string) {
	var numericChars []rune
	var nonNumericChars []rune

	for _, char := range input {
		if char >= '0' && char <= '9' {
			numericChars = append(numericChars, char)
		} else {
			nonNumericChars = append(nonNumericChars, char)
		}
	}

	numericPart = string(numericChars)
	nonNumericPart = string(nonNumericChars)
	return numericPart, nonNumericPart
}

func CheckKeyInSlice(strArray []string, key string) bool {
	if strArray == nil {
		return false
	}
	for _, val := range strArray {
		if val == key {
			return true
		}
	}
	return false
}

func ExistForm(form map[string]string, max int) map[string]formValues {
	var attributesTraits = map[string]formValues{}
	for i := 1; i <= max; i++ {
		var formKey, formType, formValue string
		for key := range form {
			numericPart, nonNumericPart := splitNumericAndNonNumeric(key)

			if numericPart != "" {
				numericValue, err := strconv.Atoi(numericPart)
				if err != nil {
					fmt.Println("Error converting numeric part to integer:", err)
					return nil
				}

				if i == numericValue {
					if nonNumericPart == "atrk" || nonNumericPart == "uatrk" {
						formKey = key
					}
					if nonNumericPart == "atrv" || nonNumericPart == "uatrv" {
						formValue = key
					}
					if nonNumericPart == "atrt" || nonNumericPart == "uatrt" {
						formType = key
					}
				}
			}
		}

		if formKey != "" && formValue != "" && formType != "" {
			formValueType := formValues{FormValue: form[formValue], FormType: form[formType]}
			attributesTraits[form[formKey]] = formValueType
		}

	}
	return attributesTraits
}

func convertRequest(original map[string]string) ConvertedRequest {
	var converted = ConvertedRequest{}

	var form1 = make(map[string]string)
	var form2 = make(map[string]string)
	var maximumValue int

	for key, v := range original {

		numericPart, nonNumericPart := splitNumericAndNonNumeric(key)

		if numericPart != "" {
			numericValue, err := strconv.Atoi(numericPart)
			if err != nil {
				fmt.Println("Error converting numeric part to integer:", err)
				return ConvertedRequest{}
			}
			fmt.Printf("Numeric Part: %d\n", numericValue)
			fmt.Printf("Non-Numeric Part: %s\n", nonNumericPart)

			if numericValue > maximumValue {
				maximumValue = numericValue
			}
		}

		if CheckKeyInSlice(basicForm1, nonNumericPart) {
			form1[key] = v
		}
		if CheckKeyInSlice(basicForm2, nonNumericPart) {
			form2[key] = v
		}
		if CheckKeyInSlice(basicForm3, nonNumericPart) {
			if nonNumericPart == "ev" {
				converted.Event = v
			}
			if nonNumericPart == "et" {
				converted.EventType = v
			}
			if nonNumericPart == "id" {
				converted.AppID = v
			}
			if nonNumericPart == "uid" {
				converted.UserID = v
			}
			if nonNumericPart == "mid" {
				converted.MessageID = v
			}
			if nonNumericPart == "t" {
				converted.PageTitle = v
			}
			if nonNumericPart == "p" {
				converted.PageURL = v
			}
			if nonNumericPart == "l" {
				converted.BrowserLang = v
			}
			if nonNumericPart == "sc" {
				converted.ScreenSize = v
			}
		}
	}

	converted.UserTraits = ExistForm(form1, maximumValue)
	converted.Attributes = ExistForm(form2, maximumValue)

	return converted
}

func sendToWebhook(converted ConvertedRequest) error {
	webhookURL := "https://webhook.site/ee78f661-4e99-4273-9c3e-9a6293f5bb69"

	// Marshal the converted request to JSON
	jsonData, err := json.Marshal(converted)
	if err != nil {
		return err
	}

	//Send the JSON data to the webhook URL
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("webhook request failed with status code %d", resp.StatusCode)
	}

	fmt.Println("Request sent to webhook successfully")
	return nil
}

func processRequest(original map[string]string) {
	//Convert the original request
	converted := convertRequest(original)

	//Send the converted request to the webhook
	if err := sendToWebhook(converted); err != nil {
		fmt.Printf("Error sending request to webhook: %v\n", err)
	}
}

func main() {
	// Create a buffered channel for original requests
	originalRequests := make(chan map[string]string, 10) // Adjust the capacity based on your needs

	// Start a goroutine to process each request
	go func() {
		for req := range originalRequests {
			go processRequest(req)
		}
	}()

	// Start the HTTP server
	port := 8080
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Parse the JSON request
		var req map[string]string
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Send the request to the channel
		originalRequests <- req

		// Send a response to the client
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Request received successfully"))
		if err != nil {
			return
		}
	})

	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return
	}
}
