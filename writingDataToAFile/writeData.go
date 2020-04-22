package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new file
	f, _ := os.Create("C:\\Users\\lekbad\\Desktop\\Folders\\Go\\writingDataToAFile\\newFile.txt")
	// Start writing
	wr := bufio.NewWriter(f)
	name := "Lekha"
	// Fprint statements are used to write data to a file
	fmt.Fprintf(wr, "Hello,%s.\n", name)
	fmt.Fprintln(wr, "Go is pretty neat!")
	fmt.Fprint(wr, "Ok, bye.")
	// Don't forget to flush, or the data won't be written.
	wr.Flush()
}
