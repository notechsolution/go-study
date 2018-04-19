package main

import (
	"os/exec"
	"fmt"
)

func main() {

	// For Windows CMD, no need to call command.Run(), just retrieve output from command.Output()
	fmt.Println("Execute external commands");
	cmd := exec.Command("cmd", "/C", "dir");
	content,_ := cmd.Output();
	fmt.Printf("output is %s\n", content)

	// for External program, need to call Run() to start them
	notepad :=exec.Command("notepad");
	error :=notepad.Run();
	if error !=nil {
		fmt.Printf("executed command with error %s\n", error);
	}
}
