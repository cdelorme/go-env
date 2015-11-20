package env

import (
	"os"
	"testing"

	"github.com/cdelorme/go-maps"
)

func TestPlacebo(t *testing.T) {
	t.Parallel()

	if !true {
		t.FailNow()
	}
}

func TestVar(t *testing.T) {
	t.Parallel()
	e := App{}

	// validate empty names are ignored
	e.Var("", "noname")
	e.Var("nokey", "")
	if len(e.envs) > 0 {
		t.Log("failed to ignore empty names or keys...")
		t.FailNow()
	}

	// test valid entries
	e.Var("valid", "ONE")
	e.Var("valid", "TWO")
	if len(e.envs) != 2 {
		t.Log("failed to accept registration of 2 env vars...")
		t.FailNow()
	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	e := App{}

	// register entries for each data-type
	e.Var("bool", "TEST_BOOL")
	e.Var("string", "TEST_STRING")
	e.Var("int", "TEST_INT")
	e.Var("float", "TEST_FLOAT")

	// manually set env vars
	os.Setenv("TEST_BOOL", "true")
	os.Setenv("TEST_STRING", "banana")
	os.Setenv("TEST_INT", "66")
	os.Setenv("TEST_FLOAT", "3.5")

	// parse
	o := e.Parse()

	// check length
	if len(o) != 4 {
		t.Logf("failed to parse expected values: %+v\n", o)
		t.FailNow()
	}

	// validate output values
	if b, _ := maps.Bool(o, false, "bool"); b != true {
		t.Log("Expected: true, Got: false")
		t.FailNow()
	}
	if s, _ := maps.String(o, "", "string"); s != "banana" {
		t.Logf("Expected: banana, Got: %s\n", s)
		t.FailNow()
	}
	if i, _ := maps.Int(o, 0, "int"); i != 66 {
		t.Logf("Expected: 66, Got: %d\n", i)
		t.FailNow()
	}
	if f, _ := maps.Float(o, 0, "float"); f != 3.5 {
		t.Logf("Expected: 3.5, Got: %f\n", f)
		t.FailNow()
	}
}
