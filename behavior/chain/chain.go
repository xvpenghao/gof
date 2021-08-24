package chain

import (
	"fmt"
	"strings"
)

// 责任链的 两种写法

// LinkSensitiveWordFilter 敏感词的过滤-链表的形式
type LinkSensitiveWordFilter interface {
	Filter(content string) bool
	SetNext(LinkSensitiveWordFilter)
}

// AdSensitiveWordFilter 广告
type AdSensitiveWordFilter struct {
	next LinkSensitiveWordFilter
}

// Filter 实现过滤算法
func (this *AdSensitiveWordFilter) Filter(content string) bool {
	fmt.Println("广告算法词过滤")
	if strings.Contains(content, "广告") {
		return true
	}
	return this.next.Filter(content)
}

func (this *AdSensitiveWordFilter) SetNext(next LinkSensitiveWordFilter) {
	this.next = next
}

// PoliticalWordFilter 政治敏感
type PoliticalWordFilter struct {
	next LinkSensitiveWordFilter
}

// Filter 实现过滤算法
func (this *PoliticalWordFilter) Filter(content string) bool {
	fmt.Println("政治敏感-过滤")
	if strings.Contains(content, "政治") {
		return true
	}
	if this.next != nil {
		return this.next.Filter(content)
	}
	return false
}

func (this *PoliticalWordFilter) SetNext(next LinkSensitiveWordFilter) {
	this.next = next
}

// 数组的形式

// ArraySensitiveWordFilter 敏感词的过滤-数组的的形式
type ArraySensitiveWordFilter interface {
	Filter(content string) bool
}

// SensitiveWordFilterChain 设置一个 数组的过滤链
type SensitiveWordFilterChain struct {
	filters []ArraySensitiveWordFilter
}

func (this *SensitiveWordFilterChain) AddFilter(f ArraySensitiveWordFilter) {
	this.filters = append(this.filters, f)
}

func (this *SensitiveWordFilterChain) Filter(content string) bool {
	for _, f := range this.filters {
		if f.Filter(content) {
			return true
		}
	}
	return false
}

// AryAdSensitiveWordFilter 广告
type AryAdSensitiveWordFilter struct{}

// Filter 实现过滤算法
func (this *AryAdSensitiveWordFilter) Filter(content string) bool {
	fmt.Println("广告算法词过滤")
	if strings.Contains(content, "广告") {
		return true
	}
	return false
}

// AryPoliticalWordFilter 政治敏感
type AryPoliticalWordFilter struct{}

// Filter 实现过滤算法
func (this *AryPoliticalWordFilter) Filter(content string) bool {
	fmt.Println("政治敏感-过滤")
	if strings.Contains(content, "政治") {
		return true
	}
	return false
}
