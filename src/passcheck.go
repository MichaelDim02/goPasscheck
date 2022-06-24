package main 

import "fmt"
import "os"


var version string = "v0.1"

var RED string = "\x1b[31m"
var YEL string = "\u001b[33m"
var BLUE string = "\033[0;34m"
var GREEN string = "\033[0;32m"
var RESET string = "\x1b[0m"

func banner() {
	fmt.Printf("goPasstest %s\n" +
		   "'(O.O)' <(better safe than sorry)\n" +
		   "by Michael Constantine Dimopoulos\n" +
		   "https://mcdim.xyz <mk@mcdim.xyz>\n", version)
}	

func help(bin string) {
	banner()
	fmt.Printf("	-p, password argument\n" +
		   "	-w, words file argument\n" +
		   "	-h, print this panel and exit\n" +
		   "	-v, print version and exit\n" +
		   "usage: %s [password] | -h | -v\n", bin);
	os.Exit(0)
}

func print_version() {
	fmt.Printf("gopasscheck-%s\n", version);
	os.Exit(0)
}

func usage(bin string) {
	fmt.Fprintf(os.Stderr, "usage: %s [password] | -h | -v\n", os.Args[0]);
	os.Exit(1)
}

func main() {

	bin := os.Args[0]

	if (len(os.Args) == 1) {
		usage(bin)
	}

	if (os.Args[1] == "-h") { help(bin); }
	if (os.Args[1] == "-v") { print_version(); }
	
	var words string = ""
	var pass string = ""
	for i := 0; i < len(os.Args); i++ {
		if (os.Args[i] == "-w") {
			if (len(os.Args) >= i+1) {
				words = os.Args[i+1]
			} else {
				usage(bin)
			} }

		if (os.Args[i] == "-p" && len(os.Args) > i+1) {
			if (len(os.Args) > i) {
				pass = os.Args[i+1]
			} else {
				usage(bin)
			}
		}
	}

	if (pass == "") {
		usage(bin)
	}

	banner()
	fmt.Printf("\nWords: %s\n", words);
	fmt.Printf("Password: %s\n\n", pass);

	length_test(pass);
	fmt.Printf("\nCharacter set:\n")
	charset_test(pass);

	if (words != "") {
		fmt.Printf("\nWords:\n")
		words_test(pass, words) // 0 = words file, 1 = password dictionary
	}

	fmt.Printf("\nRecent years:\n")
	years_test(pass)

	fmt.Printf("\nCommon masks:\n")
	mask_test(pass)
}
