package saym

type node struct {
	pattern  string  //待匹配路由，列如 /p/:lang
	part     string  //路由中的一部分，列如 :lang
	children []*node //子节点，列如[doc,intro]
	isWild   bool    //是否精准匹配，part 含有 ：或者*时为true
}

//第一个匹配成功的空的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

//所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {

}
