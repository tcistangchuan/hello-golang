package simple_factory

import (
	"testing"
)

func TestAPI_Say(t *testing.T) {
	t.Skip("456")
	t.Log(NewAPI().Say())
}
