package consts

const (
	IP        = "127.0.0.1"
	Port      = "9000"
	profile   = "/debug/pprof/profile" //cpu的占用信息
	goroutine = "/debug/pprof/goroutine"
	heap      = "/debug/pprof/heap"
	mutex     = "/debug/pprof/mutex"  //持有锁的堆栈信息
	allocs    = "/debug/pprof/allocs" //内存分配采样

	trace = "/debug/pprof/trace" //当前程序执行的trace
	//symbol="/debug/pprof/symbol"//请求中列出的程序计数器,无需刷新数据源
	block        = "/debug/pprof/block"
	threadcreate = "/debug/pprof/threadcreate" //导致操作系统创建新增的线程的堆栈信息
	BaseDataUrl  = "http://" + IP + ":" + Port + "/target/fresh?arg="
)

const (
	JobProfileCPU = "/debug/pprof/profile?seconds="

	JobHeap = "/debug/pprof/heap?seconds="

	JobGoroutine = "/debug/pprof/goroutine?seconds=" //peek

	JobMutex = "/debug/pprof/mutex?seconds="

	JobAllocs = "/debug/pprof/allocs?seconds="

	JobBlock       = "/debug/pprof/block?seconds="
	BaseReqDataURL = "http://" + IP + ":" + Port + "/target/fresh?arg="
)

// can change data
const (
	ExportAddr     = "http://localhost:60066"
	JobStepSeconds = "2" //?seconds=xxx
)
