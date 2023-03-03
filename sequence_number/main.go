package main

import (
  "fmt"
  "github.com/go-redis/redis/v8"
  "context"
  "log"
  "time"
  "sync"
  "os"
	"os/signal"
  "strconv"
  "net/http"
	"github.com/Shopify/sarama"
)

var (
	kafkaBrokers = []string{"localhost:9093"}
	KafkaTopic = "sequeceNumber"
	enqueued int
)

const TIME = 10 * time.Second

func countRedis(wg *sync.WaitGroup) {
  ctx := context.Background()
  client := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "redispassword", // no password set
    DB:       0,  // use default DB
  })

  counter := 0
  done := time.After(TIME)
  for {
    select {
    case <-done:
      log.Printf("Redis: %v", counter)
      wg.Done()
      return
    default:
      client.Incr(ctx, "myKey")
      counter += 1
      // perform the operation here
    }
  }
}

func countIterator(wg *sync.WaitGroup) {
  counter := 0
  done := time.After(TIME)
  for {
    select {
    case <-done:
      log.Printf("Counter: %v", counter)
      wg.Done()
      return
    default:
      counter += 1
      // perform the operation here
    }
  }
}

func WriteFile(counter int) {
  file, err := os.Create("example.txt")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  // write some text to the file
  text := strconv.Itoa(counter)
  _, err = file.WriteString(text)
  if err != nil {
    panic(err)
  }
}
func countFile(wg *sync.WaitGroup) {
  counter := 0
  done := time.After(TIME)
  for {
    select {
    case <-done:
      log.Printf("Written to file: %v", counter)

      wg.Done()
      return
    default:
      counter += 1
      WriteFile(counter)
      // perform the operation here
    }
  }
}

func countHttp(wg *sync.WaitGroup) {
  counter := 0
  done := time.After(TIME)
  for {
    select {
    case <-done:
      log.Printf("Random user: %v", counter)

      wg.Done()
      return
    default:
      counter += 1
      response, err := http.Get("https://randomuser.me/api/")
      if err != nil {
        panic(err)
      }
      defer response.Body.Close()
      // perform the operation here
    }
  }
}

func setupProducer() (sarama.AsyncProducer, error){

	config := sarama.NewConfig()
	return sarama.NewAsyncProducer(kafkaBrokers, config)
}

func countKafka(wg *sync.WaitGroup) {
	producer, err := setupProducer()
	if err != nil {
		panic(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

  counter := 0
  done := time.After(TIME)
  message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder(fmt.Sprintf("%v", counter))}
  for {
    select {
    case <-done:
      log.Printf("Kafka Messages: %v", counter)
			producer.AsyncClose() // Trigger a shutdown of the producer.
      wg.Done()
      return
		case producer.Input() <- message:
      counter+=1
      message= &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder(fmt.Sprintf("%v", counter))}
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
    default:
    }
  }
}

func main() {
  var wg sync.WaitGroup

  wg.Add(5)

  go countRedis(&wg)
  go countIterator(&wg)
  go countFile(&wg)
  go countHttp(&wg)
  go countKafka(&wg)
  wg.Wait()
}
