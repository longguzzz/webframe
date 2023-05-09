### 问题1
前缀树与散列表之间的权衡与比较

### 问题2
w.Header().Set()
w.WriteHeader()
这两者的顺序不能换

### 问题3
http.Error(c.Writer, err.Error(),500)里边有用到WriteHeader()和w.Header.Set()
encoder.Encode(obj)也相当于调用了Write()
二者其一使用一次，就会导致另一个无法设置header

研究一下gin中的新做法，https://github.com/gin-gonic/gin/blob/master/render/json.go#L56
```golang
var (
	jsonContentType      = []string{"application/json; charset=utf-8"}
	jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	jsonASCIIContentType = []string{"application/json"}
)

// Render (JSON) writes data with custom ContentType.
func (r JSON) Render(w http.ResponseWriter) error {
	return WriteJSON(w, r.Data)
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(w http.ResponseWriter, obj any) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
```

### 问题4

type trieNode struct {
	children trieNodeDict
	services serviceList // 经常访问的serve其实可以缓存
}

func (t *trieNode) serviceList() *serviceList {
	return &t.services
}

assert.Equal(t, trie.findRoute("/").serviceList(), (*serviceList)(nil))

        	Error:      	Not equal: 
        	            	expected: &frame.serviceList(nil)
        	            	actual  : (*frame.serviceList)(nil)
        	            	
        	            	Diff:
        	Test:       	Test_Trie

对比

type trieNode struct {
	children trieNodeDict
	services *serviceList // 经常访问的serve其实可以缓存
}

func (t *trieNode) serviceList() *serviceList {
	return t.services
}

差异就在于，services置为nil之后，前者会导致解引用nil