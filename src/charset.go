package main

import (
	"fmt"
	"unicode"
)

func charset_test(pass string) {

	let := false
	upp := false
	low := false
	num := false
	sym := false

	for _, r := range pass {

		if (unicode.IsDigit(rune(r))) {
			num = true
		}

		if (unicode.IsLetter(rune(r))) {
			let = true

			if (unicode.IsUpper(r)) {
				upp = true
			} else {
				low = true
			}
		}

		if (unicode.IsSymbol(rune(r))) {
			sym = true
		}
	}

	if (num == true) {
		fmt.Printf(GREEN + "[+] Your password includes numbers\n" + RESET);
	} else {
		fmt.Printf(RED + "[!] Your password doesn't include numbers\n" + RESET);
	}

	if (sym == true) {
		fmt.Printf(GREEN + "[+] Your password includes special symbols\n" + RESET);
	} else {
		fmt.Printf(RED + "[!] Your password doesn't include special symbols\n" + RESET);
	}

	if (let == true) {
		fmt.Printf(GREEN + "[+] Your password includes letters\n" + RESET);

		if (upp == true) {
			fmt.Printf(GREEN + "	[+] includes upper case letters\n" + RESET);
		} else {
			fmt.Printf(RED + "	[!] doesn't include upper case letters\n" + RESET);
		}

		if (low == true) {
			fmt.Printf(GREEN + "	[+] includes lower case letters\n" + RESET);
		} else {
			fmt.Printf(RED + "	[!] doesn't include lower case letters\n" + RESET);
		}

	} else {
		fmt.Printf(RED + "[!] Your password doesn't include letters\n" + RESET);
	}

}
