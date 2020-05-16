package jobpool

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

//////////////// Example //////////////////

// TestJob - holds only an ID to show state
type TestJob struct {
	ID string
}

// Process - test process function
func (t *TestJob) Process() {
	fmt.Printf("Processing job '%s'\n", t.ID)
	time.Sleep(1 * time.Second)
}

func TestJobPool(t *testing.T) {
	queue := NewJobQueue(runtime.NumCPU())
	queue.Start()
	defer queue.Stop()

	for i := 0; i < 4*runtime.NumCPU(); i++ {
		queue.Submit(&TestJob{strconv.Itoa(i)})
	}
}
