package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hoisie/mustache"
)

var dayToCursive = map[string]string{
	"1":  "dayone",
	"2":  "daytwo",
	"3":  "daythree",
	"4":  "dayfour",
	"5":  "dayfive",
	"6":  "daysix",
	"7":  "dayseven",
	"8":  "dayeight",
	"9":  "daynine",
	"10": "dayten",
	"11": "dayeleven",
	"12": "daytwelve",
	"13": "daythirteen",
	"14": "dayfourteen",
	"15": "dayfifteen",
	"16": "daysixteen",
	"17": "dayseventeen",
	"18": "dayeighteen",
	"19": "daynineteen",
	"20": "daytwenty",
	"21": "daytwentyone",
	"22": "daytwentytwo",
	"23": "daytwentythree",
	"24": "daytwentyfour",
	"25": "daytwentyfive",
	"26": "daytwentysix",
	"27": "daytwentyseven",
	"28": "daytwentyeight",
	"29": "daytwentynine",
	"30": "daythirty",
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: download <year> <day>")
	}

	year := os.Args[1]
	day := os.Args[2]

	cursive, ok := dayToCursive[day]
	if !ok {
		panic("Invalid day")
	}

	fmt.Printf("Downloading input for day %s in year %s\n", day, year)
	r, _ := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day+"/input", nil)

	session := os.Getenv("SESSION")
	if session == "" {
		fmt.Printf("Error: SESSION environment variable not set\n")
		os.Exit(1)
	}

	r.AddCookie(&http.Cookie{Name: "session", Value: session})

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Printf("Error downloading content err %v\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body %v\n", err)
		os.Exit(1)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error downloading content status code %d: %s\n", res.StatusCode, content)
		os.Exit(1)
	}

	// get pwd
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("CWD: %s\n", dir)

	// create year directory if it doesn't exist
	err = os.MkdirAll("./embed/"+year, 0755)
	if err != nil {
		fmt.Printf("Error creating directory %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("./embed/"+year+"/"+cursive+".txt", content, 0644)
	if err != nil {
		fmt.Printf("Error writing file %v\n", err)
		os.Exit(1)
	}

	dayCamelCase := strings.ToUpper(string(cursive[0])) + cursive[1:]
	args := map[string]string{
		"dayUpperCamelCase": dayCamelCase,
		"package":           cursive,
		"year":              year,
	}

	path := "handlers/" + year + "/" + cursive
	err = os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Printf("Error creating directory %v\n", err)
		os.Exit(1)
	}

	var output string
	output = mustache.RenderFile("download/handler.mustache", args)
	err = os.WriteFile(path+"/"+cursive+".go", []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file %v\n", err)
		os.Exit(1)
	}

	output = mustache.RenderFile("download/test_template.mustache", args)
	err = os.WriteFile(path+"/"+cursive+"_test.go", []byte(output), 0644)
	if err != nil {
		fmt.Printf("Error writing file %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat("embed/" + year + "/embed.go"); os.IsNotExist(err) {
		err = os.WriteFile("embed/"+year+"/embed.go", []byte("package embed\n\nimport _ \"embed\"\n"), 0644)
		if err != nil {
			fmt.Printf("Error writing file %v\n", err)
			os.Exit(1)
		}
	}

	embedContent, err := os.ReadFile("embed/" + year + "/embed.go")
	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
		os.Exit(1)
	}

	embedContent = append(embedContent, []byte("\n")...)
	embedContent = append(embedContent, []byte("//go:embed "+cursive+".txt\n")...)
	embedContent = append(embedContent, []byte("var FileInput"+dayCamelCase+" []byte\n")...)
	embedContent = append(embedContent, []byte("//go:embed "+cursive+"_example.txt\n")...)
	embedContent = append(embedContent, []byte("var FileInput"+dayCamelCase+"Example []byte\n")...)

	err = os.WriteFile("embed/"+year+"/embed.go", embedContent, 0644)
	if err != nil {
		fmt.Printf("Error writing file %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile("embed/"+year+"/"+cursive+"_example.txt", []byte(""), 0644)
	if err != nil {
		fmt.Printf("Error writing file %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downloaded input for day ./embed/%s.txt\n", cursive)

}
