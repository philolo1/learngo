package main
import "testing"

func TestSubtract(t *testing.T) {
    t.Run("test subtracting positive numbers", func(t *testing.T) {
        result := subtract(4, 2)
        expectedResult := 2
        if result != expectedResult {
            t.Errorf("expected %d but got %d", expectedResult, result)
        }
    })

    t.Run("test subtracting negative numbers", func(t *testing.T) {
        result := subtract(-4, -2)
        expectedResult := -2
        if result != expectedResult {
            t.Errorf("expected %d but got %d", expectedResult, result)
        }
    })

    t.Run("test subtracting zero", func(t *testing.T) {
        result := subtract(4, 0)
        expectedResult := 4
        if result != expectedResult {
            t.Errorf("expected %d but got %d", expectedResult, result)
        }
    })
}
