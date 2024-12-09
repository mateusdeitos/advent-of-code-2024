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

//go:embed daysix.txt
var FileInputDaySix []byte

//go:embed daysix_example.txt
var FileInputDaySixExample []byte

//go:embed dayseven.txt
var FileInputDaySeven []byte

//go:embed dayseven_example.txt
var FileInputDaySevenExample []byte
