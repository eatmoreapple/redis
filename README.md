# redis
redis client for golang



#### Example

```go
package main

import (
	"fmt"
	"github.com/eatmoreapple/redis"
	"log"
)

func main() {
	client, err := redis.Dial(":@tcp(localhost:6379)/0")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := client.Get()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()
	fmt.Println(conn.Set("foo", "bar")) // true <nil>
	fmt.Println(conn.Get("foo"))        // bar <nil>
}

```

