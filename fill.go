package main

import (
	"github.com/buger/goterm"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func e(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	//clear screen
	goterm.Clear()

	//disable input buffering
	er := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	e(er)

	//do not display characters on the screen
	er = exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	//Capture SIGTERM
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	x := 1
	y := 1

	go func() {
		<-sigs
		er = exec.Command("stty", "-F", "/dev/tty", "echo").Run()
		goterm.Clear()
		os.Exit(0)
		done <- true
	}()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		switch string(b[0]) {
		case "h":
			left(&y)
		case "j":
			down(&x)
		case "k":
			up(&x)
		case "l":
			right(&y)
		default:
			right(&y)
		}
		draw(x, y)
	}

	<-done
}

func draw(x int, y int) {
	goterm.MoveCursor(x, y)
	goterm.Print("#")
	goterm.Flush()
}

func left(y *int) {
	if 1 < *y && *y <= goterm.Width() {
		*y--
	}
}

func down(x *int) {
	if 1 <= *x && *x < goterm.Height() {
		*x++
	}
}

func up(x *int) {
	if 1 < *x && *x <= goterm.Height() {
		*x--
	}
}

func right(y *int) {
	if 1 <= *y && *y < goterm.Width() {
		*y++
	}
}
