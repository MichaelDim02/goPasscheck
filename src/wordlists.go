package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


func wordlist_test(pass string, file string) {

	pass1 := deleet(pass)

	f, err := os.Open(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fScanner := bufio.NewScanner(f)

	fScanner.Split(bufio.ScanLines)

	var found bool
	var i int
	for fScanner.Scan() {
		line := fScanner.Text()
		if (len(line) >= 3 && strings.Contains(pass1, line)) {

			fmt.Printf(RED + "\n[!!!] Password is cracking dictionary - ")

			fmt.Printf(RESET + "%s:%d : %s\n\n", file,i, line)
			found = true
			i++
		}
	}

	if (!found) {
		fmt.Printf(GREEN + "[+] Password is not in dictionary " + RESET +"%s\n", file)
	}
	
	f.Close()
}
