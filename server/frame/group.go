package frame

// trie保存底层的业务服务
// group添加控制策略和中间件

// 统一入口，都从group访问，trie只提供匹配功能……？
// serviceTrie
// groupTrie
// 匹配从groupTrie进入

// TODO:
// engine 可从engine访问root对应的group
// engine 可从engine查询特定pattern对应的group
// router 可创建group以及相应的matcher
// router 可标记group，但不创建相应的matcher
// router 可删除group，但不删除相应的matcher
// router 可修改group名

// VirtualGroup，可能会增加复杂度……？与路径不重合的控制树，相当于第二个入口……？
// 底层一棵树，上边附加group
// group是否需要id……？

// group创建子group
// group查询子group
// group删除子group
// group修改group名

// group获取底层trie
// 从trieNode获取相应的group

// 可以从group创建子group
// 与matcher有明显的对应关系

type Group struct {
	middlewares serviceList
	matcher     Matcher // 可以由trieNode转换来
	router      *Router
}

// pattern会清除开头的 / ，并且一次只支持一组
func (g *Group) NewSubGroup(pathPart string) *Group {
	if pathPart != "" && pathPart[:1] == "/" {
		pathPart = pathPart[1:]
	}
	if node := g.matcher.node().selectChild(pathPart); node != nil {
		*node.getGroup() = Group{
			matcher: node.asMatcher(),
			router:  g.router,
		}
		return node.getGroup()
	} else {
		node := g.matcher.node().createChild(pathPart)
		grp := &Group{
			matcher: node.asMatcher(),
			router:  g.router,
		}
		node.setGroup(grp)
		return grp
	}
}

func (g *Group) GetSubGroup(pattern string) *Group {
	if node := g.matcher.node().selectChild(pattern); node != nil {
		return node.getGroup()
	}
	return nil
}

func (g *Group) RegisterGET(serveFunc ServeFunc) {
	g.register("GET", serveFunc)
}

func (g *Group) RegisterPOST(serveFunc ServeFunc) {
	g.register("POST", serveFunc)
}

func (g *Group) register(method string, serveFunc ServeFunc) {
	serviceList := g.matcher.node().getServiceList()
	_ = serviceList.createOrUpdate(service{
		method,
		serveFunc,
	})
}
