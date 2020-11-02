观察者模式

可参考步骤：
1.主题（被观察的对象）
2.观察者

“主题”(被观察的对象)接口Observable
抽象方法Attach: 增加“订阅者”
抽象方法Detach: 删除“订阅者”
抽象方法Notify: 通知“订阅者”

“订阅者”(观察者)接口ObserverInterface
抽象方法Do: 自身的业务