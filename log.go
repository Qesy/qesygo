package qesygo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// DailyFileWriter 自定义带互斥锁的按天写文件器
type DailyFileWriter struct {
	mu       sync.Mutex
	basePath string
	file     *os.File
	day      int
}

// Write 实现 io.Writer 接口，确保并发安全并自动检查日期跨天
func (w *DailyFileWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 检查是否跨天，如果跨天则重新创建新日期的文件
	now := time.Now()
	if now.Day() != w.day {
		if err := w.rotate(now); err != nil {
			return 0, err
		}
	}

	return w.file.Write(p)
}

// rotate 执行文件切换操作
func (w *DailyFileWriter) rotate(now time.Time) error {
	if w.file != nil {
		_ = w.file.Close()
	}

	// 确保目标目录存在
	dir := filepath.Dir(w.basePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 生成带日期的文件名：./static/log/error-2026-08-12.log
	filePath := fmt.Sprintf("%s-%s.log", w.basePath, now.Format("2006-01-02"))
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	w.file = f
	w.day = now.Day()
	return nil
}

func Log(Path string) {
	now := time.Now()
	writer := &DailyFileWriter{
		basePath: Path,
	}

	// 初始化创建当天的文件
	if err := writer.rotate(now); err != nil {
		log.Printf("Failed to initialize log file: %v\n", err)
		return
	}

	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	log.SetOutput(writer)
	fmt.Println("Log Start Success !")
}
