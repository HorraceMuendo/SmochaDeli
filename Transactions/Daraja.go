package transactions

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jwambugu/mpesa-golang-sdk"
)

func DarajaApi(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mpesaApp := mpesa.NewApp(http.DefaultClient, "4T1XbAACROPyDG7np8nBLw2ALRQVsGkL", "qcbSipJSjFvV83Sj", mpesa.Sandbox)

	stkPushRes, err := mpesaApp.STKPush(ctx, "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919", mpesa.STKPushRequest{
		BusinessShortCode: 174379,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            10,
		PartyA:            254729664004,
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

	c.Status(fiber.StatusOK).JSON("sent....")

}
