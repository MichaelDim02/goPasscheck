package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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
	i := 1
	for fScanner.Scan() {
		line := fScanner.Text()
		if (len(line) >= 3 && pass == line) {

			fmt.Printf(RED + "\n[!!!] Password is in cracking dictionary - ")

			fmt.Printf(RESET + "%s:%d : %s\n\n", file,i, line)
			found = true
		} else if (len(line) >= 3 && strings.Contains(pass1, line)) {
			perc := get_per(len(line), len(pass1))
			if (perc == 100.0) {
				fmt.Printf(RED + "[!] Deleetified password is in cracking dictionary - %.1f%% - ", perc)
			} else if (perc >= 70.0) {
				fmt.Printf(RED + "[!] Deleetified password contains cracking dictionary password - %.1f%% - ", perc)
			} else if (perc >= 35.0) {
				fmt.Printf(YEL + "[-] Deleetified password contains cracking dictionary password - %.1f%% - ", perc)
			} else {
				fmt.Printf(BLUE + "[~] Deleetified password contains cracking dictionary password - %.1f%% - ", perc)
			}
			fmt.Printf(RESET + "%s : %s\n", file, line)
			found = true
		}
		i++
	}

	if (!found) {
		fmt.Printf(GREEN + "[+] Password is not in cracking dictionary " + RESET +"%s\n", file)
	}
	
	f.Close()
}
