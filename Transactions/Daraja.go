package transactions

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jwambugu/mpesa-golang-sdk"
)

func DarajaApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mpesaApp := mpesa.NewApp(http.DefaultClient, os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRET"), mpesa.Sandbox)

	stkPushRes, err := mpesaApp.STKPush(ctx, os.Getenv("PASSKEY"), mpesa.STKPushRequest{
		BusinessShortCode: 174379,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            10,
		PartyA:            254716881438,
		PartyB:            174379,
		PhoneNumber:       254729664004,
		CallBackURL:       "https://example.com",
		AccountReference:  "SmochaDeli",
		TransactionDesc:   "Test Request",
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v", stkPushRes)

	return c.Status(fiber.StatusOK).JSON("sent....")

}
