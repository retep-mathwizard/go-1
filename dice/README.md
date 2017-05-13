##Dice
```go
package main

import (
  c "github.com/skilstak/go-colors"
  "fmt"
)

func main() {
  fmt.Println(c.Rc() + " Dice! " + c.X)
}
```
##Dice 2.0
```go

package main

import (
  c "github.com/skilstak/go-colors"
  "github.com/skilstak/go-random"
  "fmt"
)

func main() {
  numbers := []int{1,2,3,4,5,6}
  //[]int tells golang that it will contain numbers.
  //[]string tells it that it will contain strings.

  randNum := random.Choice(numbers)

  fmt.Println(c.Rc() + str(randNum) + c.X)
  // when printing, all items have to be of the same type

}
```
##Dice 3.0
```go
package main

import (
  c "github.com/skilstak/go-colors"
  "github.com/skilstak/go-random"
  "fmt"
)

func main() {
const (
    side1 = (`
 ------
|      |
|  壹  |
|     1|
 ------
    `)

    side2 = (`
 ------
|      |
|  貳  |
|     2|
 ------
    `)

    side3 = (`
 ------
|      |
|  叁  |
|     3|
 ------
    `)

    side4 = (`
 ------
|      |
|  肆  |
|     4|
 ------
    `)

    side5 = (`
 ------
|      | 
|  伍  |
|     5|
 ------
    `)

    side6 = (`
 ------
|      | 
|  陸  |
|     6|
 ------
    `)
  )
//const means constant. Constants never change throughout the program
  die := []string{side1,side2,side3,side4,side5,side6}

  randSide := random.Choice(die)

  fmt.Println(c.Rc() + randSide + c.X)
}
```
##Dice 4.0
```go
package main

import (
  "github.com/skilstak/go-input"
  c "github.com/skilstak/go-colors"
  "github.com/skilstak/go-random"
  "fmt"
)

func main() {
const (
    side1 = (`
 ------
|      |
|  壹  |
|     1|
 ------
    `)

    side2 = (`
 ------
|      |
|  貳  |
|     2|
 ------
    `)

    side3 = (`
 ------
|      |
|  叁  |
|     3|
 ------
    `)

    side4 = (`
 ------
|      |
|  肆  |
|     4|
 ------
    `)

    side5 = (`
 ------
|      | 
|  伍  |
|     5|
 ------
    `)

    side6 = (`
 ------
|      | 
|  陸  |
|     6|
 ------
    `)
  )
  die := []string{side1,side2,side3,side4,side5,side6}
  for {
    randSide := random.Choice(die)

    fmt.Println(c.Rc() + randSide + c.X)
    input.Ask("")
  }
}
