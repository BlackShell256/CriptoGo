package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mkmueller/aes256"
)

func main() {
	color.Blue("Escoge lo que necesites\n\n")
	color.Red(" [1] Encriptar")
	color.Red(" [2] Desencriptar\n\n")

	fmt.Print(color.RedString("> "))
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	op := buf.Text()

	color.Blue("\nIngresa la contraseña de encriptado o desencriptado\n\n")

	fmt.Print(color.RedString("> "))
	buf2 := bufio.NewScanner(os.Stdin)
	buf2.Scan()
	pass := buf2.Text()

	switch op {
	case "1":
		Enc(pass)
	case "2":
		Dec(pass)
	default:
		color.Red(" [X] Opcion incorrecta")
		time.Sleep(2 * time.Second)
		main()
	}

	color.Blue("\n\n [*] Finalizado")
}

func GetFileEx() (dir string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	end := strings.HasPrefix(ex, os.Getenv("TEMP"))

	if end {
		dir = filepath.Base(strings.TrimSuffix(os.Args[0], ".exe") + ".go")
	} else {
		dir = ex
	}

	return
}

func GetFiles() (List []string) {
	Ex := GetFileEx()
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.Name() == Ex {
			continue
		} else if f.IsDir() {
			continue
		}
		List = append(List, f.Name())
	}

	return
}

func Enc(Key string) {

	for _, File := range GetFiles() {
		content, _ := os.ReadFile(File)
		Out, _ := aes256.Encrypt(Key, content)
		err := OpenFile(File, Out)
		if err != nil {
			continue
		}

	}
}

func OpenFile(name string, content []byte) (err error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}

	f.Write(content)
	err = f.Close()
	if err != nil {
		return
	}
	return
}

func Dec(Key string) {
	for _, File := range GetFiles() {
		content, _ := os.ReadFile(File)
		Out, _ := aes256.Decrypt(Key, content)
		err := OpenFile(File, Out)
		if err != nil {
			continue
		}

	}
}
