# 桥接模式

将它的抽象部分(改层自身不完成具体的工作，它需要将具体的工作委派给实现部分层)，和它的实现部分进行分离，使他们都可以独立的变化

* 优点

0、分类抽象接口及其 实现部分

0、开闭原则，新增抽象部分 和实现部分，他们之间不会有任何影响。（桥接模式提高了系统的可扩充性，在两个变化维度中任意扩展一个维度，都不需要修改原有系统。）

0、单一职责原则。 抽象部分专注于处理高层逻辑， 实现部分处理平台细节（实现细节对客户透明，可以对用户隐藏实现细节）

* 缺点

0、桥接模式的引入会增加系统的理解与设计难度，由于聚合关联关系建立在抽象层，要求开发者针对抽象进 行设计与编程

* 解决问题

mac和windows电脑， 需要安装打印机 普通，和爱普生，进行打印机打印。

方法1，mac 和 windows 分别安装 惠普 和 爱普生， 4个类， 方法2，mac 和 windows 里面 各有 2个 打印方法，

如果新增一个 打印机呢，你可能需要 6个类，或者，再修改原代码新增 打印方法了， 类变的巨多，也违反了，开闭原则，

* 使用环境

0、一个类存在两个独立变化的维度，且这两个维度都需要进行扩展

