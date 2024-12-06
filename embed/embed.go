package embed

import _ "embed"

//go:embed dayone.txt
var FileInputDayOne []byte

//go:embed daytwo.txt
var FileInputDayTwo []byte

//go:embed daythree.txt
var FileInputDayThree []byte

//go:embed dayfour.txt
var FileInputDayFour []byte

//go:embed dayfive.txt
var FileInputDayFive []byte

//go:embed dayfive_example.txt
var FileInputDayFiveExample []byte
