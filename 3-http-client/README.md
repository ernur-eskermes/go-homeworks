# HTTP Client

## Example

```go
package main

import (
	"fmt"
	"github.com/ernur-eskermes/go-homeworks/3-http-client"
	"log"
	"os"
)

func main() {
	c := cloudpayments.NewClient(os.Getenv("CP_PUBLIC_ID"), os.Getenv("CP_API_SECRET"))

	if err := c.Test(""); err != nil {
		log.Fatal(err)
	}

	transaction, secure3d, err := c.ChargeCard(cloudpayments.ChargeCardInput{
		Amount:               5,
		CardCryptogramPacket: "",
		IpAddress:            "192.168.35.24",
		Currency:             cloudpayments.USD,
		RequireConfirmation:  false,
	}, "")
	if err != nil {
		log.Fatal(err)
	}
	if secure3d != nil {
		fmt.Printf("Secure3d: %+v\n", secure3d)
	}
	if transaction != nil {
		fmt.Printf("Transaction: %+v\n", transaction)
	}
}
```