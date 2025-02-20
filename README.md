# Go Slice Modification within Range Loop
This example demonstrates a common error when modifying a slice within a range loop in Go.  Incorrectly handling index adjustments after removing elements can lead to unexpected behavior, skipping elements or causing out-of-bounds errors. The solution illustrates a safer approach using a different loop structure or creating a new slice.

## Bug
The `bug.go` file contains a function that attempts to remove elements from a slice while iterating over it using a `for...range` loop. This approach is problematic because removing elements shifts the indices, leading to potential issues.  The loop might skip elements or even cause a runtime panic due to index-out-of-bounds errors.

## Solution
The `bugSolution.go` file presents two alternatives:  First, using a `for` loop with manual index tracking provides more control. The second utilizes `append` to create a new slice containing only the desired elements.