package main
import "fmt"
import "strings"

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Valid(id string) bool {

  str := ReverseString(strings.ReplaceAll(id," ",""))
  if len(str) < 2 {
    return false
  }

  sum := 0

  for index,char := range str {
    number := int(char - '0')
    if number < 0 || number > 9 {
      return false
    }

    if index % 2 == 1 {
      if number*2 > 9 {
        sum += number*2 - 9
      } else {
        sum += number*2
      }
    } else {
      sum += number
    }
  }

  return sum % 10 == 0
}

func main() {
  fmt.Println("Hello world")
  fmt.Println("Valid: ", Valid("4539 3195 0343 6467"))
}
