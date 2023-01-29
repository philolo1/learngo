package main

import "fmt"
import "strings"
import "unicode"


type PointCalculation struct {
    points int
    str string
}

var pArr []PointCalculation = []PointCalculation{
  {points: 1, str: "AEIOULNRST"},
  {points: 2, str: "DG"},
  {points: 3, str: "BCMP"},
  {points: 4, str: "FHVWY"},
  {points: 5, str: "K"},
  {points: 8, str: "JX"},
  {points: 10, str: "QZ"},
}

func getScore(p PointCalculation, char rune) int {
  if strings.ContainsRune(p.str,char) {
    return p.points
  }

  return 0
}

func Score(val string) int {
  var sum = 0
  if len(val) == 0 {
    return 0
  }
  for _, char := range val {
    for _,p := range pArr {
      // fmt.Println(string(char))
      sum += getScore(p, unicode.ToUpper(char))
    }
  }
  return sum
}


func main() {
  fmt.Println("Scrabble Project")
  fmt.Println("Calculate", Score(""))
  fmt.Println("Calculate", Score("abcdefghijklmnopqrstuvwxyz"))
}
