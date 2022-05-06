package trie

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type Node struct {
	sub      string
	children []*Node
	uid      int
}

func NewTrie(s string, uid int) *Node {
	return &Node{sub: s, uid: uid}
}

func minNumber(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (t *Node) Print(depth int) {
	var buf bytes.Buffer
	for i := 0; i < depth; i++ {
		buf.WriteString("\t")
	}
	buf.WriteString(fmt.Sprintf("(%v)", t.uid))
	buf.WriteString(t.sub)
	log.Println(buf.String())
	for _, child := range t.children {
		child.Print(depth + 1)
	}
}

func (t *Node) Register(s string, uid int) error {
	node := t
	newSub := s
	newUID := uid

	for len(node.children) > 0 {
		childIdx := findSimilarIndex(node.children, newSub)
		if childIdx == -1 {
			break
		}
		childLength := len(node.children[childIdx].sub)

		if childLength > len(newSub) {
			child := node.children[childIdx]
			child.sub = child.sub[len(newSub):]
			node.children = append(node.children[:childIdx], node.children[childIdx+1:]...)

			newNode := NewTrie(newSub, newUID)
			newNode.children = append(newNode.children, child)

			node.children = append(node.children, newNode)
			return nil
		} else if childLength < len(newSub) {
			newSub = newSub[childLength:]
		} else {
			// 잘못된거임
			log.Println("equal")
			return errors.New("겹친다")
		}
		node = node.children[childIdx]
	}

	node.children = append(node.children, NewTrie(newSub, newUID))
	return nil
}

func findSimilarIndex(nodes []*Node, str string) int {
	for i, node := range nodes {
		min := minNumber(len(node.sub), len(str))
		if node.sub[:min] == str[:min] {
			return i
		}
	}
	return -1
}
