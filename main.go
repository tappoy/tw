package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const dirname = ".tw"
const basename = "tw.txt"

func twfile() string {
	return filepath.Join(os.Getenv("HOME"), dirname, basename)
}

func main() {
	var (
		a = flag.Bool("a", false, "Print all memos")
		t = flag.Bool("t", false, "Print today's memos")
		g = flag.Bool("g", false, "Search memos by keyword using regular expression")
	)
	flag.Parse()
	args := flag.Args()

	if *a {
		printAllMemos()
	} else if *t {
		printTodaysMemos()
	} else if *g {
		searchMemos(args)
	} else {
		if len(args) == 0 {
			flag.Usage()
		} else {
			appendMemo(args)
		}
	}
}

func appendMemo(args []string) {
	// create a directory if it does not exist
	os.Mkdir(filepath.Join(os.Getenv("HOME"), dirname), 0755)

	f, err := os.OpenFile(twfile(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04")
	host, _ := os.Hostname()
	user := os.Getenv("USER")
	content := strings.Join(args, " ")

	fmt.Fprintf(f, "%s %s@%s: %s\n", timestamp, user, host, content)
}

func printAllMemos() {
	memo, err := ioutil.ReadFile(twfile())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(memo))
}

func printTodaysMemos() {
	memo, err := ioutil.ReadFile(twfile())
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(memo), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, time.Now().Format("2006-01-02")) {
			fmt.Println(line)
		}
	}
}

func searchMemos(args []string) {
	memo, err := ioutil.ReadFile(twfile())
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(memo), "\n")
	rx := regexp.MustCompile(strings.Join(args, "|"))
	for _, line := range lines {
		if rx.MatchString(line) {
			fmt.Println(line)
		}
	}
}
