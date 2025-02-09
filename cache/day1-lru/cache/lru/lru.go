package lru

//lru  least recently Used 最近最少使用 缓存策略,优先淘汰最久未使用的数据
// 实现 如果数据最近被访问过，那么将来被访问的概率也会更高。LRU 算法的实现非常简单，维护一个队列，如果某条记录被访问了，
//则移动到队首，那么队尾则是最近最少访问的数据，淘汰该条记录即可。
import "container/list"

type Cache struct {
	maxBytes  int64                         //最大内存
	nbytes    int64                         //已用内存
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      //字典 键是字符串,值是节点对应的指针
	OnEvicted func(key string, value Value) // 是一个常见的回调函数名。它通常用于在缓存条目被驱逐（即从缓存中删除）时执行特定的操作
}

// 键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// Value 为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nbytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		//类型断言,
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

// 缓存淘汰
func (c *Cache) RemoveOldest() {
	element := c.ll.Back()
	if element != nil {
		//从链表中移除对应的元素
		c.ll.Remove(element)
		//取出元素,类型断言
		kv := element.Value.(*entry)
		//从维护字典中山删除
		delete(c.cache, kv.key)
		//删除键和值对应的内存量
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		// 删除后执行回调方法
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)

		}
	}

}

func (c *Cache) Add(key string, value Value) {
	//更新
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		element := c.ll.PushFront(&entry{key, value})
		c.cache[key] = element
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	//这里使用for 而不是if 是因为一个if语句只会检查条件一次，并执行一次代码块。
	//一个for循环会在每次执行循环体之后重新检查条件，直到条件为假。这样可以确保在每次移除最旧的条目后，重新检查缓存大小是否仍然超过限制。
	//如果缓存大小在移除一个条目后仍然超过限制，for循环会继续移除条目，直到缓存大小在限制范围内。
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}

}

func (c *Cache) Len() int {
	return c.ll.Len()
}
