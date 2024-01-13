package cerror

import "errors"

var ErrAlreadyExist = errors.New("id already exist")
var ErrDecodeFile = errors.New("decoding file db")
var ErrStringToInt = errors.New("id from string to int")
var ErrWriteByte = errors.New("write byte")
var ErrInMemoryRepo = errors.New("initializing in memory repo")
var ErrOpenFileRepo = errors.New("opening file repo")
var ErrRunningServer = errors.New("running server")
var ErrEnvParseConfig = errors.New("parsing env")
var ErrIsDeleted = errors.New("is deleted")
