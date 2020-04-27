package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://google.com")
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	f, _ := os.Create("C:\\Users\\lekbad\\Desktop\\Folders\\Go\\webDevelopmentWithGo\\test.txt")
	w := bufio.NewWriter(f)
	fmt.Println(w, "Hello\n#{b}")
	w.Flush()
	fmt.Print(b)

}
