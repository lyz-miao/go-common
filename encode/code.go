package encode

//  400 to 599
var (
	OK                  = Code{Code: 0, Message: ""}
	Unauthorized        = Code{Code: 401, Message: "Unauthorized"}
	Forbidden           = Code{Code: 403, Message: "Forbidden"}
	ParameterError      = Code{Code: 400, Message: "Parameter error"}
	InternalServerError = Code{Code: 401, Message: "Internal Server Error"}
)
