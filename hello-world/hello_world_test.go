package greeting

import (
	"testing"
)

func TestHello(t *testing.T) {
	if msg := sayHello(); msg != "Hello, World!" {
		t.Fatalf(`Hello() = %q, want "Hello, World!", error`, msg)
	}
}
