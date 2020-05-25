package csplookup

import (
	"fmt"
	"github.com/Navid2zp/easyreq"
	"github.com/Navid2zp/httptracer"
	"time"
)

const (
	BaseDomain     = "https://lookup.configserver.pro"
	LookupEndPoint = "https://lookup.configserver.pro/api/v1/lookup/ip?ip=%s"
)

// traces the connection and returns a httptracer.TracerResult type
// use httptracer.TracerResult.FullResponse to see the total time
func TestResponseTime() (*httptracer.TracerResult, error) {
	return httptracer.Tracer(BaseDomain, "GET")
}

// Creates and returns a new client for api key
func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

// Get default headers
func (c *Client) getHeaders() map[string]string {
	return map[string]string{"apiKey": c.APIKey}
}

// main function to lookup an ip
// returns error only if an error happens on connection
// use GetAPIError() to check for api errors
func (c *Client) Lookup(ip string) (*IPLookup, error) {
	var result IPLookup
	req := easyreq.Request{
		URL:              fmt.Sprintf(LookupEndPoint, ip),
		Headers:          c.getHeaders(),
		Method:           "GET",
		//ResponseDataType: "json",
		//SaveResponseTo:   &result,
	}
	res, err := req.Make()
	if err != nil {
		return &result, err
	}
	err = res.ToJson(&result)
	if err != nil {
		body, err := res.ReadBody()
		fmt.Println("error:", err)
		fmt.Println(string(body))
	}
	defer res.CloseBody()
	return &result, err
}

// returns lookup result if it took less than the provided timeLimit.
// returns TimeLimitReached error if it took more than provided timeLimit.
// NOTE: use TestResponseTime() to check for normal response time.
// NOTE: use this only if you can't effort any long lookup time.
// WARNING: do not use short time periods.
// WARNING: if requests reaches the server and then it gets aborted, you'll still be charged for request.
func (c *Client) LookupWithTL(ip string, timeLimit time.Duration) (*IPLookup, error) {
	var result *IPLookup
	var err error
	c1 := make(chan string, 1)
	go func() {
		result, err = c.Lookup(ip)
		c1 <- "done"
	}()
	select {
	case <-c1:
		return result, err
	case <-time.After(timeLimit):
		return nil, TimeLimitReached
	}
}

// checks for any error on api
// returns nil if there is no error and result is provided
// maps error code to internal errors making handling errors easier
func (l *IPLookup) GetAPIError() error {
	if l.ErrorCode == "" {
		return nil
	}
	switch l.ErrorCode {
	case "SERVER_ERROR":
		return ServerError
	case "EXPIRED_KEY":
		return ExpiredKey
	case "DAILY_LIMIT_REACHED":
		return DailyLimitReached
	case "MONTHLY_LIMIT_REACHED":
		return MonthlyLimitReached
	case "MAX_LIMIT_REACHED":
		return MaxLimitReached
	case "INVALID_KEY":
		return InvalidKey
	case "DEACTIVATED_KEY":
		return InactiveKey
	case "INVALID_IP":
		return InvalidIp
	case "NO_IP_PROVIDED":
		return NoIpProvided
	case "NO_API_KEY":
		return NoAPIKey
	default:
		return UnknownError
	}
}

// returns country code
func (l *IPLookup) GetCountryCode() string {
	return l.Result.Country.IsoCode
}

// search for english country name and return it
// use GetCountryName() for other languages
func (l *IPLookup) GetENCountryName() string {
	if val, ok := l.Result.Country.Names["en"]; ok {
		return val
	}
	return ""
}

// returns country name in provided language code if it exists
// returns empty string if it doesn't exist
func (l *IPLookup) GetCountryName(lang string) string {
	if val, ok := l.Result.Country.Names[lang]; ok {
		return val
	}
	return ""
}

// search for english city name and return it
// use GetCityName() for other languages
func (l *IPLookup) GetENCityName() string {
	if val, ok := l.Result.City.Names["en"]; ok {
		return val
	}
	return ""
}

// returns city name in provided language code if it exists
// returns empty string if it doesn't exist
func (l *IPLookup) GetCityName(lang string) string {
	if val, ok := l.Result.City.Names[lang]; ok {
		return val
	}
	return ""
}

// returns location timezone
func (l *IPLookup) GetTimeZone() string {
	return l.Result.Location.TimeZone
}

// returns location for timezone
// location will be a standard time.Location type
func (l *IPLookup) GetTimeZoneLocation() (*time.Location, error) {
	return time.LoadLocation(l.GetTimeZone())
}

// returns current time in the location
// default time.Time type if location doesn't exists in resutl
func (l *IPLookup) GetTimeInTimeZone() (time.Time, error) {
	loc, err := l.GetTimeZoneLocation()
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(loc), nil
}
