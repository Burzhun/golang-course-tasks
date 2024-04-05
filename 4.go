 package main
import (
  "fmt"
  "bufio"
  "os"
)

type pos struct{
  x int
  y int
}

func addPathToTop(x,y int,c rune, cells [][]rune){
  if(x % 2==1){
    x--
    cells[y][x] = c
  }
  for y>0{
    y--
    cells[y][x] = c
  }
  for x>0{
    x--
    cells[0][x] = c
  }
}

func addPathToBottom(x,y,h,l int,c rune, cells [][]rune){
  if(x % 2==1){
    x++
    cells[y][x] = c
  }
  for y<h-1{
    y++
    cells[y][x] = c
  }
  for x<l-1{
    x++
    cells[y][x] = c
  }
}

func main() {
	var in *bufio.Reader
  var out *bufio.Writer
  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()
  
  var n int
  fmt.Fscanln(in, &n)
  var a,b pos
  for i:=0; i<n; i++{
    var h,l int
    fmt.Fscanln(in, &h,&l)
    cells := make([][]rune, h, h)
    for j:=0; j<h; j++{
      var s string
      cells[j] = make([]rune, l, l)
      fmt.Fscanln(in, &s)
      cells[j] = []rune(s)
      for k,c := range s{
        if(c=='A'){
          a = pos{x:k,y:j}
        }
        if(c=='B'){
          b = pos{x:k,y:j}
        }
      }
    }
    if a.y<b.y{
      addPathToTop(a.x,a.y, 'a', cells)
      addPathToBottom(b.x,b.y,h,l, 'b', cells)
    }
    if a.y>b.y{
      addPathToBottom(a.x,a.y,h,l, 'a', cells)
      addPathToTop(b.x,b.y, 'b', cells)
    }
    if(a.y==b.y){
      if(a.x>b.x){
        addPathToBottom(a.x,a.y,h,l, 'a', cells)
      addPathToTop(b.x,b.y, 'b', cells)
      }else{
        addPathToTop(a.x,a.y, 'a', cells)
      addPathToBottom(b.x,b.y,h,l, 'b', cells)
      }
    }
    for j:=0; j<h; j++{
      fmt.Fprintln(out,string(cells[j]))
    }
  }
  
}