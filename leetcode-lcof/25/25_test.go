package leetcode

import (
	"testing"

	"github.com/lijinglin2019/algorithm-go/leetcode"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	result := "[1, 2, 3, 4, 5, 6, 7]"
	listNode1 := "[1, 3, 5, 7]"
	listNode2 := "[2, 4, 6]"
	assert.Equal(t, leetcode.NewListNode(result),
		mergeTwoLists(
			leetcode.NewListNode(listNode1),
			leetcode.NewListNode(listNode2),
		),
	)
	assert.Equal(t, leetcode.NewListNode(result),
		mergeTwoLists(
			leetcode.NewListNode(listNode2),
			leetcode.NewListNode(listNode1),
		),
	)
}
