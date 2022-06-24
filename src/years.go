package main

import (
	"fmt"
	"regexp"
)

func years_test(pass string) {

	r, _ := regexp.Compile("((19)|(20))[0-9]{2}")
	match := r.MatchString(pass)

	if (match == true) {
		for _, i := range r.FindAllString(pass, -1) {
			perc := get_per(4, len(pass))

			if (perc == 100.0) {
				fmt.Printf(RED + "[!] Password is a recent year- %.1f%% - ", perc)
			} else if (perc >= 50.0) {
				fmt.Printf(RED + "[!] Password contains a recent year - %.1f%% - ", perc)
			} else if (perc >= 35.0) {
				fmt.Printf(YEL + "[-] Password contains a recent year - %.1f%% - ", perc)
			} else {
				fmt.Printf(BLUE + "[~] Password contains a recent year - %.1f%% - ", perc)
			}
			fmt.Printf(RESET + "%s\n", i)
		}
	} else {
		fmt.Printf(GREEN + "[~] Password does not contain a recent year\n" + RESET)
	}
}
