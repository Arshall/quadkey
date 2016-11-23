# quadkey
Generate Quad Keys

### Setup
<pre>go get github.com/onef9day/quadkey</pre>

### Example
```go
package main

import "fmt"
import "github.com/onef9day/quadkey"

func main() {
    quadKey := quadkey.Init(rand.NewSource(time.Now().UnixNano()))
    quadKey.Alphabet = "ABCDEF0123456789"
    quadKey.Segments = 4
    for i := 1; i <= 20; i++ {
        fmt.Println(quadKey.Key())
    }
}
```