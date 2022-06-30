package greetings

import(
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T){
	name:= "Alice"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg,err := Hello("Alice")
	if !want.MatchString(msg) || err != nil{
		t.Fatalf("Hello(%q) = %q, %v, want %q, nil", name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil{
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)	}
}