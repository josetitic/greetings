package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q,%v, quiere coincidencia para %q,nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q,%v, quiere "", error `, msg, err)
	}
}

func TestHellosNames(t *testing.T) {
	names := []string{"Sami", "Mary", "José", "Juan"}
	msg, err := Hellos(names)
	if err != nil {
		t.Fatalf(`Hellos(names) = %q,%v, quiere coincidencia para los nombres`, msg, err)
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{"", "", ""}
	msg, err := Hellos(names)
	if err == nil {
		t.Fatalf(`Hellos("") = %q,%v, quiere "", error `, msg, err)
	}
}
