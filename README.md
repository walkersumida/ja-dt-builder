# jadtbuilder

## Getting Started

```sh
go get -u github.com/walkersumida/jadtbuilder
```

## Usage

```go
import 	"github.com/walkersumida/jadtbuilder"

now := time.Now()
fmt.Println(now)
// => 2021-09-05 16:00:00 +0900 JST

var datetime string

datetime = jadtbuilder.Build("2021 9 4 13")
fmt.Println(datetime)
// => "9月4日(土) 13:00"

datetime = jadtbuilder.Build("21")
fmt.Println(datetime)
// => "9月5日(日) 21:00"

datetime = jadtbuilder.Build("4 21")
fmt.Println(datetime)
// => "10月4日(月) 21:00"
```