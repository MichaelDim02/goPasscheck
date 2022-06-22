package main 

import "fmt"

func length_test(pass string) {

	length := len(pass);
	fmt.Printf("Length: %d\n", length);

	if (length <= 7) { 
		fmt.Printf(RED + "[!] Your password is very short! " + RESET);
		fmt.Printf("It is vulnerable to a pure bruteforce attack\n");
	} else if (length <= 10) {
		fmt.Printf(RED + "[!] Your password is short. " + RESET);
		fmt.Printf("It is vulnerable to a pure bruteforce attack\n");
	} else if (length <= 20) {
		fmt.Printf(GREEN + "[+] Your password is pretty long.\n" + RESET);
	} else if (length <= 30) {
		fmt.Printf(GREEN +"[+] Your password is very long!\n" + RESET);
	}
}
