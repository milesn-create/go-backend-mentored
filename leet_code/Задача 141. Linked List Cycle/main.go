package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// самый оптимальный, меньше всего памяти расходует и оч быстрый
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}
func HasCycle(head *ListNode) bool {
	s := []*ListNode{}
	if head == nil || head.Next == nil {
		return false
	}
	for head != nil {
		for _, n := range s {
			if n == head {
				return true
			}
		}
		s = append(s, head)
		head = head.Next

	}
	return false
}
func HHasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if m[current] {
			return true
		}
		m[current] = true
		current = current.Next
	}
	return false
}
