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
		runtime.GC()
		debug.FreeOSMemory()
		time.Sleep(time.Second)
	}
	fmt.Printf("Size sum %d\n", sum)
}

/*
Allocated mem
msMem=514159Kb psMem=30628Kb
msMem=514200Kb psMem=30648Kb
Filled mem
msMem=514241Kb psMem=563864Kb
msMem=514282Kb psMem=563864Kb
msMem=514322Kb psMem=563868Kb
Shrinking all slices
msMem=487Kb psMem=563976Kb
msMem=528Kb psMem=563976Kb
msMem=568Kb psMem=563976Kb
msMem=608Kb psMem=563976Kb

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
directly. Indeed, we are using 500+ M and this is RSS (resident memory). It seems that the GC is
trying to save some work by not returning memory to the OS unless it is necessary.

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
	time.Sleep(3 * time.Second)
	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = 55
		}
	}
	fmt.Println("Filled mem")
	time.Sleep(3 * time.Second)
	fmt.Println("Shrinking all slices")
	for i := 0; i < n; i++ {
		a[i] = make([]byte, 2)
	}
	runtime.GC()
	debug.FreeOSMemory()
}

/*
Similar to allocDemo2. However, for the last stage, instead of shrinking all slices, we
alternate between small and large slices. We see that in this case, ps says we use the most
amount of memory.

The reason I think is that before the GC can give us back the space, there were new allocations
requested for large slices. Hence, the OS has to give us more memory. This might explain why
if the dgraph server says the mem limit is 1G, then it ends up using 2G. Whereas if it requests
for 2G, it ends up using 4G.

msMem=663Kb psMem=13632Kb
Allocated mem
msMem=514157Kb psMem=30624Kb
msMem=514197Kb psMem=30628Kb
Filled mem
msMem=514237Kb psMem=563844Kb
msMem=514277Kb psMem=563856Kb
msMem=514318Kb psMem=563856Kb
Varying slice sizes
msMem=535770Kb psMem=564924Kb
msMem=155477Kb psMem=730632Kb
msMem=197695Kb psMem=730632Kb
msMem=165465Kb psMem=730636Kb
*/
func allocDemo3() {
	n := 50000
	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = make([]byte, 10000)
	}
	fmt.Println("Allocated mem")
	time.Sleep(3 * time.Second)
	for i := 0; i < n; i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = 55
		}
	}
	fmt.Println("Filled mem")
	time.Sleep(3 * time.Second)

	fmt.Println("Varying slice sizes")
	var sum int
	for {
		for i := 0; i < n; i++ {
			for _, b := range a[i] {
				// Explicitly iterate over each a[i] so that the mem doesn't get paged to disk and ps
				// will include that.
				sum += int(b)
			}
			if (i % 10) < 7 {
				a[i] = make([]byte, 2)
			} else {
				a[i] = make([]byte, 10000)
			}
		}
		runtime.GC()
		debug.FreeOSMemory()
	}
	fmt.Printf("Size sum %d\n", sum)
}

func main() {
	x.Init()

	//	go putDemo()
	//	go getDemo()
	//go writeBatchDemo()
	//	go iterDemo()
	//	go allocDemo()
	go allocDemo2()
	//	go allocDemo3()

	var ms runtime.MemStats
	for {
		runtime.ReadMemStats(&ms)
		msKbs := ms.HeapAlloc / 1024
		psKbs := ps()
		fmt.Printf("msMem=%dKb psMem=%dKb\n", msKbs, psKbs)
		time.Sleep(time.Second)
	}
}
