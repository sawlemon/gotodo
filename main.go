package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"slices"
	"strconv"
)

func main() {
	var tasks []string
    fmt.Println("Go TODO")
	for {
		var action, param string

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		
		if len(parts) < 1 {
			fmt.Println("No input")
			continue
		}
		
		action = parts[0]
		param = strings.Join(parts[1:], " ")
		
		fmt.Println("action ",action," param:",param)
		
		allowedActions := []string{"/bye", "/add", "/del", "ls"}
		
		if ! slices.Contains(allowedActions, action)  {
			fmt.Println("Unrecognized Action")
		}
		
		if action == "/bye" {
			fmt.Println("Bye Bye..")
			break
		}

		if action == "/add" {
			if len(param) < 1 {
				fmt.Println("No Parameter")
				continue
			}
			fmt.Println("adding task...")
			tasks = append(tasks, param)
		}
		
		if action == "/ls" {
			for i, v := range tasks {
				fmt.Println(i+1," : ",v)
			}	
		}

		if action == "/del"{
			// TODO:
			// max index must not go out of range
			// if the input is 0 it should be handled
			// handle the err when converting non into to into
			taskNumber, _ := strconv.Atoi(param)
			fmt.Println("taskNumeber:", taskNumber)
			tasks = slices.Delete(tasks	,taskNumber-1, taskNumber)
		}

		for i, v := range tasks {
			fmt.Println(i+1," : ",v)
		}
	}
}