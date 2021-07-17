package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	flag.Parse()
	var url = flag.Arg(0)

	var abh float64

	defaultUmask := syscall.Umask(0)
	os.Mkdir("/tmp/abh", 0777)
	syscall.Umask(defaultUmask)
	os.Mkdir("/tmp/abh/abh", 0777)
	os.Mkdir("/tmp/abh/usersite", 0777)

	abh_start := time.Now()
	exec.Command("wget", " -p -np -P /tmp/abh/abh/", " http://abehiroshi.la.coocan.jp/").Run()
	abh_speed := time.Since(abh_start).Milliseconds()
	fmt.Printf("経過: %vms\n", time.Since(abh_start).Milliseconds())

	url_start := time.Now()
	exec.Command("wget", " -p -np -P /tmp/abh/usersite/", " ", url).Run()
	url_speed := time.Since(url_start).Milliseconds()
	fmt.Printf("経過: %vms\n", time.Since(url_start).Milliseconds())

	os.RemoveAll("/tmp/abh")

	fmt.Println("結果 阿部寛のホームページ:", abh_speed, " 指定されたURL:", url_speed)
	abh = float64(url_speed) / float64(abh_speed)
	fmt.Println(url, "は", abh, "ABHです")

}
