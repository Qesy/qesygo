package qesygo

import "os"

type FileLib struct {
}

func NewFileLib() *FileLib {
	return &FileLib{}
}

func (q *FileLib) Write(filePath, content string) { // 以追加模式打开文件（如果文件不存在则创建）
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}
}
