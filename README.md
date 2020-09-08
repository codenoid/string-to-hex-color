# string-to-hex-color

Generate hex color uniquely based on given string

![Example use of string-to-hex-color](https://github.com/codenoid/string-to-hex-color/blob/master/screenshot.png?raw=true)
_example use of string-to-hex-color applied to another library_

## Installation

```go
go get -u github.com/codenoid/string-to-hex-color
```

## Usage

```go
package main

import (
	"fmt"

	shex "github.com/codenoid/string-to-hex-color"
)

func main() {
	fmt.Println(shex.GenerateColor("Let's write Stable Software")) // #B25320
	fmt.Println(shex.GenerateColor("Let's write Stable Software")) // #B25320
}
```

