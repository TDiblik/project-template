package utils

type IConstants struct {
	TokenHeaderName      string
	TokenClaimsLocalsKey string
}

var Constants IConstants = IConstants{
	TokenHeaderName:      "x-user-token",
	TokenClaimsLocalsKey: "user_info",
}

const NANOSECONDS_MULTIPLIER = int64(1000000000)
