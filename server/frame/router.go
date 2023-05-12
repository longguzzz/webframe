package frame

import (
	"net/http"
)

// 具体实现可以是Map或Trie
// 接口方式1：范型
// 接口方式2：返回值使用标记接口
type Matcher interface {
	node() matcherNode
	addRoute(path string) matcherNode
	findRoute(path string) matcherNode
	deleteRoute(path string) bool
}

// 匹配所得的节点【有】控制策略的功能，所以这里使用组合
// 因为同时也具备匹配功能，所以同时也实现了matcherNode接口
// 不应该在matcher里混入ctrlPanel。可以在routerMatcher返回之后，router对外提供能力的时候，插补ctrlPanel（组合）
// 关键在于，把trie的实现给隔离开，在写ctrlPanel相关内容的时候，可以不影响trie的实现、轻微影响router的实现
// 特定的模式串会返回一组方法集，在这个方法集上再通过ctrlPanel附加一些性质
//
//	type serviceNode struct {
//		trieNode
//		ctrlPanel
//	}
type serviceNode struct {
	matcherNode
	ctrlNode
}

type ctrlNode interface {
	middlewares() []ServeFunc
}

type matcherNode interface {
	asMatcher() Matcher
	selectChild(pathPart string) matcherNode
	createChild(pathPart string) matcherNode
	getServiceList() *serviceList
	setServiceList(*serviceList) 
	getGroup() *Group
	setGroup(*Group)
}

type serviceList []service

func (sList serviceList) findBy(method string) *service {
	for i := range sList {
		if sList[i].Method == method {
			return &sList[i]
		}
	}
	return nil
}

func (sList *serviceList) createOrUpdate(newService service) (oldService service) {
	foundIt := false
	// TODO: 注意修改slice可能的潜在bug
	for i := range *sList {
		if (*sList)[i].Method == newService.Method {
			foundIt = true
			oldService = (*sList)[i]
			(*sList)[i].doWith = newService.doWith
		}
	}
	if !foundIt {
		*sList = append(*sList, newService)
	}
	return oldService
}

type service struct {
	Method string
	doWith ServeFunc
}

type Router struct {
	Matcher
	*Group
	// ServeFuncMap map[string]ServeFunc
}

func newRouter(matcher Matcher) *Router { // 依赖注入
	r := &Router{Matcher: matcher}
	r.Group = &Group{router: r, matcher: matcher}
	return r
}

// func (r *Router) RigisterGET(pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
// 	return r.register("GET", pattern, serveFunc)
// }

// func (r *Router) RigisterPOST(pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
// 	return r.register("POST", pattern, serveFunc)
// }

// // oldServeFunc为nil则之前没有别的服务
// func (r *Router) register(method, pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
// 	matched := r.addRoute(pattern)
// 	sList := matched.getServiceList()
// 	if sList == nil {
// 		*sList = serviceList{service{
// 			method,
// 			serveFunc,
// 		}}
// 	}

// 	sList.createOrUpdate(service{
// 		method, serveFunc,
// 	})
// 	return oldServeFunc
// }

// TODO: 修改函数签名，添加错误码
func (r *Router) serveWith(c *RequestContext) {
	path := c.Request.URL.Path
	if matched := r.findRoute(path); matched != nil {
		if service := matched.getServiceList().findBy(c.Request.Method); service != nil {
			service.doWith(c)
		}
	} else {
		c.RespString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request.URL)
	}
}
