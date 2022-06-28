# The Maplerad Go client



The library follows an object-oriented approach and methods are grouped under categories.

There are currently ten (10) base categories namely:
 * Customer
 * Collections
 * Transfer
 * Bills
 * Wallets
 * Issuing
 * Counterparty
 * Forex
 * Institutions
 * Misc

 #### Learn more from the [docs](https://maplerad.dev/reference)

### Installation
```shell
    $ go get -u github.com/wirepay/maplerad-go
```

## Authorization
 A secret key is needed for authorization. It can be gotten from the Maplerad dashboard

## Environments
Maplerad provides two environments to ensure a smooth and easy experience.

* sandbox: for development
* live: for production

### Sandbox
Sandbox is your playground. You can credit your test wallets and use that to test your integrations, no real money will be debited or credited.
Ensure to switch to Live when you are ready to launch.

### Live
All method calls under Live will be charged and real money will be debited or credited.
You are advised to use this when you have fully tested your integrations and are ready to launch your product.


```go
package main

import (
	"fmt"
	maplerad "github.com/wirepay/maplerad-go/lib"
	"log"
	"os"
)

func main() {
	// maplerad.NewClient(secret_key, environment)
	client, err := maplerad.NewClient(os.Getenv("SECRET_KEY"), "sandbox")
	if err != nil {
		log.Fatalln(err)
	}
	// To get all wallets
	
	                // client.Base-class.Base-class-method
	wallets, err := client.Wallet.GetWallets()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(wallets)
}


```