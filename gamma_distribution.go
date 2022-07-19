package main

/*
#cgo CFLAGS: -I./include
// C 标志io头文件，你也可以使用里面提供的函数
#cgo LDFLAGS: -L./lib -lgamma_lib -Wl,-rpath,lib

#include "library.h"

*/
import "C" // 切勿换行再写这个

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"text/tabwriter"
	"time"
)

type ResponseData struct {
	GammaRate float64 `json:"gamma"`
}
type BidObj struct {
	Id    string  `json:"id"`
	Impid string  `json:"impid,omitempty"`
	Price float64 `json:"price"`
}

var gamma_a float64
var gamma_b float64

func main() {
	timeStart := time.Now().Unix()
	fmt.Println("time start:", timeStart)
	CORE_NUM := runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(CORE_NUM * 4)
	fs := flag.NewFlagSet("main", flag.ExitOnError)
	var (
		s_a    = fs.String("gamma_a", "", "gamma a value")
		s_b    = fs.String("gamma_b", "", "gamma b value ")
		s_port = fs.String("port", "", "http port")
	)
	fs.Usage = UsageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])
	//fmt.Println("flag:", fs)
	fmt.Println("gamma a:", *s_a)
	fmt.Println("gamma b:", *s_b)
	fmt.Println("port:", *s_port)
	gamma_a, _ = strconv.ParseFloat(*s_a, 64)

	gamma_b, _ = strconv.ParseFloat(*s_b, 64)

	HandleKillSignal()

	http.HandleFunc("/gamma_random", DoGammaRandom)

	if err := http.ListenAndServe("0.0.0.0:"+*s_port, nil); err != nil {
		fmt.Println("start http server failed: ", err.Error())
		return
	}
	for i := 0; i < 100; i++ {
		fmt.Println(C.gamma_random(2.0, 2.0))
		fmt.Println("i:", i)
	}

	print("hello world")
}

func DoGammaRandom(res http.ResponseWriter, req *http.Request) {
	var res_data ResponseData
	temp := C.gamma_random(C.double(gamma_a), C.double(gamma_b))
	res_data.GammaRate = float64(temp)
	res.Header().Set("Content-Type", "application/json")
	if b, err := json.Marshal(res_data); err == nil {
		fmt.Println(string(b))
		res.WriteHeader(http.StatusOK)
		res.Write(b)
	} else {
		res.WriteHeader(http.StatusNoContent)

	}

}

func UsageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func HandleKillSignal() {
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-sigchan
		fmt.Println("get shutdown signal.")
		// 关闭HTTP服务
		//manners.Close()
		// 停止extractor
		//extractor.Stop()
		fmt.Println("close http server.")
		os.Exit(0)
	}()
}
