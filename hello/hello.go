package main

import "fmt"
import "errors"

func subtract(a int,b int) int {
  return a - b;
}

func subtractThree(a int,b int, c int) int {
  return a - b - c;
}

func multiply(a int,b int) int {
  return a * b;
}

func isOdd(a int) bool {
  if a % 2 == 0{
    return false
  }
  return true

}

func hasEvenLength(str string) bool {
  return len(str) % 2 == 0
}

func hasAllOnes(str string) bool {
  val := true

  for _, c := range str {
    if c != '1' {
      val = false
    }
  }

  return val
}

func createName(a string, b string) string {
   return fmt.Sprintf("%s %s", a, b)
}

func myArray(a []int) {
  a[0] = 3
}

func myArray2(a *[]int) {
  (*a)[1] = 3
}

func max(a []int) (int, error) {
  if len(a) == 0 {
    return 0, errors.New("empty array")
  }
  var m = 0

  for _, x := range a {
    if x > m {
      m = x
    }
  }

  return m, nil
}

func filterOdd(a []int) []int {
  var list = []int{}

  for _, x := range a {
    if x % 2 == 1 {
      list = append(list, x)
    }
  }
  return list
}

type Person struct {
  FirstName string
  LastName  string
  Age uint
}

func newPerson(firstName, lastName string, age uint) *Person {
  return &Person{FirstName: firstName, LastName: lastName, Age: age}
}



func (p Person) String() string {
  return fmt.Sprintf("(%s, %s, %d)", p.FirstName, p.LastName, p.Age)
}

func (p *Person) hello() {
  fmt.Println("HELLO!!!", p.FirstName)
}

func youngest(arr []Person) (*Person, error) {
  if len(arr) == 0 {
    return nil, errors.New("Empty Array")
  }

  index := 0

  for i, el := range arr {
    if el.Age < arr[index].Age {
      index = i
    }
  }

  fmt.Println("el: ", arr[index])
  return &arr[index], nil
}

type Name struct {
  Title string
  First string
  Last  string
}

type DOB struct {
  Date string
  Age  int
}

type User struct {
  Gender string
  Name   Name
  DOB    DOB
}




func main() {
    fmt.Println("Hello, World!")
    fmt.Println("10-2", subtract(10, 2))
    fmt.Println("10-2-3", subtractThree(10, 2,3))
    fmt.Println("3*3", multiply(3, 3))
    fmt.Println("isOdd 3", isOdd(3))
    fmt.Println("isOdd 8", isOdd(8))
    fmt.Println("hasEvenLength test", hasEvenLength("test"))
    fmt.Println("hasEvenLength est", hasEvenLength("tst"))
    fmt.Println("hasAllOnes 1", hasAllOnes("11111"))
    fmt.Println("hasAllOnes 1", hasAllOnes("1x1111"))
    fmt.Println("hasAllOnes 1", hasAllOnes(""))
    fmt.Println("Name:", createName("Patrick", "Klitzke"))

    var arr = []int{10,11,12}
    fmt.Println("GO array", arr)
    myArray(arr)
    fmt.Println("GO array", arr)
    myArray2(&arr)

    maxValue, err := max(arr)
    fmt.Println("GO array", maxValue, err)
    maxValue2, err2 := max([]int{})
    fmt.Println("GO array", maxValue2, err2)

    person := newPerson("Patrick", "Klitzke", 30)
    fmt.Println("Person", person)
    person.hello()
    fmt.Println("Person", person)

    people := []Person{
      {FirstName: "Peter", Age: 12},
      {FirstName: "Maria", Age: 3},
      {FirstName: "Philip", Age: 15},
    }

    y, err := youngest(people)

    fmt.Println("youngest ", y)

    user  := User{
      Gender: "male",
      Name: Name{
        Title: "mr",
        First: "brad",
        Last:  "gibson",
      },
      DOB: DOB{
        Date: "1993-07-20T09:44:18.674Z",
        Age:  26,
      },
    }

    fmt.Println("User", user)

}
