package main
import (
  "fmt"
  "bufio"
  "os"
)

func main() {
	var in *bufio.Reader
  var out *bufio.Writer
  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()
  
  var n,t,p int
  fmt.Fscanln(in, &t)
  for i:=0; i<t; i++{
  	fmt.Fscanln(in, &n, &p)
  	s:=0.0
  	for j:=0; j<n; j++{
  		var a int
  		fmt.Fscanln(in, &a)
  		var f float64
  		f = 1.0*float64(a*p)/100
  		s += f- float64(int(f))		
  	}
  	fmt.Fprintf(out,"%.2f\n", s)
  }
  
  
}