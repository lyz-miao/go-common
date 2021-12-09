package encode

// User Code start form 500 to 699
var (
    // AccountInfoError 账号密码错误
    AccountInfoError    = Code{Code: 500, Message: "Username or password is error"}
    RemoveTokenError    = Code{Code: 501, Message: ""}
    // EncodingHashIDError 编码哈希ID错误
    EncodingHashIDError = Code{Code: 502, Message: "Encoding hash Id error"}
    // DecodingHashIDError 解码哈希ID错误
    DecodingHashIDError = Code{Code: 502, Message: "Decoding hash Id error"}
    // NotFindUser 找不到用户
    NotFindUser         = Code{Code: 503, Message: "Can not find user"}
)
