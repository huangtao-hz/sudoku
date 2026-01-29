package main

// Set 定义数独单元格的可选数，采用整数位形式，使用第1到第9位来表示
type Set int

// NewSet 构造函数
func NewSet(vals ...int) *Set {
	var s Set = 0
	s.Add(vals...)
	return &s
}

// Add 添加元素
func (s *Set) Add(vals ...int) {
	for _, i := range vals {
		*s |= (1 << i)
	}
}

// Remove 删除元素
func (s *Set) Remove(vals ...int) {
	for _, i := range vals {
		*s &^= (1 << i)
	}
}

// String 转化为字符串
func (s *Set) String() string {
	result := make([]byte, 0, 9)
	value := *s
	for i := 1; value != 0; i++ {
		value >>= 1
		if value&1 == 1 {
			result = append(result, byte(i+'0'))
		}
	}
	return string(result)
}

// Difference 求减集
func (s Set) Difference(other *Set) *Set {
	s &^= *other
	return &s
}

// Intersection 求交集
func (s Set) Intersection(other *Set) *Set {
	s &= *other
	return &s
}

// Union 求并集
func (s Set) Union(other *Set) *Set {
	s |= *other
	return &s
}

// Contains 判断是否包含
func (s *Set) Contains(vals ...int) bool {
	d := NewSet(vals...)
	return *(d.Difference(s))&-1 == 0
}

// ToSlice 转换为 Slice
func (s *Set) ToSlice() (result []int) {
	result = make([]int, 0, 9)
	value := *s
	for i := 1; value != 0; i++ {
		value >>= 1
		if value&1 == 1 {
			result = append(result, i)
		}
	}
	return
}

// Clone 制作副本
func (s Set) Clone() *Set {
	return &s
}

// Count 统计数量
func (s *Set) Count() (count int) {
	value := int(*s)
	for value != 0 {
		value >>= 1
		count += value & 1
	}
	return
}

// GetOnlyValue 获取唯一的值，值不唯一时返回 0
func (s *Set) GetOnlyVlaue() (value int) {
	vals := s.ToSlice()
	if len(vals) == 1 {
		value = vals[0]
	}
	return
}
