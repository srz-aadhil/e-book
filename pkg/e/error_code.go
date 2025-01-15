package e

//400 errors
const (
	//ErrInvalidRequest : when post body, query param, or path param is invalid
	ErrInvalidRequest = 400001 + iota

	//ErrValidateRequest : error when validating the request
	ErrValidateRequest

	//ErrDecodeRequest : error when decoding the request body
	ErrDecodeRequestBody
)

//404 Errors
const (
	//ErrResourceNotFound: When no records corresponding to the request is found in the DB
	ErrResourceNotFound = 404001
)

//500 errors
const (
	//ErrInternalServer : unexpected error
	ErrInternalServer = 500001
)