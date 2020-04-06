package csplookup

import "errors"

var NoAPIKey = errors.New("NO_API_KEY")
var ServerError = errors.New("SERVER_ERROR")
var ExpiredKey = errors.New("EXPIRED_KEY")
var DailyLimitReached = errors.New("DAILY_LIMIT_REACHED")
var MonthlyLimitReached = errors.New("MONTHLY_LIMIT_REACHED")
var MaxLimitReached = errors.New("MAX_LIMIT_REACHED")
var NoIpProvided = errors.New("NO_IP_PROVIDED")
var InvalidIp = errors.New("INVALID_IP")
var InvalidKey = errors.New("INVALID_KEY")
var InactiveKey = errors.New("DEACTIVATED_KEY")
var UnknownError = errors.New("UNKNOWN_ERROR")
var TimeLimitReached = errors.New("TIME_LIMIT_REACHED")
