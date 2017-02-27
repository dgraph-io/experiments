/*
DIR=$HOME/rdbdemo
go build . && rm -Rf $DIR && mkdir -p $DIR && ./rdbdemo -dir $DIR
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/dgraph/rdb"
	"github.com/dgraph-io/dgraph/store"
	"github.com/dgraph-io/dgraph/x"
)

var (
	dir = flag.String("dir", "", "Directory to write data")
)

func ps() int {
	pid := os.Getpid()
	cmd := fmt.Sprintf("ps -ao rss,pid | grep %v", pid)
	buf, err := exec.Command("bash", "-c", cmd).Output()
	x.Check(err)
	s := strings.TrimSpace(string(buf))
	tokens := strings.Split(string(s), " ")
	kbs, err := strconv.Atoi(tokens[0])
	x.Check(err)
	return kbs
}

// This does not seem to have any memory leak. Memory usage fluctuates around 160000K for a
// really really long time.
func putDemo() {
	x.AssertTrue(len(*dir) > 0)
	st, err := store.NewStore(*dir)
	x.Check(err)
	defer st.Close()

	for i := 0; i < 100000000; i++ {
		if (i % 100000) == 0 {
			fmt.Printf("Written %d keys\n", i)
		}
		key := []byte(fmt.Sprintf("key%09x", i%10000))
		// Intentionally reallocate value.
		val := bytes.Repeat([]byte("v"), 5000*5+9)
		valSuffix := val[5000*5:]
		x.AssertTrue(9 == copy(valSuffix, []byte(fmt.Sprintf("%09x", i))))
		x.Check(st.SetOne(key, val))
	}
}

// This does not seem to have any major memory leak. Over 1 minute, it increases by 4K. Might be
// just due to memory fragmentation?
func getDemo() {
	x.AssertTrue(len(*dir) > 0)
	st, err := store.NewStore(*dir)
	x.Check(err)
	defer st.Close()

	for i := 0; i < 100000; i++ {
		key := []byte(fmt.Sprintf("key%09x", i))
		val := bytes.Repeat([]byte("v"), 5000*5+9)
		valSuffix := val[5000*5:]
		x.AssertTrue(9 == copy(valSuffix, []byte(fmt.Sprintf("%09x", i))))
		x.Check(st.SetOne(key, val))
	}

	var sum int
	for i := 0; i < 100000000; i++ {
		if (i % 100000) == 0 {
			fmt.Printf("Read %d keys\n", i)
		}
		key := []byte(fmt.Sprintf("key%09x", i%100000))
		slice, err := st.Get(key)
		x.Check(err)
		sum += slice.Size()
		slice.Free() // It is crucial to deallocate this. Otherwise, C memory will blow up very fast.
	}
	fmt.Printf("Size sum %d\n", sum)
}

// We see some memory leak here but it doesn't seem severe. Over ~10 minutes, it grows by 100M.
// Grows to 5.1% after 53 mins.
func writeBatchDemo() {
	x.AssertTrue(len(*dir) > 0)
	st, err := store.NewStore(*dir)
	x.Check(err)
	defer st.Close()

	for i := 0; i < 100000; i++ {
		key := []byte(fmt.Sprintf("key%09x", i))
		val := bytes.Repeat([]byte("v"), 5000*5+9)
		valSuffix := val[5000*5:]
		x.AssertTrue(9 == copy(valSuffix, []byte(fmt.Sprintf("%09x", i))))
		x.Check(st.SetOne(key, val))
	}

	var sum int
	wb := rdb.NewWriteBatch()
	//	defer wb.Destroy()

	for i := 0; i < 100000000; i++ {
		if (i % 100000) == 0 {
			fmt.Printf("Written %d keys\n", i)
		}
		if (i % 1000) == 0 { // Smaller batch size so that we never use too much mem.
			st.WriteBatch(wb)
			wb.Clear()
			//			wb.Destroy()
			//			wb = rdb.NewWriteBatch()
		}
		key := []byte(fmt.Sprintf("key%09x", i))
		val := bytes.Repeat([]byte("v"), 5000*5+9)
		valSuffix := val[5000*5:]
		x.AssertTrue(9 == copy(valSuffix, []byte(fmt.Sprintf("%09x", i))))
		wb.Put(key, val)
	}
	fmt.Printf("Size sum %d\n", sum)
}

func allocDemo() {
	n := 50000
	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = make([]byte, 10000)
	}
	fmt.Println("Allocated mem")
	var sum int
	for {
		for i := 0; i < n; i++ {
			for _, b := range a[i] {
				// Explicitly iterate over each a[i] so that the mem doesn't get paged to disk and ps
				// will include that.
				sum += int(b)
			}
			// As expected, the memory is not released. However, both memstats and ps take this memory
			// into account. In our use case, it is a discrepancy between memstats and ps. So it can't
			// be due to this.
			a[i] = a[i][:10]
		}
		// We find that the memory remains allocated, as expected, and memstats and ps both reflect that.
		debug.FreeOSMemory()
		time.Sleep(time.Second)
	}
	fmt.Printf("Size sum %d\n", sum)
}

// I think this explains what we are seeing.
/*
Allocated mem
msMem=514163Kb psMem=31680Kb
msMem=514204Kb psMem=31688Kb
msMem=514245Kb psMem=31692Kb
msMem=514285Kb psMem=31696Kb
Filled mem
msMem=514325Kb psMem=564904Kb
msMem=514366Kb psMem=564908Kb
msMem=514406Kb psMem=564920Kb
msMem=514447Kb psMem=564920Kb
msMem=514487Kb psMem=564924Kb
Resizing slices
msMem=514763Kb psMem=565188Kb
msMem=2450Kb psMem=565788Kb
msMem=2450Kb psMem=566564Kb
msMem=2451Kb psMem=566636Kb
msMem=2451Kb psMem=566676Kb

There are three stages.
Stage 1: Allocate big chunks of memory in Go.
Stage 2: Start filling in stuff in these big chunks of memory.
Stage 3: Reallocate new slices that are much smaller and call debug.FreeOSMemory.

We report two values every second. The first value is from memstats. It is what Go thinks it is
using. The second value is from ps. It is what the OS thinks we are using.

In Stage 1, we see that msMem >>> psMem. This is most likely because the OS is lazy in its memory
allocation. Until we really touch and use the memory, it does not really give you the memory.

In Stage 2, we see that msMem ~ psMem. This is because we start using the memory and the OS has
to indeed give us the memory.

In Stage 3, we see the problem we face. It is when msMem <<< psMem. To be sure I stare at ps
directly. Indeed, we are using ~567M and this is RSS (resident memory). It is not clear to me why
the garbage collector did not return the memory back to the OS.

USER               PID  %CPU %MEM      VSZ    RSS   TT  STAT STARTED      TIME COMMAND
jchiu            60333   1.3  6.8 556649020 566724 s000  S+    9:42AM   0:09.24 ./rdbdemo -dir /Users/jchiu/rdbdemo
*/
func allocDemo2() {
	n := 50000
	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = make([]byte, 10000)
	}
	fmt.Println("Allocated mem")
	time.Sleep(5 * time.Second)
	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = 55
		}
	}
	fmt.Println("Filled mem")
	time.Sleep(5 * time.Second)
	fmt.Println("Resizing slices")

	var sum int
	for {
		for i := 0; i < n; i++ {
			for _, b := range a[i] {
				// Explicitly iterate over each a[i] so that the mem doesn't get paged to disk and ps
				// will include that.
				sum += int(b)
			}
			a[i] = make([]byte, 10)
		}
		// We find that the memory remains allocated, as expected, and memstats and ps both reflect that.
		debug.FreeOSMemory()
		time.Sleep(time.Second)
	}
	fmt.Printf("Size sum %d\n", sum)
}

