package embed

import _ "embed"

//go:embed dayone.txt
var FileInputDayone []byte
//go:embed dayone_example.txt
var FileInputDayoneExample []byte

//go:embed daytwo.txt
var FileInputDaytwo []byte
//go:embed daytwo_example.txt
var FileInputDaytwoExample []byte

//go:embed daythree.txt
var FileInputDaythree []byte
//go:embed daythree_example.txt
var FileInputDaythreeExample []byte
