package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func deleet(str1 string) string {
	str2 := strings.ReplaceAll(str1, "3", "e")
	str3 := strings.ReplaceAll(str2, "@", "a")
	str4 := strings.ReplaceAll(str3, "1", "i")
	str5 := strings.ReplaceAll(str4, "4", "a")
	str6 := strings.ReplaceAll(str5, "0", "o")
	str7 := strings.ReplaceAll(str6, "9", "g")
	str8 := strings.ReplaceAll(str7, "6", "b")

	str9 := strings.ToLower(str8)

	return str9
}

func words_test(pass string, file string) {

	pass1 := deleet(pass)

	f, err := os.Open(file)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fScanner := bufio.NewScanner(f)

	fScanner.Split(bufio.ScanLines)

	var found bool
	for fScanner.Scan() {
		line := fScanner.Text()
		if (len(line) >= 3 && strings.Contains(pass1, line)) {
			fmt.Printf(YEL + "[-] Password contains dictionary words - "+RESET)
			fmt.Printf("%s : %s\n", file, line)
			found = true
		}
	}

	if (found) {
		fmt.Printf("    The above is not necessarily an issue if your password\n")
		fmt.Printf("    contains multiple distinct dictionary words, or if the\n")
		fmt.Printf("    words are a small part of the password\n");
	} else {
		fmt.Printf(GREEN + "[+] Password was not found in any files\n"+RESET)
	}
	
	f.Close()
}