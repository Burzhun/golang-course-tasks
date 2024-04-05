package main

import (
	"fmt"
	"os"
	"io"
	//"strings"
	"bytes"
)

func printTree(out io.Writer, left string, path string, printFiles bool) {
	outputDirRead, _ := os.Open(path)
	list,_ := outputDirRead.ReadDir(0) 
	lastIndex := len(list)-1
	if(!printFiles){
		for i, e := range list {
			if e.IsDir() {
				lastIndex = i
			}		
		}
	}
	
	for i, e := range list {
		c := "├"
		b := i==lastIndex
		if b{
			c="└"
		}
		if e.IsDir() {
			fmt.Fprintln(out, left+c+"───"+e.Name())
			if b{
				printTree(out, left+"\t",path+string(os.PathSeparator)+e.Name(), printFiles)
			}else{
				printTree(out, left+"│\t",path+string(os.PathSeparator)+e.Name(), printFiles)
			}
		}else{
			if printFiles {
				info,_ := e.Info()
				size := info.Size()
				
				if size>0 {				
					s := fmt.Sprintf(c+"───%s (%db)", e.Name(), size)
					fmt.Fprintln(out,left+s)
				}else{
					s := fmt.Sprintf(c+"───%s (empty)", e.Name())
					fmt.Fprintln(out,left+s)
				}
				
			}
		}		
	}
}

func dirTree(out io.Writer, path string,  printFiles bool) error {
	printTree(out,"",path,printFiles)
	return nil
	
}

func main() {
	//out := os.Stdout
	out := new(bytes.Buffer)

	err := dirTree(out, "testdata", true)
	if err != nil {
		panic(err.Error())
	}

}
