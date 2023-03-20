package transactions
import(
	"fmt"
)


BASE_API := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

//variable declarations
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
AccountReference:=""
//additional info sent by the system
TransactionDesc:=""
