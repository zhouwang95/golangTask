package main

import (
	"fmt"
	"strconv"
)

/*只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
 */
func onlyOnceNum(arr []int) int {
	mapNum := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if mapNum[arr[i]] == 1 {
			delete(mapNum, arr[i])
		} else {
			mapNum[arr[i]] = 1
		}
	}

	var onceNum int
	for key, _ := range mapNum {
		onceNum = key
	}
	return onceNum
}

// 回文数判断
func isHuiWen(num int) bool {
	var str string = strconv.Itoa(num)
	m1 := make([]string, len(str)/2)
	m2 := make([]string, len(str)/2)

	m := make([]string, len(str))

	for index, val := range str {
		m[index] = string(val)
	}

	m1 = m[:len(str)/2]

	var f int = len(str) % 2
	if f == 0 {
		m2 = m[len(str)/2:]
	} else {
		m2 = m[len(str)/2+1:]
	}

	for i := 0; i < len(str)/2; i++ {
		if m1[i] != m2[len(str)/2-1-i] {
			return false
		}
	}
	return true
}

// 有效的括号
func isValid(s string) bool {
	f := len(s) % 2
	if f == 1 {
		return false
	}

	p := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if p[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != p[s[i]] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 最长公共前缀
func longestCommonPrefix(arr []string) string {
	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[0] == "" {
				return ""
			}
			l := len(arr[j]) - 1
			if l == i && arr[j][i] == arr[0][i] {
				return arr[j]
			}
			if i < l && arr[j][i] != arr[0][i] {
				str := arr[j][:i]
				return str
			}
		}
	}
	return ""
}

/*
加一
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func addOneArr(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		n := arr[i] + 1
		if n > 9 {
			arr[i] = 0
		} else {
			arr[i] = n
			break
		}
	}
	return arr
}

/*
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func deleteValue(nums []int) []int {
	j := 0
	if len(nums) < 2 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			j++
		}
	}
	return nums
}

/*
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
//排序方法
func sortInterval(intervals [][]int) [][]int {
	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] < intervals[j][0] {
				tem := intervals[i]
				intervals[i] = intervals[j]
				intervals[j] = tem
			}
		}
	}
	return intervals
}

// 合并区间
func mergeInterval(intervals [][]int) [][]int {
	var intervalsSort [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] >= intervals[i][0] {
			arr := []int{intervals[i-1][0], intervals[i][1]}
			intervalsSort = append(intervalsSort, arr)
			if len(intervals) == i {
				return intervalsSort
			}
		} else {
			intervalsSort = append(intervalsSort, intervals[i-1])
			if (len(intervals) - 1) == i {
				intervalsSort = append(intervalsSort, intervals[i])
				return intervalsSort
			}
		}
	}
	return mergeInterval(intervalsSort)
}

// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func sumOfTwoNum(nums []int32, target int32) [][2]int32 {
	map0 := make(map[int32]int)
	for _, num := range nums {
		map0[num] = 1
	}

	var res [][2]int32
	for key, _ := range map0 {
		num1 := target - key
		_, err := map0[num1]
		if err != false {
			n := [2]int32{key, num1}
			res = append(res, n)
			delete(map0, num1)
		}
	}
	return res
}

func main() {
	nums := []int32{1, 8, 3, 4, 6, 15, 2, 7, 2}
	fmt.Printf("output value： %d\n", sumOfTwoNum(nums, 9))

	/*intervals := [][]int{{3, 5}, {1, 2}, {4, 8}, {10, 15}, {7, 11}}
	fmt.Printf("output value： %d\n", mergeInterval(sortInterval(intervals)))*/

	/*arr := []int{0, 7, 8, 9, 9, 9, 10}
	fmt.Printf("output value： %d\n", deleteValue(arr))
	*/
	/*arr := []int{0, 9, 9, 9, 9}
	fmt.Printf("output value： %d\n", addOneArr(arr))*/

	/*arr := []string{"qwer123", "qwertasd", "qwer"}
	fmt.Printf("output value： %s\n", longestCommonPrefix(arr))*/

	/*arr := []int{3, 7, 3, 5, 5, 6, 6, 8, 8, 9, 9, 0, 0}
	fmt.Printf("output value： %d\n", onlyOnceNum(arr))*/

	//fmt.Printf("是否为回文数字：", isHuiWen(1234543210))

	//fmt.Printf("是否为有效的括号：", isValid("()[]{}"))

}
