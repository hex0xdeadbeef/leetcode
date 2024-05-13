package middle

// https://leetcode.com/problems/find-duplicate-file-in-system/description/

import (
	"strings"
)

func FindDuplicate(dirs []string) [][]string {
	var (
		seenContents           = map[string][]string{}
		dirRoot, name, content string
	)

	for _, dir := range dirs {
		info := strings.Fields(dir)
		dirRoot = info[0] + "/"

		for _, file := range info[1:] {
			name, content = getFileAndContent(file)

			if _, ok := seenContents[content]; !ok {
				seenContents[content] = make([]string, 0, 8)
				seenContents[content] = append(seenContents[content], dirRoot+name)
				continue
			}
			seenContents[content] = append(seenContents[content], dirRoot+name)

		}
	}

	result := make([][]string, 0, len(seenContents))
	for _, group := range seenContents {
		if len(group) >= 2 {
			result = append(result, group)
		}
	}

	return result
}

func getFileAndContent(fileStr string) (string, string) {
	var (
		bracketIndex = strings.Index(fileStr, "(")
	)
	return fileStr[:bracketIndex], fileStr[bracketIndex+1 : len(fileStr)-1]
}
