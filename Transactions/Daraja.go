package transactions

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getPassword() string {
	//passphrase generated from smochadeliveryapp
	passphrase := "ConquestpDefendvHarvestrHotp9"
	envCode := os.Getenv("ENVCODE")
	data := passphrase + envCode
	hash := sha256.Sum256([]byte(data))
	password := hex.EncodeToString(hash[:])
	return password
}
func TimeStamp() string {
	return time.Now().UTC().Format("20060102150405")
}

func DarajaApi(c *fiber.Ctx) error {
	BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	TOKEN_API := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

	//ACCESS_TOKEN := os.Getenv("")

	// lipa na mpesa parameters

	ShortBusinessCode := "174379"
	Amount := ""
	PhoneNumber := ""
	AccountReference := "SMOCHADELI"
	CallBackURL := ""
	TransactionDesc := "test"

	// encoding of consumer key and customer secret

	Consumerkey := os.Getenv("CONSUMERKEY")
	consumerSecret := os.Getenv("CONSUMERSECRET")
	Auth := consumerSecret + ":" + Consumerkey

	//req body json

	// creating the http request

	req, err := http.NewRequest("GET", TOKEN_API, nil)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error.....",
		})
	}

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

	//authenticaton encoding
	AuthEncode := base64.StdEncoding.EncodeToString([]byte(Auth))

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

	//***setting up the request headers***

	//set-up the request headers
	//req.Header.Add("Authorization", "Basic"+AuthEncode)
	//req.Header.Add("Authorization", "Bearer"+)
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Api_", "application/json")
	// req.Header.Add("Accept", "application/json")
	// req.Header.Add("Cache-Control", "no-cache")
	// sending the http req

	return c.Status(200).JSON(fiber.Map{
		"respbody": body,
	})
}
