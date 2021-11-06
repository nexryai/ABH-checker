package main

import (
	"flag"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	var url = flag.Arg(0)

	var abh float64

	defaultUmask := syscall.Umask(0)
	syscall.Umask(defaultUmask)

	abh_start := time.Now()
	exec.Command("wget", " -p -np -P /dev/null", " http://abehiroshi.la.coocan.jp/").Run()
	abh_speed := time.Since(abh_start).Milliseconds()
	fmt.Printf("経過: %vms\n", time.Since(abh_start).Milliseconds())

	url_start := time.Now()
	exec.Command("wget", " -p -np -P /dev/null", " ", url).Run()
	url_speed := time.Since(url_start).Milliseconds()
	fmt.Printf("経過: %vms\n", time.Since(url_start).Milliseconds())


	fmt.Println("結果 阿部寛のホームページ:", abh_speed, " 指定されたURL:", url_speed)
	abh = float64(url_speed) / float64(abh_speed)
	fmt.Println(url, "は", abh, "ABHです")

}
