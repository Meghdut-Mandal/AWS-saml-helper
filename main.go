package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "=")
		if len(arr) >= 2 {
			println(arr[0], " ", arr[1])
			key := strings.ToUpper(strings.TrimSpace(arr[0]))
			value := strings.TrimSpace(arr[1])
			println("Setting ", key, " as ", value)

			cmdstr := "-c \"export " + key + "=" + value + "\""
			cmd := exec.Command("bash", cmdstr)
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	os.Remove(file.Name())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
