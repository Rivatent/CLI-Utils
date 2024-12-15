package xargs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Xargs struct {
	command string
	args    []string
}

func Init() (Xargs, error) {
	res := Xargs{}

	if len(os.Args) < 2 {
		return res, errors.New("command should be specified")
	}
	command := os.Args[1]
	commandArgs := os.Args[2:]
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			commandArgs = append(commandArgs, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return res, err
		}
	}
	res.command = command
	res.args = append(res.args, commandArgs...)
	return res, nil
}

func (x *Xargs) Exec(out io.Writer) {
	cmd := exec.Command(x.command, x.args...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func MyXargs(in io.Reader, out io.Writer, cliArgs []string) {

	command := cliArgs[1]
	commandArgs := cliArgs[2:]
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(in)
		for scanner.Scan() {
			commandArgs = append(commandArgs, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
	}

	cmd := exec.Command(command, commandArgs...)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
