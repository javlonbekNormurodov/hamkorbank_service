package http

// Status ...
type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded",
	}

	BadEnvironment = Status{
		Code:        400,
		Status:      "BAD_ENVIRONMENT",
		Description: "The service has an invalid environment value",
	}

	NotFound = Status{
		Code:        404,
		Status:      "Not found",
		Description: "Phone number not found",
	}

	InternalServerError = Status{
		Code:        500,
		Status:      "Server error",
		Description: "Server error",
	}
)

// Can be added as many as need like belove examples
// 400	BAD_CONTINUATION_TOKEN	Invalid continuation token passed.
// 400	BAD_PAGE	Page number does not exist or is an invalid format (e.g. negative).
// 400	BAD_REQUEST	The resource you’re creating already exists.
// 400	INVALID_ARGUMENT	Invalid argument value passed.
// 400	INVALID_AUTH	Authentication/OAuth token is invalid.
// 400	INVALID_AUTH_HEADER	Authentication header is invalid.
// 400	INVALID_BATCH	Batched request is missing or invalid.
// 400	INVALID_BODY	A request body that was not in JSON format was passed.
// 400	UNSUPPORTED_OPERATION	Requested operation not supported.
// 401	ACCESS_DENIED	Authentication unsuccessful.
// 401	NO_AUTH	Authentication not provided.
// 403	NOT_AUTHORIZED	User has not been authorized to perform that action.
// 404	NOT_FOUND	Invalid URL.
// 405	METHOD_NOT_ALLOWED	Method is not allowed for this endpoint.
// 409	REQUEST_CONFLICT	Requested operation resulted in conflict.
// 429	HIT_RATE_LIMIT	Hourly rate limit has been reached for this token. Default rate limits are 2,000 calls per hour.
// 500	EXPANSION_FAILED	Unhandled error occurred during expansion; the request is likely to succeed if you don’t ask for expansions, but contact Eventbrite support if this problem persists.
// 500	INTERNAL_ERROR	Unhandled error occurred in Eventbrite. contact Eventbrite support if this problem persists.
