package main
import "testing"
func TestN(t *testing.T) {
    name:=getName()
    if name!="World" {
        t.Error("Respone from getName is unexpected value")
    }
}
