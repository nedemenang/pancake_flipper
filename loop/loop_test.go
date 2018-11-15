package loop

import "testing"

func TestLoop(t *testing.T) {
	var tests = []struct {
		s        string
		k        int
		expected string
	}{
		{"---+-++-", 3, "3"},
		{"+++++", 4, "0"},
		{"-+-+-", 4, "IMPOSSIBLE"},
	}
	for _, test := range tests {
		if result := Loop(test.s, test.k); result != test.expected {
			t.Error("Test Failed: Expected {} to equal {}, recieved: {}", test.s, test.expected, result)
		} else {
			t.Logf("Test passed")
		}
	}
}


func benchmarkLoop(s string, k int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(s, k)
	}
}

func BenchmarkLoop1(b *testing.B) { benchmarkLoop("---+-++-", 3, b)}
func BenchmarkLoop2(b *testing.B) { benchmarkLoop("---++++-++--+--+-+--+++--+-+-++++-++-+-----+++-+-+++++-+++--++-+++++-+++++", 6, b)}
func BenchmarkLoop3(b *testing.B) { benchmarkLoop("-++-++--+-----+----+-+-+---+--+++++++-++++-++++++++--++++-+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++-++--+--+-----+----+-+-+---+--+++++++-++++-++++++++--++++-", 117, b)}
