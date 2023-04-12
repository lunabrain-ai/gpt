# aicli
Use ChatGPT in your terminal. Kind of feels like a bash utility?

## Installation
```bash
go get github.com/lunabrain-ai/aicli
```
or if you want the binary release, go to [releases](https://github.com/lunabrain-ai/aicli/releases/).

## Usage
````bash
export OPENAI_API_KEY="<your openai api key>"
aicli "Write me a go function that prints 'Hello World'"
Here's an example Go function that prints "Hello World" to the console:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
````

## Hack

### Run locally
```bash
go run main.go
```

### Build
```bash
go build main.go
```
