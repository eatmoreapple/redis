# redis
redis client for golang



#### Example

```go
package main

import (
	"fmt"
	"github.com/eatmoreapple/redis"
)

func main() {
	client, err := redis.NewClient("tcp", "localhost:6379", 2)
	if err != nil {
		panic(err)
	}
	conn, err := client.Conn()
	if err != nil {
		panic(err)
	}
	defer conn.Release()
	fmt.Println(conn.Set("foo", "bar")) // true <nil>
	fmt.Println(conn.Get("foo"))        // bar <nil>
}
```

