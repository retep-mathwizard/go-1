#Badgers

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("badgers!")
}
```

##2.0

```go
package main

import (
  "fmt"
  c "github.com/skilstak/go/colors"
)
func main() {
  for i := 1; i <= 12; i++ {
    fmt.Println(c.B3 + "badger" + c.X)
  }
}
```

##3.0

```go
package main
import (
  "fmt"
  "time"
  c "github.com/skilstak/go/colors"
)
const timeSleep = 387
func main() {
  for i := 1; i <= 12; i++ {
    time.Sleep(time.Millisecond * timeSleep)
    fmt.Println(c.B3 + "badger" + c.X)
  }
}
```
##4.0
```go
package main

import (
  "fmt"
  "time"
  c "github.com/skilstak/go/colors"
)
func main() {
  const sleepTime = 387
  for i := 1; i <= 4; i++ {
    for i := 1; i <= 12; i++ {
      time.Sleep(sleepTime * time.Millisecond)
      fmt.Println(c.B3 + "badger" + c.X)
    }
    for i := 1; i <= 2; i++ {
      time.Sleep((sleepTime * 2) * time.Millisecond)
      fmt.Println(c.R + "mushroom" + c.X)
      }
 }
}
```
##5.0
```go

package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"time"
)

func main() {
	const snake = (`
		 /\ /\         /\ /\         /\ /\              _______   
		/ / \ \       / / \ \       / / \ \             \   _  \  
		/ /   \ \     / /   \ \     / /   \ \            /  /_\  \ 
		/ /    \ \    / /    \ \    / /    \ \           
	/ /     \ \   / /     \ \   / /     \ \           \  \_/   \ 
____________ / /       \ \ / /       \ \ / /       \ \___________\_____  /
/_____/_____/ \/         \/ \/         \/ \/         \/_____/_____/     \/ 
`)
	//const is a constant. Constant's never change.
	const sleepTime = 387
	for i := 1; i <= 4; i++ {
		for i := 1; i <= 12; i++ {
			time.Sleep(sleepTime * time.Millisecond)
			fmt.Println(c.B3 + "badger" + c.X)
		}
		for i := 1; i <= 2; i++ {
			time.Sleep((sleepTime * 2) * time.Millisecond)
			fmt.Println(c.R + "mushroom" + c.X)
		}
	}

	time.Sleep((sleepTime * 10) * time.Millisecond)
	fmt.Println(c.G + snake + c.X)

}
```
