package frame

import (
	"net/http"
)

// 具体实现可以是Map或Trie
// 接口方式1：范型
// 接口方式2：返回值使用标记接口
type routerMatcher interface {
	addRoute(path string) (node matcherNode)
	findRoute(path string) (node matcherNode)
	deleteRoute(path string) bool
}

type matcherNode interface {
	tagPathNode()
	serviceList() *serviceList
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
	routerMatcher
	// ServeFuncMap map[string]ServeFunc
}

func newRouter(matcher routerMatcher) *Router { // 依赖注入
	return &Router{matcher}
}

func (r *Router) RigisterGET(pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
	return r.register("GET", pattern, serveFunc)
}

func (r *Router) RigisterPOST(pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
	return r.register("POST", pattern, serveFunc)
}

// oldServeFunc为nil则之前没有别的服务
func (r *Router) register(method, pattern string, serveFunc ServeFunc) (oldServeFunc ServeFunc) {
	matched := r.addRoute(pattern)
	sList := matched.serviceList()

	sList.createOrUpdate(service{
		method, serveFunc,
	})
	return oldServeFunc
}

// TODO: 修改函数签名，添加错误码
func (r *Router) serveWith(c *Context) {
	path := c.Request.URL.Path
	if matched := r.findRoute(path); matched != nil {
		if service := matched.serviceList().findBy(c.Request.Method); service != nil {
			service.doWith(c)
		}
	} else {
		c.RespString(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Request.URL)
	}
}
