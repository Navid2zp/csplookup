# csplookup
Client package for CSP Lookup API.


### Install
```
go get github.com/Navid2zp/csplookup
```

### Usage

```go
import (
	"github.com/Navid2zp/csplookup"
)

func main() {

    client := csplookup.NewClient("YOUR_API_KEY")
    response, err = client.Lookup("4.2.2.4")
    
    if err != nil {
        panic(err)
    }
    // Note: it will return error only for connection errors
    // Note: use GetAPIError() to check api errors
    apiError := response.GetAPIError()
    if apiError != nil {
        fmt.Println("error on api:", apiError)
    }
    fmt.Println("lookup result:", response.Result)

}
```

Look up with time limit:

```go
// aborts if lookup took more than 10ms
response, err := client.LookupWithTL("4.2.2.4", time.Millisecond * 10)

if err != nil {
    if err != csplookup.TimeLimitReached {
        panic(err)
    } else {
        fmt.Println("it took too long!")
    }
}
apiError := response.GetAPIError()
if apiError != nil {
    fmt.Println("error on api:", apiError)
}
fmt.Println("lookup result:", response.Result)
```

NOTE: You might get charged even if your request gets aborted.

#### Methods

Some helper methods that can help you get the data you need faster:
```go
// returns country iso code
countryCode := response.GetCountryCode()

// returns country name in english
countryEnglishName := response.GetENCountryName()

// returns city name in english
cityEnglishName := response.GetENCityName()

// returns city timezone
cityTimeZone := response.GetTimeZone()

// returns time in location timezone
// error if timezone is not available in result
timeThere, err := response.GetTimeInTimeZone()

// returns location for timezone
// location will be a standard time.Location type
// error if timezone is not available in result
timezoneLocation, err := response.GetTimeZoneLocation()
```

License
----

MIT
