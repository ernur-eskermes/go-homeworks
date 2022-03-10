# Cache

In-memory cache

## Example

```go
package main

import (
	"fmt"
	cache "github.com/ernur-eskermes/go-homeworks/1-in-memory-cache"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)
	userId, err := c.Get("userId")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(userId)
	}

	c.Delete("userId")
}
```