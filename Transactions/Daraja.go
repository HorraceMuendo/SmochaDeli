package transactions

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func DarajaApi(c *fiber.Ctx) error {
	BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	ACCESS_TOKEN := os.Getenv("")

	// lipa na mpesa parameters
	ShortBusinessCode := ""
	Amount := ""
	PhoneNumber := "254729664004"
	AccountReference := "SMOCHADELIVERY"
	CallBackURL := ""
	TransactionDesc := "test"
	// encoding of consumer key and customer secret
	Consumerkey := os.Getenv("")
	consumerSecret := os.Getenv("")
	Auth := consumerSecret + ":" + Consumerkey
	AuthEncode := base64.StdEncoding.EncodeToString([]byte(Auth))
	//req body json

	RequestBody := fmt.Sprintf(`{
	"ShortBusinessCode": "%s",
	"Amount":"%s",
	"PhoneNumber" :"%s",
	"AccountReference":"%s",
	"CallBackURL":"%s",
	"TransactionDesc":"%s",
	"Tmestamp":"%s",
	"Password":"%s"


}`, ShortBusinessCode, Amount, PhoneNumber, AccountReference, CallBackURL, TransactionDesc, TimeStamp(), getPassword())

	// creating the http request
	req, err := http.NewRequest("POST", BASE_API, bytes.NewBufferString(RequestBody))
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error",
		})
	}

	//set-up the request headers
	req.Header.Set("Authorization", "Basic"+AuthEncode)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	// making the http req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error",
		})
	}
	// read the response body
	defer resp.Body.Close()
	responsBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server Error",
		})
	}
	fmt.Println(string(responsBody))

	return c.SendStatus(200)
}
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
