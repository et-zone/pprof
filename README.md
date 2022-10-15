# Introduction

pprof is a tool for visualization and analysis of profiling data.

# Run
`./pprof`

### open browser to update like cpu
    `http://127.0.0.1:9000/target/fresh?arg=http://localhost:8080/debug/pprof/profile?seconds=1`

### open browser to look pic or table
  `http://127.0.0.1:9000/target/`

# Config
  ### your go server open pprof
```
import (
	"log"
	"net/http"
	_ "net/http/pprof"
)
func main(){
	log.Println(http.ListenAndServe("localhost:8080", nil))
}
```

## eg
  `http://localhost:8080/debug/pprof/profile?seconds=1`

  `http://localhost:8080/debug/pprof/heap?seconds=1`

  `http://localhost:8080/debug/pprof/goroutine?seconds=1`

  `http://localhost:8080/debug/pprof/mutex?seconds=1`

  `http://localhost:8080/debug/pprof/allocs?seconds=1`

  `http://localhost:8080/debug/pprof/block?seconds=1`