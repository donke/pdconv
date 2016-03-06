# pdconv
packed decimal (packed bcd) converter.

## Usage
```go
package main

import (
	"fmt"
	"github.com/kuked/pdconv"
)

func main() {
	i, err := pdconv.Ptoi([]byte{0x01, 0x23, 0x4C})
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

    b, err := pdconv.Itop(i, 3)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%x\n", b)
}
```

## Installation
```
$ go get github.com/kuked/pdconv
```

## License

MIT

## Author

Ken Kudo
