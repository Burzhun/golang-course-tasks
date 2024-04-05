package main
import (
  "fmt"
  "bufio"
  "os"
  "time"
)

func LastDayOfMonth(t time.Time) time.Time {
    firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
    lastDay := firstDay.AddDate(0, 1, 0).Add(-time.Nanosecond)
    return lastDay
}

func printData(t time.Time) string{
    return t.Format("2006-01-02")
}

func main() {
    var in *bufio.Reader
  var out *bufio.Writer
  in = bufio.NewReader(os.Stdin)
  out = bufio.NewWriter(os.Stdout)
  defer out.Flush()
  
  var t,start,end string
  fmt.Fscanln(in, &t)
  fmt.Fscanln(in, &start, &end)
  startTime, _ := time.Parse("2006-01-02", start)
    endTime, _ := time.Parse("2006-01-02", end)
    next := startTime
    var outputs  []string

  fmt.Fprintf(out,"%s %s %s \n", start, end,t)
  
  switch t{
      case "MONTH":
	    i :=0
	    for next.Before(endTime){
	        //fmt.Fprintf(out,"debug %s %d", printData(next), len(outputs) )
	        next2 := next.AddDate(0,1,0)
	        if next2.After(endTime){
	            firstDay := time.Date(next.Year(), next.Month(), 1, 0, 0, 0, 0, time.UTC)
	            s1 := printData(firstDay)
	            s2 := printData(endTime)
	            fmt.Fprintf(out,"%d\n", i+1)
	            for _,s := range outputs{
	                fmt.Fprintf(out,"%s\n", s)
	            }
	            fmt.Fprintf(out,"%s %s\n", s1, s2)
	            
	        }else{
	            firstDay := time.Date(next.Year(), next.Month(), 1, 0, 0, 0, 0, time.UTC)
	            if i==0{
	                firstDay = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, time.UTC)
	            }
	            l := LastDayOfMonth(next)
	            s1 := printData(firstDay)
	            s2 := printData(l)
	            outputs = append(outputs, s1+" "+s2)
	        }
	        next=next2
	        i++
	    }
        break
      case "WEEK":
        w := int(next.Weekday())
        if w != 6{
            next = next.AddDate(0,0,w-6)
        }
        //fmt.Fprintf(out,"debug %s %d", printData(next.AddDate(0,0,7)), len(outputs) )
        if next.AddDate(0,0,7).After(endTime){
           // fmt.Fprintf(out,"debug %s %d", printData(next), len(outputs) )
            s1 := printData(startTime)
            s2 := printData(endTime)
            fmt.Fprintf(out,"%d\n", 1)
            fmt.Fprintf(out,"%s %s\n", s1, s2)
            return
        }
        i :=0
        
        for !next.After(endTime){
	        
	        next2 := next.AddDate(0,0,7)
	        //fmt.Fprintf(out,"debug %s %d \n", printData(next2), len(outputs) )
	        if next2.After(endTime){
	           // fmt.Fprintf(out,"debug2 %s %d \n", printData(next), len(outputs) )
	            s1 := printData(next)
	            s2 := printData(endTime)
	            fmt.Fprintf(out,"%d\n", i+1)
	            for _,s := range outputs{
	                fmt.Fprintf(out,"%s\n", s)
	            }
	            fmt.Fprintf(out,"%s %s\n", s1, s2)
	            
	        }else{
	            firstDay := time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, time.UTC)
	            if i==0{
	                firstDay = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, time.UTC)
	            }
	            s1 := printData(firstDay)
	            s2 := printData(next2.AddDate(0,0,-1))
	            outputs = append(outputs, s1+" "+s2)
	        }
	        next=next2
	        i++
	    }
        break
      case "REVIEW":
        summerStart := time.Date(next.Year(), 4, 1, 0, 0, 0, 0, time.UTC)
        summerEnd := time.Date(next.Year(), 9, 30, 0, 0, 0, 0, time.UTC)
         winterStart := time.Date(next.Year(), 10, 1, 0, 0, 0, 0, time.UTC)
        winterEnd := time.Date(next.Year(), 3, 31, 0, 0, 0, 0, time.UTC)
        i := 0
        for !next.AddDate(i,0,0).After(endTime){
            if next.After(winterEnd) && next.Before(winterStart){
                s1 := printData(next.AddDate(i,0,0))
	            s2 := printData(summerEnd.AddDate(i,0,0))
	            outputs = append(outputs, s1+" "+s2)
	            next = winterStart
	            continue
            }
            if next.After(summerEnd){
                s1 := printData(next.AddDate(i,0,0))
                next = summerStart
                i++
	            s2 := printData(next.AddDate(i,0,-1))
	            outputs = append(outputs, s1+" "+s2)
	            continue
            }
            if !next.After(winterStart){
                s1 := printData(next.AddDate(i,0,0))
	            s2 := printData(winterEnd.AddDate(i,0,0))
	            next=summerStart
	            outputs = append(outputs, s1+" "+s2)
	            continue
            }
        }
        fmt.Fprintf(out,"%d\n", len(outputs))
        for _,s := range outputs{
            fmt.Fprintf(out,"%s\n", s)
        }
        break
  }
  
  
  
}