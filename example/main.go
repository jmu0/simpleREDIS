package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmu0/simpleREDIS"
)

func main() {
	var host = "localhost:6379"
	r, err := simpleREDIS.NewRedis(host)
	printErr(err)
	args := os.Args[1:]
	if len(args) < 2 {
		help()
	}
	switch args[0] {
	case "get":
		var val string
		val, err = r.Get(args[1])
		printErr(err)
		fmt.Println(val)
	case "set":
		if len(args) == 3 {
			err = r.Set(args[1], args[2])
			printErr(err)
		} else {
			help()
		}
	case "del":
		var n int64
		n, err = r.Del(args[1])
		printErr(err)
		if n == 0 {
			printErr(errors.New("Deleted nothing"))
		}
	default:
		help()
	}
}

func help() {
	fmt.Println("get: redis get <key>")
	fmt.Println("set: redis set <key> <value>")
	fmt.Println("del: redis del <key>")
	os.Exit(0)
}

func printErr(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
