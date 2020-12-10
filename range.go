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

func (r Range) GetAllPoints() string {

	RealRange := r.ToRealRange()
	strRange := ""

	start := RealRange.Start
	if !(RealRange.StartInclusive) {
		start = RealRange.Start + 1
	}
	end := RealRange.End
	if RealRange.EndInclusive {
		end = RealRange.End + 1
	}

	for i := start; i < end; i++ {
		if !(i+1 == end) {
			strRange = fmt.Sprintf("%s%d,", strRange, i)
		} else {
			strRange = fmt.Sprintf("%s%d", strRange, i)
		}
	}
	return fmt.Sprintf("{%s}", strRange)
}

func (r Range) ContainsRange(cr Range) string {
	RealRange := r.ToRealRange()
	compareRealRange := cr.ToRealRange()

	startBracket := "["
	start := RealRange.Start
	if !(RealRange.StartInclusive) {
		startBracket = "("
		start = RealRange.Start + 1
	}

	endBracket := ")"
	end := RealRange.End
	if RealRange.EndInclusive {
		endBracket = "]"
		end = RealRange.End + 1
	}
	compareStartBracket := "["
	compareStart := compareRealRange.Start
	if !(compareRealRange.StartInclusive) {
		compareStartBracket = "("
		compareStart = compareRealRange.Start + 1
	}

	compareEndBracket := ")"
	compareEnd := compareRealRange.End
	if compareRealRange.EndInclusive {
		compareEndBracket = "]"
		compareEnd = compareRealRange.End + 1
	}

	if compareStart < start || compareEnd > end {
		return fmt.Sprintf("%v%d,%d%v doesn't contain %v%d,%d%v", startBracket, RealRange.Start, RealRange.End, endBracket, compareStartBracket, compareRealRange.Start, compareRealRange.End, compareEndBracket)
	}
	return fmt.Sprintf("%v%d,%d%v contains %v%d,%d%v", startBracket, RealRange.Start, RealRange.End, endBracket, compareStartBracket, compareRealRange.Start, compareRealRange.End, compareEndBracket)

}

func (r Range) EndPoints() string {
	rr := r.ToRealRange()

	firstEndpoint := rr.Start
	if !(rr.StartInclusive) {
		firstEndpoint = rr.Start + 1
	}

	lastEndpoint := rr.End
	if !(rr.EndInclusive) {
		lastEndpoint = rr.End - 1
	}

	return fmt.Sprintf("{%d,%d}", firstEndpoint, lastEndpoint)
}

func (r Range) OverlapsRange(cr Range) string {
	RealRange := r.ToRealRange()
	compareRealRange := cr.ToRealRange()

	startBracket := "["
	start := RealRange.Start
	if !(RealRange.StartInclusive) {
		startBracket = "("
		start = RealRange.Start + 1
	}

	endBracket := ")"
	end := RealRange.End
	if RealRange.EndInclusive {
		endBracket = "]"
		end = RealRange.End + 1
	}
	compareStartBracket := "["
	compareStart := compareRealRange.Start
	if !(compareRealRange.StartInclusive) {
		compareStartBracket = "("
		compareStart = compareRealRange.Start + 1
	}

	compareEndBracket := ")"
	compareEnd := compareRealRange.End
	if compareRealRange.EndInclusive {
		compareEndBracket = "]"
		compareEnd = compareRealRange.End + 1
	}

	containsAtLeastOne := false
	for i := start; i < end; i++ {
		for j := compareStart; j < compareEnd; j++ {
			if i == j {
				containsAtLeastOne = true
				break
			}
		}
	}

	if !containsAtLeastOne {
		return fmt.Sprintf("%v%d,%d%v doesn't overlap with %v%d,%d%v", startBracket, RealRange.Start, RealRange.End, endBracket, compareStartBracket, compareRealRange.Start, compareRealRange.End, compareEndBracket)
	}

	return fmt.Sprintf("%v%d,%d%v overlaps with %v%d,%d%v", startBracket, RealRange.Start, RealRange.End, endBracket, compareStartBracket, compareRealRange.Start, compareRealRange.End, compareEndBracket)

}

func main() {
	fmt.Println("Hey")
}
