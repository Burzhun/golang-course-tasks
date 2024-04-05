 package main
import (
  "fmt"
  "bufio"
  "os"
)
func check(s string) bool{
  var prev rune
  for i,c := range s {
    if i==0 && c !='M'{
      return false
      prev = c
      continue
    }
    switch c{
      case 'M':
        if prev=='R' || prev == 'M'{
          return false
        }
    case 'R':
      if prev!='M'{
        return false
      }
    
    case 'C':
      if prev!='R'{
        return false
      }
    case 'D':
      if prev!='M'{
        return false
      }
    }
    prev = c
  }
  if prev != 'D'{
    return false
  }
  return true
}

func main() {
	var in *bufio.Reader
  var out *bufio.Writer
  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()
  
  var n int
  fmt.Fscanln(in, &n)
  var s string
  for i:=0; i<n; i++{
    fmt.Fscanln(in, &s)
    if(check(s)){
      fmt.Fprintln(out,"YES")
    }else{
      fmt.Fprintln(out,"NO")
    }
  }
  
}