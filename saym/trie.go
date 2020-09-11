package saym

type node struct {
	pattern  string  //待匹配路由，列如 /p/:lang
	part     string  //路由中的一部分，列如 :lang
	children []*node //子节点，列如[doc,intro]
	isWild   bool    //是否精准匹配，part 含有 ：或者*时为true
}
