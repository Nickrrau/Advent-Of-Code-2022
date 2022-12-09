package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Folder struct {
	parentFolder *Folder
	name         string
	files        []int
	folders      map[string]*Folder
}

func (f *Folder) CalcSize() int {
	total := 0
	for _, v := range f.files {
		total = total + v
	}
	for _, v := range f.folders {
		total = total + v.CalcSize()
	}
	return total
}

func main() {
	fmt.Println("AoC Day 7 Part 1")

	if len(os.Args) < 2 {
		fmt.Println("Need to provide input")
		return
	}

	input := os.Args[1]
	processInput(input)
}

func processInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rootFolder := Folder{name: "/", folders: map[string]*Folder{}, files: []int{}}
	curFolder := &rootFolder

	listed := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '$' {
			listed = false
			if line[2:4] == "cd" {
				if line[5:] == "/" {
					curFolder = &rootFolder
				} else if line[5:] == ".." {
					curFolder = curFolder.parentFolder
				} else {
					newFolder := &Folder{name: line[5:], parentFolder: curFolder, folders: map[string]*Folder{}, files: []int{}}
					curFolder.folders[line[5:]] = newFolder
					curFolder = newFolder
				}
			} else if line[2:4] == "ls" {
				listed = true
			}
		} else {
			if listed {
				if line[:3] != "dir" {
					size, _ := strconv.Atoi(strings.Split(line, " ")[0])
					curFolder.files = append(curFolder.files, size)
				}
			}
		}
	}

	fmt.Println(getBestCandidatesForDeletion(rootFolder))
}

func getBestCandidatesForDeletion(root Folder) int {
	largest := 0
	for _, f := range root.folders {
		if f.CalcSize() < 100000 {
			largest = largest + f.CalcSize()
		}
		largest = largest + getBestCandidatesForDeletion(*f)
	}
	return largest
}
