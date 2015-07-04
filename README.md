# pdconv
packed decimal (packed bcd) converter.

## Usage
```go
package main

import (
	"fmt"
	"github.com/donke/pdconv"
)

func main() {
	i, err := pdconv.Ptoi([]byte{0x01, 0x23, 0x4C})
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}
```

## Installation
```
$ go get github.com/donke/pdconv
```

## License

MIT

## Author

Ken Kudo (aka.kudoken@gmail.com)
