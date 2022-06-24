# The Maplerad Go client

### To install, run in terminal
```shell
    $ go get -u https://github.com/wirepay/maplerad-go
```

#
```go
    package main
    
    import (
    maplerad "github.com/wirepay/maplerad-go"
	"os"
    )
    
    func main(){
		client := maplerad.NewClient(os.Getenv("SECRET_KEY"), "sandbox")
		
		
    }
    

```