package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 찾은 라인 정보
type LineInfo struct {
	lineNo int
	line   strings.Builder
}

// 파일 내 라인 정보
type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.3 word filepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2]
	findInfos := []FindInfo{}
	for _, path := range files {
		// 파일 찾기
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("--------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("--------------------------------")
		fmt.Println()

	}

}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	filelist, err := filepath.Glob(path) // 실행 인수 가져오기
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}


	for findInfo := append(findInfos, findINfo)
	recvCnt++
	if recvCnt == cnt {

	}

}
