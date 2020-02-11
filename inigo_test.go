package configo
import "testing"
import "fmt"
func TestInigo(t *testing.T){
	a:=NewIni("temp.ini")
	fmt.Println(a)
	fmt.Println(a.Get("xxx","qwr"))
}