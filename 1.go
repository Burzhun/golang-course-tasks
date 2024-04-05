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
  l := 10
  fmt.Fscanln(in, &n)
  for i:=0; i<t; i++{
	fmt.Fscanln(in, &n, &p)
	s:=0.0
	for i:=0; i<t; i++{
		var a int
		fmt.Fscanln(in, &a)
		var f float64
		f = 1.0*a*p/100
		s += f- float(int(f))		
	}
	fmt.Fprintf(out,"%f\n" s)
  }
  
  
}