// We repeatedly create iterators. However, it seems that there is no memory leak here either
// as long as you remember to destroy the iterator.
func iterDemo(st *store.Store) {
	for i := 0; i < 100000; i++ {
		key := []byte(fmt.Sprintf("key%09x", i))
		val := bytes.Repeat([]byte("v"), 5000*5+9)
		valSuffix := val[5000*5:]
		x.AssertTrue(9 == copy(valSuffix, []byte(fmt.Sprintf("%09x", i))))
		x.Check(st.SetOne(key, val))
	}

	var sum int
	for i := 0; i < 10000; i++ {
		if (i % 1000) == 0 {
			fmt.Printf("Iter i\n", i)
		}
		it := st.NewIterator()
		for ; it.Valid(); it.Next() {
			sum += len(it.Value().Data())
		}
		it.Close()
	}
	fmt.Printf("Size sum %d\n", sum)
}

func main() {
	x.Init()

	//	go putDemo()
	//	go getDemo()
	//  go writeBatchDemo()
	//	go allocDemo()
	go allocDemo2()
	//	go iterDemo()

	var ms runtime.MemStats
	for {
		runtime.ReadMemStats(&ms)
		msKbs := ms.HeapAlloc / 1024
		psKbs := ps()
		fmt.Printf("msMem=%dKb psMem=%dKb\n", msKbs, psKbs)
		time.Sleep(time.Second)
	}
}
