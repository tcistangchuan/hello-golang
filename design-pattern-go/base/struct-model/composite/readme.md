代码建模
组合模式采用的树结构的思想
参考步骤：
1.组件（节点）接口
2.组件对象（可采用基础组件对象镶嵌的方式）
3.设置根组件
4.组装组件

组件结构：
    成员属性
    ChildComponents: 子组件列表 -> 稳定不变的
    成员方法
    Mount: 添加一个子组件 -> 稳定不变的
    Remove: 移除一个子组件 -> 稳定不变的
    Do: 执行组件&子组件 -> 变化的
