package main
import(
  "fmt"
  "math"
)
func main(){
  a:=16
  a = int(math.Sqrt(float64(a)))
  fmt.Println(a)
}