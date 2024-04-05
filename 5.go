package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Folder struct {
	Dir	 string     `json:"dir"`
	Files []string  `json:"files"`
	Folders []Folder `json:"folders"`
}

func getCount(f Folder,n int, infected bool) int{
  t:=0
  if(!infected){
    for _,s := range f.Files{
      if strings.HasSuffix(s,".hack"){
        infected = true
        break
      }
    }
  }
  if(infected){
    t += len(f.Files)
  }  

	
  for _,folder := range f.Folders {
    t += getCount(folder, n, infected)
  }
	
	return n+t
  
}

func main() {	
	
	var in *bufio.Reader
  var out *bufio.Writer
  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()
  
  var n int
  fmt.Fscanln(in, &n)
  var strBuilder strings.Builder
  for i:=0; i<n; i++{
    var m int
    fmt.Fscanln(in, &m)
    strBuilder.Reset()
    for j:=0; j<m; j++{
      s2, _ := in.ReadString('\n')
      strBuilder.WriteString(s2)
    }

    var folder1 Folder
  	err := json.Unmarshal([]byte(strBuilder.String()), &folder1)
  
  	if err != nil {
  
  		fmt.Fprintln(out,"0")
  	}else{
  	  fmt.Fprintln(out, getCount(folder1, 0, false))
  	}
    
  }	
	

}
