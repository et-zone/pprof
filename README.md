# Introduction

pprof is a tool for visualization and analysis of profiling data.


# Config
### your go server open pprof
```
import (
	"log"
	"net/http"
	_ "net/http/pprof"
)
func main(){
	go func() {
		fmt.Println(http.ListenAndServe("localhost:60066", nil))
	}()
}
```

# Run
* run your service
* run this tool `./pprof`

* open browser to update base Data like cpu
    ```http://127.0.0.1:9000/target/fresh?arg=http://localhost:60066/debug/pprof/profile?seconds=1```

## eg
  `http://localhost:8080/debug/pprof/profile?seconds=1`

  `http://localhost:8080/debug/pprof/heap?seconds=1`

  `http://localhost:8080/debug/pprof/goroutine?seconds=1`

  `http://localhost:8080/debug/pprof/mutex?seconds=1`

  `http://localhost:8080/debug/pprof/allocs?seconds=1`

  `http://localhost:8080/debug/pprof/block?seconds=1`

## open browser to look View
`http://127.0.0.1:9000/target/`
