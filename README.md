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
output

```
0A11-9D2C-A288-5DE2
F229-B8E3-E0B1-0565
F4A4-77E0-D210-910C
D64D-BC3B-5870-DD87
A571-C1D5-157F-2286
527F-EE93-EFFA-43A0
13DF-452E-CCB4-6202
6019-5380-FCC7-E25E
0D36-091E-B66E-FB9B
B3CD-6FE0-0D9A-837D
1188-3D7A-63D3-7FAF
E3EB-5E00-58B1-4093
5DBC-D0B1-C69C-0214
E514-83C4-A65D-119E
241C-92BD-7F32-45BA
55A8-7688-2265-72E1
CF2C-6743-6FB1-8B70
A919-398A-9DE9-DC91
3503-56F6-BA3D-7CF2
193F-FD95-3635-3FD5
```