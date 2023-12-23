# Go 语言设计模式

设计模式是为了让程序具有更好的`代码重用性`、`可读性`、`可扩展性`、`可靠性`、使程序呈现高内聚、低耦合的特性。

## 设计模式的七大原则

1. 单一职责原则：一个类应该只负责一项职责
2. 接口隔离原则：一个类对另一个类的依赖应该建立在最小的接口上
3. 依赖倒转原则：抽象不应该依赖细节，细节应该依赖抽象（面向接口编程）
4. 里氏替换原则：所有引用基类的地方必须能“透明”地使用其子类的对象（使用聚合、组合、依赖 替代 继承）
5. 开闭原则：一个软件实体（类、模块、函数）应该对扩展开放，对修改关闭。用抽象构建框架，用实现扩展细节
6. 迪米特法则（最少知道原则）：一个类对自己依赖的类知道的越少越好
7. 合成复用原则：尽量使用 合成/聚合 的方式，而不是使用继承

## 示例代码

### 创建型模式

- Singleton 单例模式
- Simple Factory 简单工厂模式
- Factory Method 工厂方法模式
- Abstract Factory 抽象工厂模式
- Prototype 原型模式
- Builder 创建者模式

### 结构型模式

- Adapter 适配器模式
- Bridge 桥模式
- Decorator 装饰模式
- Composite 组合模式
- Facade 外观模式
- Flyweight 享元模式
- Proxy 代理模式

### 行为型模式

- Template Method 模板方法模式
- Command 命令模式
- Visitor 访问者模式
- Iterator 迭代器模式
- Observer 观察者模式
- Mediator 中介者模式
- Memento 备忘录模式
- Interpreter 解释器模式
- State 状态模式
- Strategy 策略模式
- Chain of Responsibility 职责链模式
