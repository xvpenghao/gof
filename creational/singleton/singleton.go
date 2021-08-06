package singleton

// Singleton 饿汉式单例
type Singleton struct {
	Id int
} // 他是有个空的结构体 导致每次创建的new 打印他的地址，都一样。

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
