package helpers

import "io/ioutil"

//GetFileContentsAsStr ...
func GetFileContentsAsStr(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	CheckErr(err)
	return string(content)
}

//utils "github.com/nagamocha3000/aoc2019/src/helpers"
