# Cache

In-memory cache

## Example

```go
package main

import (
	"fmt"
	cache "github.com/ernur-eskermes/go-homeworks/2-cache-ttl"
	"log"
	"time"
)

func main() {
	c := cache.New()
	c.Set("userId", 42, time.Second*5)
	userId, err := c.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 6) // прошло 5 секунд

	userId, err = c.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}
}
```