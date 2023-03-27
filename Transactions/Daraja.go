package transactions

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getPassword() string {
	passphrase := os.Getenv("PASSPRHASE")
	data := "174379" + passphrase + TimeStamp()
	password := base64.StdEncoding.EncodeToString([]byte(data))
	return password
}
func TimeStamp() string {
	return time.Now().UTC().Format("20060102150405")
}

func DarajaApi(c *fiber.Ctx) error {
	BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	TOKEN_API := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	ShortBusinessCode := "174379"
	Amount := ""
	PhoneNumber := ""
	AccountReference := "SMOCHADELI"
	CallBackURL := ""
	TransactionDesc := "test"

	Consumerkey := os.Getenv("CONSUMERKEY")
	consumerSecret := os.Getenv("CONSUMERSECRET")

	// creating the http request

	req, err := http.NewRequest("GET", TOKEN_API, nil)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error.....",
		})
	}
	req.Close = true
	req.SetBasicAuth(Consumerkey, consumerSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error",
		})
	}
	defer resp.Body.Close()

	// read the response body

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}

	// Extract the access token from the response
	accessToken, ok := response["access_token"].(string)
	if !ok {
		panic("Access token not found in response")
	}

	// Print the access token
	fmt.Println("Access token:", accessToken)

	//writing the post request
	RequestBody := fmt.Sprintf(`{
		"ShortBusinessCode": "%s",
		"Amount":"%s",
		"PhoneNumber" :"%s",
		"AccountReference":"%s",
		"CallBackURL":"%s",
		"TransactionDesc":"%s",
		"Timestamp":"%s",
		"Password":"%s"
	
	}`, ShortBusinessCode, Amount, PhoneNumber, AccountReference, CallBackURL, TransactionDesc, TimeStamp(), getPassword())

	//TO-DO send a post request bearing the token,requestbody and

	req2, err := http.NewRequest("POST", BASE_API, bytes.NewBufferString(RequestBody))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error.....",
		})
	}
	req2.Close=true

	req2.Header.Set("Authorization", "Bearer"+accessToken)
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Add("Accept", "application/json")
	req2.Header.Add("Cache-Control", "no-cache")
	// sending the http req
	resp2,err := client.Do(req2)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error.....",
		})
		defer resp2.Body.Close()
		body,err := ioutil.ReadAll(resp2)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server Error.....",
			})
			fmt.Println(string(body))

	return c.Status(200).JSON(fiber.Map{
		"message": "transaction was succesful......",
	})
}
