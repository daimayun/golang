package random

import (
	"fmt"
	"github.com/daimayun/golang/date"
	"os"
	"sync/atomic"
	"time"
)

var num int64

// NumberString 生成24位数字字符串 [前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号]
func NumberString() string {
	t := time.Now()
	s := t.Format(date.TimeLayoutYMDHIS)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
