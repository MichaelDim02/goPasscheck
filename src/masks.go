package main

import (
	"fmt"
	"regexp"
)

func must_match(regex string, pass string) bool {
	match, _ := regexp.MatchString(regex, pass)
	return match
}


func mask_test(pass string) {

	str := " Password matches common mask "
	var found bool

	if (must_match("^[A-Za-z]{1,}$", pass)) {
		fmt.Printf(RED + "[!]" + str + "^alpha{1,}$ (only letters)\n" + RESET)
		found = true
	}

	if (must_match("^[0-9]{1,}$", pass)) {
		fmt.Printf(RED + "[!]" + str + "^num{1,}$ (only numbers)\n" + RESET)
		found = true
	}

	if (must_match("^[A-Za-z]{1,}[0-9]{1,}$", pass)) {
		fmt.Printf(YEL + "[-]" + str + "^alpha{1,}num{1,}$\n" + RESET)
		found = true
	}

	if (must_match("^[0-9]{1,}[A-Za-z]{1,}$", pass)) {
		fmt.Printf(YEL + "[-]" + str + "^num{1,}alpha{1,}$\n" + RESET)
		found = true
	}

	if (must_match("^[0-9]{1,}[^A-Za-z0-9]{1,}$", pass)) {
		fmt.Printf(YEL + "[-]" + str + "^num{1,}symbol{1,}$\n" + RESET)
		found = true
	}

	if (must_match("^[A-Za-z]{1,}[0-9]{1,}[^A-Za-z0-9]{1,}$", pass)) {
		fmt.Printf(YEL + "[-]" + str + "^alpha{1,}num{1,}symbol{1,}$\n" + RESET)
		found = true
	}

	if (must_match("[^A-Za-z0-9]{1,}$", pass)) {
		fmt.Printf(BLUE + "[~]" + str + "any{1,}symbol{1,}$\n" + RESET)
		found = true
	}

	if (must_match("^[A-Za-z]{1,}123$", pass)) {
		fmt.Printf(YEL + "[-]" + str + "^alpha{1,}123$\n" + RESET)
		found = true
	}

	if (must_match("123$", pass)) {
		fmt.Printf(BLUE + "[~]" + str + "*123$\n" + RESET)
		found = true
	}

	if (!found) {
		fmt.Printf(GREEN + "[+] Password does not match any common mask\n" + RESET)
	}
}
