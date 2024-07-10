package main

import (
	"fmt"
	"io"
	"os"
)

var path = "./test.txt"

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func createFile() {
     // detect if file exists
    var _, err = os.Stat(path)

    // create new file if it's not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) { return }
        defer file.Close()
    }

    fmt.Println("==> file created", path)
}
func writeFile() {
    // open filw with READ & WRITE access
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) { return }
    defer file.Close()

    // write string to the file
    _, err = file.WriteString("Hello World\n")
    if isError(err) { return }
    _, err = file.WriteString("write from the code\n")
    if isError(err) { return }

    // save the changes with Sync function
    err = file.Sync()
    if isError(err) { return }

    fmt.Println("==> File Edited")
}

func readFile() {
    // Open file with Read only access
    var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
    if isError(err) { return }
    defer file.Close()

    // Read the string and store it into our array of byte
    var text = make([]byte, 1024)
    for {
        n, err := file.Read(text)
        if err != io.EOF {
            if isError(err) { break }
        }
        if n == 0 {
            break
        }
    }
    if isError(err) { return }

    fmt.Println("==> File Successfully read")
    fmt.Println(string(text))
}
func deleteFile() {
    var err = os.Remove(path)
    if isError(err) { return }

    fmt.Println("==> file deleted")
}



func File() {
    createFile()
	writeFile()
	readFile()
	deleteFile()
}