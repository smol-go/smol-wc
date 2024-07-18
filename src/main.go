package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
	"unicode/utf8"
)

type Info struct {
	numLines int
	numWords int
	numBytes int
	numChars int
}

func parse_args(args []string) (map[string]bool, []string) {
	options := make(map[string]bool)
	files := []string{}

	for _, arg := range args {
		if arg[0] == '-' {
			options[arg] = true
		} else {
			files = append(files, arg)
		}
	}

	return options, files
}

func get_info(file string) Info {
	bytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return get_info_from_bytes(bytes)
}

func print_info(info Info, file string, options map[string]bool) {
	if len(options) == 0 {
		fmt.Printf(" %d %d %d", info.numLines, info.numWords, info.numBytes)
	} else {
		if options["-l"] {
			fmt.Printf(" %d", info.numLines)
		}
		if options["-w"] {
			fmt.Printf(" %d", info.numWords)
		}
		if options["-m"] {
			fmt.Printf(" %d", info.numChars)
		}
		if options["-c"] {
			fmt.Printf(" %d", info.numBytes)
		}
	}

	if file != "" {
		fmt.Printf(" %s", file)
	}

	fmt.Println()
}

func read_bytes_from_stdin() []byte {
	bytes := []byte{}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytes = append(bytes, line...)
		if err == io.EOF {
			break
		}
	}

	return bytes
}

func get_info_from_bytes(bytes []byte) Info {
	info := Info{0, 0, len(bytes), utf8.RuneCount(bytes)}
	inWord := false

	for _, byte := range bytes {
		if unicode.IsSpace(rune(byte)) {
			if byte == '\n' {
				info.numLines++
			}
			if inWord {
				info.numWords++
			}
			inWord = false
		} else {
			inWord = true
		}
	}

	if inWord {
		info.numWords++
	}

	return info
}

func main() {
	options, files := parse_args(os.Args[1:])

	if len(files) == 0 {
		bytes := read_bytes_from_stdin()
		print_info(get_info_from_bytes(bytes), "", options)
		return
	}

	for _, file := range files {
		print_info(get_info(file), file, options)
	}
}
