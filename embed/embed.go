package embed

import _ "embed"

//go:embed dayone.txt
var FileInputDayOne []byte

//go:embed daytwo.txt
var FileInputDayTwo []byte
