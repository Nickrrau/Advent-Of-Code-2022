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
	fmt.Println("AoC Day 7 Part 2")

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
	result := getBestCandidatesForDeletion([]int{}, 30000000-(70000000-rootFolder.CalcSize()), rootFolder)

	fmt.Println(result[0])
}

func getBestCandidatesForDeletion(candidates []int, targetFree int, root Folder) []int {
	for _, f := range root.folders {
		if targetFree-f.CalcSize() <= 0 {
			candidates = insertValue(candidates, f.CalcSize())
		}

		candidates = getBestCandidatesForDeletion(candidates, targetFree, *f)
	}
	return candidates
}

func insertValue(slice []int, v int) []int {
	if v == 0 {
		return slice
	}
	for i, s := range slice {
		if v < s {
			return append(slice[0:i], append([]int{v}, slice[i:]...)...)
		}
	}
	return append(slice, v)
}