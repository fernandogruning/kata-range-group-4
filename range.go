package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	r string
}

type RealRange struct {
	Start          int64
	StartInclusive bool
	End            int64
	EndInclusive   bool
}

func (r Range) ToRealRange() RealRange {
	splitted := strings.Split(r.r, ",")

	firstPart := splitted[0]
	endPart := splitted[1]

	startInclusive := false
	start := int64(0)
	end := int64(0)
	endInclusive := false

	firstRunes := []rune(firstPart)
	startInclusive = string(firstRunes[0:1]) == "["
	// (232
	s, _ := strconv.Atoi(string(firstRunes[1:len(firstPart)]))
	start = int64(s)

	endRunes := []rune(endPart)
	endInclusive = string(endRunes[len(endPart)-1:len(endPart)]) == "]"
	e, _ := strconv.Atoi(string(endRunes[0 : len(endPart)-1]))
	end = int64(e)

	return RealRange{start, startInclusive, end, endInclusive}
}

func main() {
	fmt.Println("Hey")
}
