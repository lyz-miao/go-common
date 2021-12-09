package encode

// People Code start form 800 to 999
var (
    // AccountInfoError 账号密码错误
    FileOpenError        = Code{Code: 800, Message: "File open error"}
    FileReadError        = Code{Code: 801, Message: "File read error"}
    UploadFileIncomplete = Code{Code: 802, Message: "The upload file is incomplete"}
    UnsupportedMediaType = Code{Code: 803, Message: "TUnsupported Media Type"}
    PhotoAlreadyExists   = Code{Code: 804, Message: "The photo already exists"}
)
