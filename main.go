package main

import (
	"flag"
	"log"
	"os"
	"sync/atomic"
	"time"

	"fmt"

	"github.com/nashera/merge-data/civic"
	"github.com/nashera/merge-data/cosmic"
)

const (
	// Version 版本号
	Version = "v1.0"
)

var (
	cosmicTsv string
	civicTsv  string
)

// CosmicVar variant in cosmic database
var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func init() {
	flag.StringVar(&cosmicTsv, "cosmic", "/cgdata/zhangxi/project/oncogene/cosmic/CosmicMutantExport.tsv.gz", "the path to cosmicTSV")
	flag.StringVar(&civicTsv, "civic", "/cgdata/zhangxi/project/oncogene/civic/variants.with_hgvs.tsv", "the path to civicTSV")
	flag.Parse()
}

func main() {
	cosmicVarList := cosmic.ReadCosmic(cosmicTsv)
	// for _, variant := range cosmicVarList {
	// 	fmt.Println(variant)
	// }

	civicVarList := civic.ReadCivic(civicTsv)
	fp, err := os.OpenFile("/cgdata/zhangxi/project/oncogene/merge-data/civicVsComic.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panic(err)
	}
	ch := make(chan *civic.VarCivicComplete)
	// defer fp.Close()
	// defer close(ch)
	// defer func(ch chan civic.VarCivicComplete, err )

	// var varList []*civic.VarCivicComplete
	go civic.PrintVariantList(ch, fp)
	for _, civicVar := range civicVarList {
		ch <- civic.SearchCivicVar(civicVar, cosmicVarList)
	}
	fp.Sync()
	// civic.PrintVariantList(varList)

	// defer ants.Release()

	// runTimes := 1000

	// // Use the common pool.
	// var wg sync.WaitGroup
	// syncCalculateSum := func() {
	// 	demoFunc()
	// 	wg.Done()
	// }
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	_ = ants.Submit(syncCalculateSum)
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", ants.Running())
	// fmt.Printf("finish all tasks.\n")

	// // Use the pool with a function,
	// // set 10 to the capacity of goroutine pool and 1 second for expired duration.
	// p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
	// 	myFunc(i)
	// 	wg.Done()
	// })
	// defer p.Release()
	// // Submit tasks one by one.
	// for i := 0; i < runTimes; i++ {
	// 	wg.Add(1)
	// 	_ = p.Invoke(int32(i))
	// }
	// wg.Wait()
	// fmt.Printf("running goroutines: %d\n", p.Running())
	// fmt.Printf("finish all tasks, result is %d\n", sum)
}
