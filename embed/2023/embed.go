package embed

import _ "embed"

//go:embed dayone.txt
var FileInputDayone []byte
//go:embed dayone_example.txt
var FileInputDayoneExample []byte
