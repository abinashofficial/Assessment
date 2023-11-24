package services

import (
	"Assessment/model"
	"Assessment/store"
	"Assessment/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func New(store store.Store) Service {
	return &formServ{
		formsRepo: store.FormsRepo,
	}
}

type formServ struct {
	formsRepo store.Repository
}

func (s *formServ) Create(req map[string]string) (model.ConvertedRequest, error) {
	//create the channel for request
	model.RequestChannel <- req

	form, err := s.worker(model.RequestChannel)
	if err != nil {
		fmt.Println("Error converting numeric part to integer:", err)
		return model.ConvertedRequest{}, err
	}

	return form, nil
}

func (s *formServ) worker(req chan map[string]string) (model.ConvertedRequest, error) {

	// Receive requests from the channel
	receivedData := <-req

	converted, err := convertRequest(receivedData)
	if err != nil {
		fmt.Println("Error converting numeric part to integer:", err)
		return model.ConvertedRequest{}, err
	}

	//Send the converted request to the webhook
	go func() {
		if err := sendToWebhook(converted); err != nil {
			fmt.Printf("Error sending request to webhook: %v\n", err)
		}
	}()
	// Process the request (in this example, just print it)
	fmt.Println("Received request from worker:", req)
	return converted, nil
}

func ExistForm(form map[string]string, max int) map[string]model.FormValues {
	var attributesTraits = map[string]model.FormValues{}
	for i := 1; i <= max; i++ {
		var formKey, formType, formValue string
		for key := range form {
			numericPart, nonNumericPart := utils.SplitNumericAndNonNumeric(key)

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
			formValueType := model.FormValues{FormValue: form[formValue], FormType: form[formType]}
			attributesTraits[form[formKey]] = formValueType
		}

	}
	return attributesTraits
}

func convertRequest(original map[string]string) (model.ConvertedRequest, error) {
	var converted = model.ConvertedRequest{}

	var form1 = make(map[string]string)
	var form2 = make(map[string]string)
	var maximumValue int
	if original != nil {
		for key, v := range original {

			numericPart, nonNumericPart := utils.SplitNumericAndNonNumeric(key)

			if numericPart != "" {
				numericValue, err := strconv.Atoi(numericPart)
				if err != nil {
					fmt.Println("Error converting numeric part to integer:", err)
					return model.ConvertedRequest{}, err
				}
				if numericValue > maximumValue {
					maximumValue = numericValue
				}
			}

			if utils.CheckKeyInSlice(model.BasicForm1, nonNumericPart) {
				form1[key] = v
			}
			if utils.CheckKeyInSlice(model.BasicForm2, nonNumericPart) {
				form2[key] = v
			}
			if utils.CheckKeyInSlice(model.BasicForm3, nonNumericPart) {
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

	} else {
		fmt.Println("Pointer to map is nil")
	}

	return converted, nil
}

func sendToWebhook(converted model.ConvertedRequest) error {
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
