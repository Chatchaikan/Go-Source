package main
import "fmt"
 
func main() {
    var rows int
    var k int = 0
    fmt.Print("Enter number of rows :")
    fmt.Scan(&rows)     
    for i := 1; i <= rows; i++ {     
        k=0
        for space := 1; space <= rows-i; space++ {
            //fmt.Print("  ")  
            fmt.Print(" ")       
        }
        for {
            //fmt.Print("* ")
            fmt.Print("*")
            k++
            if(k == 2*i-1){             
                break
            }
        }       
        fmt.Println("")
    }
}