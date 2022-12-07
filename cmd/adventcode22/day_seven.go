package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Dirent struct {
	TotalSize   int
	Name        string
	IsDirectory bool
	Parent      *Dirent
	Content     []*Dirent
}

func CreateDirectory(content string, parent *Dirent) *Dirent {
	dir := strings.Split(content, " ")
	directory := new(Dirent)
	directory.Name = dir[1]
	directory.IsDirectory = true
	directory.Parent = parent

	return directory
}

func CreateFile(content string, parent *Dirent) *Dirent {
	entries := strings.Split(content, " ")
	fileSize, _ := strconv.Atoi(entries[0])
	file := new(Dirent)
	file.TotalSize = fileSize
	file.Name = entries[1]
	file.IsDirectory = false
	file.Parent = parent

	return file
}

func HandleCdCommand(input string, currentDirectory *Dirent) *Dirent {
	commands := strings.Split(input, " ")
	for _, dir := range currentDirectory.Content {
		if dir.Name == commands[2] {
			return dir
		}
	}

	if commands[2] == currentDirectory.Name {
		return currentDirectory
	}

	if commands[2] == ".." {
		return currentDirectory.Parent
	}

	return nil
}

func AddDirentToDirectory(input string, currentDirectory *Dirent) {
	if strings.Contains(input, "dir") {
		currentDirectory.Content = append(currentDirectory.Content, CreateDirectory(input, currentDirectory))
	} else {
		file := CreateFile(input, currentDirectory)
		currentDirectory.TotalSize += file.TotalSize
		currentDirectory.Content = append(currentDirectory.Content, file)
	}
}

func CalculateSizeForAllDirectories(root *Dirent) map[string]int {
	result := make(map[string]int)
	var key string
	if root.Parent != nil {
		delimiter := "/"
		if root.Parent.Name == "/" {
			delimiter = ""
		}
		key = root.Parent.Name + delimiter + root.Name
	} else {
		key = root.Name
	}

	result[key] = root.TotalSize

	for _, dir := range root.Content {
		if dir.IsDirectory {
			subResult := CalculateSizeForAllDirectories(dir)
			for k, v := range subResult {
				result[k] += v
			}
			result[key] += dir.TotalSize
		}
	}

	if root.Parent != nil {
		root.Parent.TotalSize += root.TotalSize
	}

	return result
}

func DaySeven() {
	openFile := readFileByLines("inputs/d7.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return
	}

	root := CreateDirectory("dir /", nil)
	current := root
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()
		if strings.Contains(input, "$ cd") {
			current = HandleCdCommand(input, current)
			continue
		} else {
			AddDirentToDirectory(input, current)
		}
	}

	totalSizes := CalculateSizeForAllDirectories(root)

	currentDiskSpace := 70000000 - totalSizes["/"]
	missingSpace := 30000000 - currentDiskSpace
	candidatesForRemoval := []struct {
		key   string
		value int
	}{}
	totalSizeA := 0
	for key, size := range totalSizes {
		if size < 100000 {
			totalSizeA += size
		}
		if size >= missingSpace {
			candidatesForRemoval = append(candidatesForRemoval, struct {
				key   string
				value int
			}{key, size})
		}
	}

	fmt.Println(totalSizeA)

	sort.SliceStable(candidatesForRemoval, func(i, j int) bool {
		return candidatesForRemoval[i].value < candidatesForRemoval[j].value
	})
	fmt.Println(candidatesForRemoval[0].value)

	openFile.File.Close()
}
