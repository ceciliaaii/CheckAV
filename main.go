package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 指定 avList.txt 文件的路径
	avListPath := "E:\\GolandProject\\src\\avList.txt"
	fmt.Println("avList.txt 文件的路径:", avListPath)

	// 打开 avList.txt 文件
	avListFile, err := os.Open(avListPath)
	if err != nil {
		fmt.Println("无法打开 avList.txt 文件:", err)
		return
	}
	defer avListFile.Close()

	// 读取 avList.txt 文件内容，存储到 map 中
	avList := make(map[string]string)
	scanner := bufio.NewScanner(avListFile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, `": "`)
		if len(parts) == 2 {
			avList[strings.Trim(parts[0], `" `)] = strings.Trim(parts[1], `" `)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取 avList.txt 文件失败:", err)
		return
	}

	// 执行 tasklist 命令
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行 tasklist 命令失败:", err)
		return
	}

	// 将 tasklist 命令的结果存储到 tasklist.txt 文件中
	outputFile, err := os.Create("tasklist.txt")
	if err != nil {
		fmt.Println("无法创建 tasklist.txt 文件:", err)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	defer outputWriter.Flush()

	outputWriter.WriteString(string(output))

	// 检查杀软是否存在
	found := false
	for process, description := range avList {
		if strings.Contains(string(output), process) {
			fmt.Printf("发现杀软进程 %s 存在：%s\n", process, description)
			found = true
		}
	}

	if !found {
		fmt.Println("未检测到杀软存在")
	}
}
