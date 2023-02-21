package main
import "fmt"
import "math/big"
import "crypto/rand"


func PrivateKey(p *big.Int) *big.Int {
	randInt, err := rand.Int(rand.Reader, p)

  if (randInt.Cmp(big.NewInt(2)) <= 0) {
    randInt = big.NewInt(2)
  }

  if err != nil {
    panic("could not create random number")
	}

  return randInt
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
  var result big.Int
  return  result.Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
  public := PublicKey(private, p, g)
  return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
  return p.Exp(public2, private1, p)
}

func main() {
  fmt.Println("Hello world")
  fmt.Println("Number: ", PrivateKey(big.NewInt(10000)))
}
