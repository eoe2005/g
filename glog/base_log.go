package glog

import (
	"log"
	"os"
	"time"
)

const (
	SPLIT_FILESIZE = 1
	SPLIT_DAY      = 2
	SPLIT_HOUR     = 3
)

type Glog struct {
	maxsize   int64
	lastdata  int64
	fd        *os.File
	fileName  string
	splittype int
}

func new(filename string, ftype int, maxsize int64) *Glog {
	ret := &Glog{
		fileName:  filename,
		maxsize:   maxsize,
		splittype: ftype,
	}
	ret.reset()
	return ret
}
func (l *Glog) reset() {
	if l.fd != nil {
		l.fd.Close()
		rfilename := l.fileName + "."
		switch l.splittype {
		case SPLIT_FILESIZE:
			l.lastdata = 0
			rfilename += time.Now().Format("2006010215")
		case SPLIT_HOUR:
			rfilename += time.Now().Format("2006010215")
		default:
			rfilename += time.Now().Format("20060102")
		}
		os.Rename(l.fileName, rfilename)
	}
	fd, err := os.OpenFile(l.fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	l.fd = fd
	ctime := time.Now()
	switch l.splittype {
	case SPLIT_FILESIZE:
		l.lastdata = 0
	case SPLIT_HOUR:
		y, m, d := ctime.Date()
		l.lastdata = int64(y*1000000 + int(m)*10000 + d*100 + ctime.Hour())
	default:
		y, m, d := ctime.Date()
		l.lastdata = int64(y*10000 + int(m)*100 + d)
	}
}
func NewGlogSplitDay(filename string) *Glog {
	return new(filename, SPLIT_DAY, 0)
}
func NewGlogFileSize(filename string, maxsize int64) *Glog {
	return new(filename, SPLIT_FILESIZE, maxsize)
}
func NewGlogSplitHour(filename string) *Glog {
	return new(filename, SPLIT_HOUR, 0)
}

func (l *Glog) NewLog(prefix string, flag int) *log.Logger {
	return log.New(l, prefix, flag)
}
func (l *Glog) Write(p []byte) (n int, err error) {
	switch l.splittype {
	case SPLIT_FILESIZE:
		if l.lastdata >= l.maxsize {
			l.reset()
		}
	case SPLIT_HOUR:
		ctime := time.Now()
		y, m, d := ctime.Date()
		ct := int64(y*1000000 + int(m)*10000 + d*100 + ctime.Hour())
		if ct >= l.lastdata {
			l.reset()
		}
	default:
		ctime := time.Now()
		y, m, d := ctime.Date()
		ct := int64(y*10000 + int(m)*100 + d)
		if ct >= l.lastdata {
			l.reset()
		}
	}

	n, err = l.fd.Write(p)
	if err != nil {
		return
	}
	l.lastdata += int64(n)
	return
}
