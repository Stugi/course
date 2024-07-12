package main

func main() {
	el3 := ListNode{Val: 7}
	el2 := ListNode{Val: 8, Next: &el3}
	el1 := ListNode{Val: 6, Next: &el2}

	el4 := ListNode{Val: 8}
	el5 := ListNode{Val: 6, Next: &el4}
	el6 := ListNode{Val: 7, Next: &el5}

	addTwoNumbers(&el1, &el6)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	result := head
	var increase int
	for l1 != nil || l2 != nil || increase != 0 {
		result.Next = new(ListNode)
		result = result.Next
		if l1 != nil && l2 != nil {
			if sum := l1.Val + l2.Val; sum >= 10 {
				result.Val = sum - 10 + increase
				increase = 1
			} else {
				if sum2 := sum + increase; sum2 >= 10 {
					result.Val = sum2 - 10
					increase = 1
				} else {
					result.Val = sum + increase
					increase = 0
				}
			}
			l1 = l1.Next
			l2 = l2.Next
			continue
		} else if l1 != nil {
			if sum3 := l1.Val + increase; sum3 >= 10 {
				result.Val = sum3 - 10
				increase = 1
			} else {
				result.Val = sum3
				increase = 0
			}
			l1 = l1.Next
			continue
		} else if l2 != nil {
			if sum4 := l2.Val + increase; sum4 >= 10 {
				result.Val = sum4 - 10
				increase = 1
			} else {
				result.Val = sum4
				increase = 0
			}
			l2 = l2.Next
			continue
		}
		result.Val = increase
		break
	}
	return head.Next
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	// ret := &ListNode{}
// 	// tmpRet := ret
// 	counter := 0
// 	res := 0
// 	for l1 != nil || l2 != nil {
// 		if l1 != nil || l2 != nil {
// 			res += int(math.Pow(10, float64(counter)))*l1.Val + int(math.Pow(10, float64(counter)))*l2.Val
// 			l1 = l1.Next
// 			l2 = l2.Next
// 			counter++
// 			continue
// 		}
// 		if l1 != nil {
// 			sum := calc(counter, l1.Val, 0)
// 			res += sum
// 			l1 = l1.Next
// 			counter++
// 			continue
// 		}
// 		if l2 != nil {
// 			sum := calc(counter, 0, l2.Val)
// 			res += sum
// 			l2 = l2.Next
// 			counter++
// 			continue
// 		}
// 	}

// 	//fmt.Println(res)
// 	return getReverseList(res)
// }

// func calc(index int, val1 int, val2 int) int {
// 	res1 := int(math.Pow(10, float64(index))) * val1
// 	res2 := int(math.Pow(10, float64(index))) * val2
// 	res := res1 + res2
// 	return res
// }

// func getReverseList(i int) (listNode *ListNode) {
// 	str := strconv.Itoa(i)
// 	var node ListNode
// 	for _, c := range str {
// 		i, _ := strconv.Atoi(string(c))
// 		node = ListNode{Val: i}
// 		if listNode != nil {
// 			node.Next = listNode
// 		}
// 		listNode = &ListNode{Val: node.Val, Next: node.Next}
// 	}
// 	return
// }

// // /**
// //   - Definition for singly-linked list.
// //     */
type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	res1, _ := getReverseStr(l1)
// 	res2, _ := getReverseStr(l2)

// 	ret := getReverseList(res1 + res2)

// 	fmt.Println(ret.Val, ret.Next.Val, ret.Next.Next.Val)

// 	return ret
// }

// // i, err := strconv.Atoi("-42")
// // s := strconv.Itoa(-42)
// func getReverseStr(listNode *ListNode) (int, error) {
// 	var res string
// 	node := listNode
// 	for node.Next != nil {
// 		res = strconv.Itoa(node.Val) + res
// 		node = node.Next
// 	}
// 	res = strconv.Itoa(node.Val) + res
// 	node = node.Next

// 	return strconv.Atoi(res)
// }
