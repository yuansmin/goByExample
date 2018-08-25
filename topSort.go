// 拓扑排序
// 根据课程的前置关系，排出合理的课程列表
// prereqs 为关系拓扑，key，value 分别是课程、该课程的前置依赖
// 算法关键在于找到当前课程的依赖、依赖的依赖，并记录已出现的课程，以此递归查找

package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for _, course := range topSort(&prereqs) {
		fmt.Println(course)
	}
}

func topSort(top *map[string][]string) []string {
	seen := make(map[string]bool)
	sortedList := []string{}
	var addCourse func(course string)
	addCourse = func(course string) {
		if !seen[course] {
			seen[course] = true
			for _, pCource := range (*top)[course] {
				addCourse(pCource)
			}
			sortedList = append(sortedList, course)
		}
	}
	for course, _ := range *top {
		addCourse(course)
	}
	return sortedList
}
