package journal_hook

import "testing"

func TestStringifyEntries(t *testing.T) {
	input := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}

	output := stringifyEntries(input)
	if output["foo"] != "bar" {
		t.Fatalf("%v", output)
		t.Fatalf("expected value 'bar'. Got %q", output["foo"])
	}
	if output["baz"] != "123" {
		t.Fatalf("expected value '123'. Got %q", output["baz"])
	}
}
