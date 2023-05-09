package frame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsePattern(t *testing.T) {

}

// fix: 注意strings.Split的特殊行为，"/"会被分成["",""]
// TODO，通过objdump验证一下，匿名函数是会被存在.text段还是.data段
func Test_Trie(t *testing.T) {
	// trie := newTrie()

	// f1 := ServeFunc(func(c *Context) {
	// })
	// trie.addRoute("/test")

	// p := func(path string) uintptr { // 返回值不是unsafe.Pointer
	// 	return reflect.ValueOf(trie.findRoute(path).serve).Pointer()
	// }
	// ptrOfFunc := func(f ServeFunc) uintptr {
	// 	return reflect.ValueOf(f).Pointer()
	// }
	// assert.Equal(t, trie.findRoute("/tx"), (*trieNode)(nil)) // 类型也要相同
	// assert.Equal(t, p("/test"), ptrOfFunc(f1))               // 函数间可以赋值，但是不可以直接比较

	trie := newTrie()

	matched1 := trie.addRoute("/test")
	matched2 := trie.addRoute("/test/t")
	matched3 := trie.addRoute("/")
	matched4 := trie.addRoute("/t")

	assert.Equal(t, trie.findRoute("/tx"), (*trieNode)(nil)) // 类型也要相同
	assert.Equal(t, trie.findRoute("/test"), (matched1))     // 函数间可以赋值，但是不可以直接比较
	assert.Equal(t, trie.findRoute("/test/t"), (matched2))
	assert.Equal(t, trie.findRoute("/"), (matched3))
	assert.Equal(t, trie.findRoute("/t"), (matched4))
	assert.Equal(t, trie.findRoute("/test/"), (*trieNode)(nil))

	assert.Equal(t, trie.deleteRoute("/"), true)
	assert.Equal(t, trie.findRoute("/tx"), (*trieNode)(nil)) // 类型也要相同
	assert.Equal(t, trie.findRoute("/test"), (matched1))     // 函数间可以赋值，但是不可以直接比较
	assert.Equal(t, trie.findRoute("/test/t"), (matched2))
	// assert.Equal(t, trie.findRoute("/"), (nil)) // 注意root的特殊性
	
	assert.Equal(t, trie.findRoute("/").serviceList(), (*serviceList)(nil))
	assert.Equal(t, trie.findRoute("/t"), (matched4))
	assert.Equal(t, trie.findRoute("/test/"), (*trieNode)(nil))

	assert.Equal(t, trie.deleteRoute("/test/"), false)
	assert.Equal(t, trie.deleteRoute("/t"), true)
	assert.Equal(t, trie.root.children["t"], (*trieNode)(nil))
	assert.Equal(t, len(trie.root.children), 1)
}
