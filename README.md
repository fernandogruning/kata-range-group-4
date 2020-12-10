# Members

- 1077014 - Alan Hosking Rivera
- 1084757 - Raúl Luna Pérez
- 1075303 - Jeremy Martínez Matos
- 1076825 - Fernando Gruning Moore

# Description

This repo hosts our version of the Kata Range, it was developed with Go on Visual Studio and using the TDD approach to produce a reliable piece of code. We currently have 20 test cases.


# Requirements:

- REQ-0: Range has to be built from a input string “[n,n]”

- REQ-1: Range has to have a start point and an endpoint, either of which can be inclusive (get included in counting) or exclusive.

- REQ-2: Points have to be valid integers only.

- REQ-3: Range should have a GetAllPoints method for displaying all numbers contained in the range.

- REQ-4: Range should have a ContainsTange method for evaluating if a range contains another in it.

- REQ-5: Range should have an EndPoints method for displaying the first and the last included numbers.

- REQ-6: Range should have an OverLapsRange method for evaluating if a range overlaps another.

- REQ-7: Range should have an Equals method for evaluating if two ranges are equal.

# Acceptance Criteria:

- CRI-1-1: The start point cannot be greater than the endpoint.

- CRI-1-2: The endpoint cannot be lesser than the start point.

- CRI-1-3: The Range can be all-inclusive.

- CRI-1-4: The range can be all-exclusive.

- CRI-2-1: The start point should be a valid integer.

- CRI-2-2: The endpoint should be a valid integer.

- CRI-3-1: GetAllPoints should have no parameters and return a string containing every point inside a range separated by a comma.

- CRI-4-1: ContainsRange should accept  another valid Range and return a string displaying whether or not the Range contains the other.

- CRI-5-1: Endpoints should have no parameters and return a string displaying the first included point and the last included point.

- CRI-6-1: OverlapsRange should accept another valid Range and return a string displaying whether or not the Range overlaps with the other.

- CRI-7-1: Equals should accept another valid Range and return a string displaying whether or not the two Ranges are equal.


# Test cases:

- SCE-0-1-1: `Range = “(2,5]” ToRealRange = (2,5]`

- SCE-0-1-2: `Range = “(3,6]” ToRealRange = (3,6]`

- SCE-3-1-1: `Range = [2,6) GetAllPoints = {2,3,4,5}`

- SCE-3-1-2: `Range = [1,10) GetAllPoints = {1,2,3,4,5,6,7, 8,9}`

- SCE-4-1-1: `Range = [2,5) ContainsRange = [2,5) doesn’t contain [7,10)`

- SCE-4-1-2: `Range = [2,5) ContainsRange = [2,5) doesn’t contain [3,10)`

- SCE-4-1-3: `Range = [2,5) ContainsRange = [2,5) contains [3,5)`

- SCE-4-1-4: `Range = [3,5] ContainsRange = [3,5] contains [3,5)`

- SCE-5-1-1: `Range = [2,6) EndPoints = {2,5}`

- SCE-5-1-2: `Range = [2,6] EndPoints = {2,6}`

- SCE-5-1-3: `Range = (2,6) EndPoints = {3,5}`

- SCE-5-1-4: `Range = (2,6] EndPoints = {3,6}`

- SCE-6-1-1: `Range = [2,5) OverlapsRange [2,5) doesn’t overlap with [7,10)`

- SCE-6-1-2: `Range = [2,10) OverlapsRange [2,10) overlaps with [3,5)`

- SCE-6-1-3: `Range =[2,5) OverlapsRange [2,5) overlaps with [3,10)`

- SCE-6-1-4: `Range =[3,5) OverlapsRange [3,5) overlaps with [2,10)`

- SCE 7-1-1: `Range =[2,10) Equals [2,10) neq [3,5)`

- SCE-7-1-2: `Range =[3,5) Equals [3,5) Equals [3,5) equals [3,5)`

- SCE-7-1-3: `Range =[2,5) Equals [2,5) neq [3,10)`

- SCE-7-1-4: `Range =[3,5) Equals [3,5) neq [2,10)`
