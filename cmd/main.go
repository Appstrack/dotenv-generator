package main

import (
	"bufio"
	dotenv_generator "dotenv-generator"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	getenvPattern string
	fileName      string
)

//
func main() {
	flag.StringVar(&getenvPattern, "pattern", "os.Getenv", "pattern for searching getenv method")
	flag.StringVar(&fileName, "f", "env.go", "file name for searching pattern")
	flag.Parse()

	//log.Printf("searching for pattern: %s\n", getenvPattern)
	//log.Printf("working file: %s\n", fileName)

	if fileName == "" {
		log.Printf("file name empty, stop searching\n")
		os.Exit(1)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Printf("file open error: %s\n", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var (
		txt, temp string
		index     int
		parts     []string
		varList   = make(dotenv_generator.Environments, 0, 1)
	)
	for scanner.Scan() {
		txt = scanner.Text()

		if strings.Contains(txt, getenvPattern) {
			//log.Printf("line: %s\n", txt)
			index = strings.Index(txt, getenvPattern) + len(getenvPattern)
			temp = txt[index:]
			//log.Printf("index: %s\n", temp)
			temp = strings.Trim(temp, "()")
			//log.Printf("index 2: %s\n", temp)
			parts = strings.Split(temp, ",")
			//log.Printf("parts: %v\n", parts)
			if len(parts) != 2 {
				continue
			}
			for i, it := range parts {
				parts[i] = strings.Trim(it, " \"")
			}
			varList = append(varList, &dotenv_generator.Environment{
				Name:  parts[0],
				Value: parts[1],
			})
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = dotenv_generator.UpdateEnvironments(varList)
	if err != nil {
		log.Printf("save file error: %s\n", err)
		os.Exit(1)
	}
}
