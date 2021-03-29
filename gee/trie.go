package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern string
	part string
	children []*node
	isWild bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入路由节点树
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height] // 根据层数赋值
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} // 新建子节点，:或*为真
		n.children = append(n.children, child) // 子节点添加进父节点的子节点表
	}
	child.insert(pattern, parts, height + 1)
}

func (n *node) search(parts []string, height int) *node{
	// 查找到头或者匹配到*
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil // 找不到
		}
		return n // 找到了
	}

	part := parts[height] // 查找层数
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height + 1) // 向下层递归查找
		if result != nil {
			return result
		}
	}
	return nil
}
