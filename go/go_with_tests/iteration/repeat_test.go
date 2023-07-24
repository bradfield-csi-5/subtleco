package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeating 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
    assertCorrectInteger(t, got, want)
	})
  t.Run("repeating 10 times", func(t *testing.T){
    got := Repeat("b", 10)
    want := "bbbbbbbbbb"
    assertCorrectInteger(t, got, want)
  })

}

func assertCorrectInteger(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
