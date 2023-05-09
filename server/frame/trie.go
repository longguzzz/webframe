package frame

import "strings"

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{newTrieNode()}
}

type trieNode struct {
	children trieNodeDict
	services *serviceList // 经常访问的serve其实可以缓存
}

func newTrieNode() *trieNode {
	return &trieNode{
		children: make(trieNodeDict),
		services: nil,
	}
}

func (t trieNode) tagPathNode() {
}

// slice可能扩容，所以这里应该用指针
// func (t trieNode) serviceList() *serviceList {

// fix: 两个都应该用指针
func (t *trieNode) serviceList() *serviceList {
	return t.services
}

type trieNodeDict map[string]*trieNode

func (tdict *trieNodeDict) getOrCreate(key string) (node matcherNode, exist bool) {
	if nextNode, ok := (*tdict)[key]; ok {
		return nextNode, true
	} else {
		(*tdict)[key] = newTrieNode()
		return (*tdict)[key], false
	}
}

/*
*  输入示例
*  /
*  /test
*  /test/
*  /test/test
 */

// 解耦router与service
func (t *trie) addRoute(path string) (Node matcherNode) {
	// TODO: 合法性检测
	if path == "/" { // 特殊情况： "/"在Split之后是["",""]
		return t.root
	}

	curNode := t.root
	parts := strings.Split(path, "/")[1:] // root跳过
	for i := 0; i < len(parts); i++ {
		// TODO: 内存不足时的错误处理
		nextNode, _ := curNode.children.getOrCreate(parts[i])
		curNode = nextNode.(*trieNode)
	}
	return curNode
}

func (t *trie) findRoute(path string) (target matcherNode) {
	if path == "/" {
		return t.root // 特殊情况： "/"在Split之后是["",""]
	}
	curNode := t.root
	parts := strings.Split(path, "/")[1:] // root跳过
	for i := 0; i < len(parts); i++ {
		if nextNode, ok := curNode.children[parts[i]]; ok {
			curNode = nextNode
		} else {
			curNode = nil
			break
		}
	}
	return curNode
}

// TODO：完成到一半的时候更换服务……？是否需要加锁？似乎……没必要？
func (t *trie) deleteRoute(path string) bool {
	if path == "/" {
		t.root.services = nil
		return true
	}
	lastSlash := strings.LastIndex(path, "/")
	var parentNode = t.root
	// if lastSlash > 0 { // fix: "/"、"/test"、".../test"的区别
	// 如果没有后继路径，则完全删除，还需要考虑更新Parent上的信息
	// 有后续路径则只是serve置为nil
	lastPathPart := path[lastSlash+1:]
	parentNode = t.findRoute(path[:lastSlash]).(*trieNode)
	if parentNode == nil {
		return false // 原先就不存在
	}
	if targetNode, ok := parentNode.children[lastPathPart]; ok {
		if len(targetNode.children) == 0 {
			delete(parentNode.children, lastPathPart)
			return true
		} else {
			targetNode.services = nil
			return true
		}
	} else {
		return false // 原先就不存在
	}

	// } else {
	// t.root.serve = nil
	// return true
	// }
}
