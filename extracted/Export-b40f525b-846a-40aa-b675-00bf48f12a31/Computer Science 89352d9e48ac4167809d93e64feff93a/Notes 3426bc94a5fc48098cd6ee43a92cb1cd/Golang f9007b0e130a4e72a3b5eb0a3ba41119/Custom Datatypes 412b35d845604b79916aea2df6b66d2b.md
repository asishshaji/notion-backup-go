# Custom Datatypes

## Enum

```go
type LogEnum int

const (
	DEBUG LogEnum = iota + 1
	WARNING
	INFO
	ERROR
)

func (l LogEnum) String() string {
	return [...]string{"DEBUG", "WARNING", "INFO", "ERROR"}[l]
}

func main() {

	var l LogEnum = DEBUG
	fmt.Println(l.String())
}
```