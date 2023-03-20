package transactions

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)
func DarajaApi(c* fiber.Ctx) error {
	BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
ACCESS_TOKEN:= os.Getenv("")

//lipa na mpesa parameters
ShortBusinessCode:=""
Amount:=""
PhoneNumber :="254729664004"
AccountReference:="SMOCHADELIVERY"
CallBackURL:=""
TransactionDesc:="test"
// encoding of consumer key and customer secret
Consumerkey:= os.Getenv("")
consumerSecret := os.Getenv("")
Auth:=consumerSecret+":"+Consumerkey
AuthEncode:=base64.StdEncoding.EncodeToString([]byte(Auth))
//req body json

RequestBody:=fmt.Sprintf(`{
	"ShortBusinessCode": "%s",
	"Amount":"%s",
	"PhoneNumber" :"%s",
	"AccountReference":"%s",
	"CallBackURL":"%s",
	"TransactionDesc":"%s",
	"Tmestamp":"%s",
	"Password":"%s"


}`,ShortBusinessCode,Amount,PhoneNumber,AccountReference,CallBackURL,TransactionDesc,TimeStamp(),getPassword())





}
func getPassword() string {
	//passphrase generated from smochadeliveryapp
	passphrase:="ConquestpDefendvHarvestrHotp9"
	envCode:= os.Getenv("ENVCODE")
	data := passphrase+envCode
	hash:= sha256.Sum256([]byte(data))
	password := hex.EncodeToString(hash[:])
	return password
}
func TimeStamp() string{
	return time.Now().UTC().Format("20060102150405")
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

