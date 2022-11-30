# redis
redis client for golang


### Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/eatmoreapple/redis"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(ctx)
	defer client.Close()

	if err := client.Ping(ctx).Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ok")
	}
}
```