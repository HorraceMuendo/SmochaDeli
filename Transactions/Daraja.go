package transactions

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)
func DarajaApi(c* fiber.Ctx) error {
	BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
ACCESS_TOKEN:= os.Getenv("")

//lipa na mpesa parameters
ShortBusinessCode:=""
Amount:=""
PhoneNumber :=""
AccountReference:="SMOCHADELIVERY"
CallBackURL:=""
TransactionDesc:="test"
// encoding of consumer key and customer secret
Consumerkey:= os.Getenv("")
consumerSecret := os.Getenv("")
Auth:=consumerSecret+Consumerkey

}





//paybill/Buy_goods
ShortBusinessCode:=""
Password:=""
//timestamp of when the transcation took place
Timestamp:=""
//For paybill
TranscationType:="CustomerPayBillOnline"
Amount:=""
//phone no sending money
PartyA:=""
//Org recieving
PartyB:=""
//phone no to recieve stk
PhoneNumber :=""
CallBackURL:=""
//refrence defined for the acknowledgement of the transaction by the customer
AccountReference:="SMOCHADELIVERY"
//additional info sent by the system
TransactionDesc:="test"

