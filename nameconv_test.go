package structexport

import "testing"

func TestNameconv(t *testing.T) {
	tab := []struct {
		i     int
		input string
		want  string
	}{
		{1, "", ""},
		{2, "a", "A"},
		{3, "A", "A"},
		{4, "abc", "ABC"},
		{5, "foo", "FOO"},
		{6, "Foo", "FOO"},
		{7, "FooBar", "FOO_BAR"},
		{8, "fooBar", "FOO_BAR"},
		{9, "FOOBar", "FOOBAR"},
		{10, "QuiteLongName", "QUITE_LONG_NAME"},
		{11, "SQLite", "SQLITE"},
		{12, "TCPProxy", "TCPPROXY"},
		{13, "SNAKESEverywhere", "SNAKESEVERYWHERE"},
	}

	for _, tc := range tab {
		have := nameconv(tc.input)

		if have != tc.want {
			t.Errorf("tc#%d mismatch:\nhave %s\nwant %s", tc.i, have, tc.want)
		}
	}
}
