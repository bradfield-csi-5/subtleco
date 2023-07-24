package iteration

func Repeat(char string, number int) string {
  var repeated string
  for i := 0; i < number; i++ {
    repeated += char
  }
  return repeated
}
