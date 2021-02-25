# PHP 规范



## 1. 概览



- 源代码文件`必须` 以 `<?php` 标签开始，纯 PHP 代码文件必须省略最后的`?>`标签。
- 源代码文件`必须` 以 `不带BOM` 的 `UTF-8` 编码。
- 一个源文件只用来做声明（`类(class)`，`函数(function)`，`常量(constant)`等）或其他会产生`从属效应` 的操作（例如：输出信息，修改`.ini`配置等）,但`不能`同时做两件事。
- `命名空间(namespace)`和`类(class)` `必须`遵守 `PSR-0` 标准。
- `类名(class name)` `必须`使用`骆驼式(StudlyCaps)`写法。
- `类(class)`中的常量`必须`只由大写字母和`下划线(_)`组成。
- `方法名(method name)` `必须`使用`驼峰式(cameCase)`写法。
- 代码`必须`使用4个空格来进行缩进，而不是用制表符。
- 一行代码的字符数应该保持在`80`个之内，一定不可多于`120`个，但不硬性限制。
- 在`命名空间(namespace)`的声明下面`必须`有一行空行，并且在`导入(use)`的声明下面也`必须`有一行空行。
- `类(class)`的左花括号`必须`放到其声明下面自成一行，右花括号则`必须`放到类主体下面自成一行。
- `方法(method)`的左花括号`必须`放到其声明下面自成一行，右花括号则`必须`放到方法主体的下一行。
- 所有的`属性(property)`和`方法(method)` `必须`有可见性声明；`抽象(abstract)`和`终结(final)`声明`必须`在可见性声明之前；而`静态(static)`声明`必须`在可见性声明之后。
- 在控制结构关键字的后面`必须`有一个空格；而`方法(method)`和`函数(function)`的关键字的后面`不可`有空格。
- 控制结构的左花括号`必须`跟其放在同一行，右花括号`必须`放在该控制结构代码主体的下一行。
- 控制结构的左括号之后`不可`有空格，右括号之前也`不可`有空格。



## 2. 通则



### 2.1. PHP标签



PHP代码`必须`只使用`长标签(<?php ?>)`；而`不可`使用其他标签。



### 2.2. 字符编码



PHP代码的编码格式`必须`只使用不带`字节顺序标记(BOM)`的`UTF-8`。



### 2.3 源文件



所有的PHP源文件`必须`使用Unix LF(换行)作为行结束符。

所有PHP源文件`必须`以一个空行结束。

纯PHP代码源文件的关闭标签`?>` `必须`省略。



### 2.3. 行



行长度`没有`硬限制。

行长度的软限制`必须`是120个字符；对于软限制，代码风格检查器`必须`警告但`不可`报错。

一行代码的长度`不建议`超过80个字符；较长的行`建议`拆分成多个不超过80个字符的子行。

在非空行后面`不可`有空格。

空行`可以`用来增强可读性和区分相关代码块。

一行`不可`多于一个语句。



### 2.4. 缩进



代码`必须`使用4个空格，且`不可`使用制表符来作为缩进。

> 注意：代码中只使用空格，且不和制表符混合使用，将会对避免代码差异，补丁，历史和注解中的一些问题有帮助。空格的使用还可以使通过调整细微的缩进来改进行间对齐变得更加的简单。



### 2.5. 关键字和 True/False/Null



PHP关键字([keywords](http://php.net/manual/en/reserved.keywords.php))`必须`使用小写字母。

PHP常量`true`, `false`和`null` `必须`使用小写字母。



### 2.6. 副作用



一个源文件只用来做声明（`类(class)`，`函数(function)`，`常量(constant)`等）或者只用来做一些引起副作用的操作（例如：输出信息，修改`.ini`配置等）,但`不能`同时做这两件事。

短语`副作用(side effects)`的意思是 *在包含文件时* 所执行的逻辑与所声明的`类(class)`，`函数(function)`，`常量(constant)`等没有直接的关系。

`副作用(side effects)`包含但不局限于：产生输出，显式地使用`require`或`include`，连接外部服务，修改ini配置，触发错误或异常，修改全局或者静态变量，读取或修改文件等等

下面是一个既包含声明又有副作用的示例文件；即应避免的例子：

```
<?php``// 副作用：修改了ini配置``ini_set``(``'error_reporting'``, E_ALL);` `// 副作用：载入了文件``include` `"file.php"``;` `// 副作用：产生了输出``echo` `"<html>\n"``;` `// 声明``function` `foo()``{``  ``// 函数体``}

```

下面是一个仅包含声明的示例文件；即应提倡的例子：

```php
<?php``// 声明``function` `foo()``{``  ``// 函数体``}` `// 条件式声明不算做是副作用``if` `(! function_exists(``'bar'``)) {``  ``function` `bar()``  ``{``    ``// 函数体``  ``}``}

```



## 3. 命名空间和类



`命名空间(namespace)`和`类(class)`必须遵守 [PSR-0](https://github.com/hfcorriez/fig-standards/blob/zh_CN/接受/PSR-0.md).

这意味着一个源文件中只能有一个`类(class)`，并且每个`类(class)`至少要有一级`空间名（namespace）`：即一个顶级的`组织名(vendor name)`。

`类名(class name)` `必须`使用`StudlyCaps`写法。

代码`必须`使用正式的`命名空间(namespace)` 例子：

```
<?php``namespace` `Vendor\Model;` `class` `Foo``{``}
```



### 3.1 命名空间（namespace）和导入（use）声明



`命名空间(namespace)`的声明后面`必须`有一行空行。

所有的`导入(use)`声明`必须`放在`命名空间(namespace)`声明的下面。

一句声明中，`必须`只有一个`导入(use)`关键字。

在`导入(use)`声明代码块后面`必须`有一行空行。

示例：

```
<?php``namespace` `Vendor\Package;` `use` `FooClass;``use` `BarClass ``as` `Bar;``use` `OtherVendor\OtherPackage\BazClass;` `// ... 其它PHP代码 ...
```



## 4. 类(class)，属性(property)和方法(method)



术语“类”指所有的`类(class)`，`接口(interface)`和`特性(trait)`。



### 4.1. 扩展(extend)和实现(implement)



一个类的`扩展(extend)`和`实现(implement)`关键词`必须`和`类名(class name)`在同一行。

`类(class)`的左花括号`必须`放在下面自成一行；右花括号必须放在`类(class)`主体的后面自成一行。

```
<?php``namespace` `Vendor\Package;` `use` `FooClass;``use` `BarClass ``as` `Bar;``use` `OtherVendor\OtherPackage\BazClass;` `class` `ClassName ``extends` `ParentClass ``implements` `\ArrayAccess, \Countable``{``  ``// 常量、属性、方法``}
```



`实现(implement)`列表`可以`被拆分为多个缩进了一次的子行。如果要拆成多个子行，列表的第一项`必须`要放在下一行，并且每行`必须`只有一个`接口(interface)`。

```
<?php``namespace` `Vendor\Package;` `use` `FooClass;``use` `BarClass ``as` `Bar;``use` `OtherVendor\OtherPackage\BazClass;` `class` `ClassName ``extends` `ParentClass ``implements``  ``\ArrayAccess,``  ``\Countable,``  ``\Serializable``{``  ``// 常量、属性、方法``}
```





### 4.2. 常量



类常量`必须`只由大写字母和`下划线(_)`组成。 例子：

```
<?php``namespace` `Vendor\Model;` `class` `Foo``{``  ``const` `VERSION = ``'1.0'``;``  ``const` `DATE_APPROVED = ``'2012-06-01'``;``}
```





### 4.3. 属性



所有的`属性(property)`都`必须`声明其可见性。

`变量(var)`关键字`不可`用来声明一个`属性(property)`。

一条语句`不可`声明多个`属性(property)`。

`属性名(property name)` `不推荐`用单个下划线作为前缀来表明其`保护(protected)`或`私有(private)`的可见性。

一个`属性(property)`声明看起来应该像下面这样。

```
<?php``namespace` `Vendor\Package;` `class` `ClassName``{``  ``public` `$foo` `= null;``}
```





### 4.4. 方法(method)



所有的`方法(method)`都`必须`声明其可见性。

`方法名(method name)` `不推荐`用单个下划线作为前缀来表明其`保护(protected)`或`私有(private)`的可见性。

`方法名(method name)`在其声明后面`不可`有空格跟随。其左花括号`必须`放在下面自成一行，且右花括号`必须`放在方法主体的下面自成一行。左括号后面`不可`有空格，且右括号前面也`不可`有空格。

一个`方法(method)`声明看来应该像下面这样。 注意括号，逗号，空格和花括号的位置：

```
<?php``namespace` `Vendor\Package;` `class` `ClassName``{``  ``public` `function` `fooBarBaz(``$arg1``, &``$arg2``, ``$arg3` `= [])``  ``{``    ``// 方法主体部分``  ``}``}
```



### 4.5. 方法(method)的参数



在参数列表中，逗号之前`不能`有空格，而逗号之后则`必须`要有一个空格。

`方法(method)`中有默认值的参数必须放在参数列表的最后面。

```
<?php``namespace` `Vendor\Package;` `class` `ClassName``{``  ``public` `function` `foo(``$arg1``, &``$arg2``, ``$arg3` `= [])``  ``{``    ``// 方法主体部分``  ``}``}
```



参数列表`可以`被拆分为多个缩进了一次的子行。如果要拆分成多个子行，参数列表的第一项`必须`放在下一行，并且每行`必须`只有一个参数。

当参数列表被拆分成多个子行，右括号和左花括号之间`必须`有一个空格并且自成一行。

```
<?php``namespace` `Vendor\Package;` `class` `ClassName``{``  ``public` `function` `aVeryLongMethodName(``    ``ClassTypeHint ``$arg1``,``    ``&``$arg2``,``    ``array` `$arg3` `= []``  ``) {``    ``// 方法主体部分``  ``}``}
```



### 4.6. 抽象(abstract)，终结(final)和 静态(static)



当用到`抽象(abstract)`和`终结(final)`来做类声明时，它们`必须`放在可见性声明的前面。

而当用到`静态(static)`来做类声明时，则`必须`放在可见性声明的后面。

```
<?php``namespace` `Vendor\Package;` `abstract` `class` `ClassName``{``  ``protected` `static` `$foo``;` `  ``abstract` `protected` `function` `zim();` `  ``final` `public` `static` `function` `bar()``  ``{``    ``// 方法主体部分``  ``}``}
```





### 4.7. 调用方法和函数



调用一个方法或函数时，在方法名或者函数名和左括号之间`不可`有空格，左括号之后`不可`有空格，右括号之前也`不可`有空格。参数列表中，逗号之前`不可`有空格，逗号之后则`必须`有一个空格。

```
<?php``bar();``$foo``->bar(``$arg1``);``Foo::bar(``$arg2``, ``$arg3``);
```



参数列表`可以`被拆分成多个缩进了一次的子行。如果拆分成子行，列表中的第一项`必须`放在下一行，并且每一行`必须`只能有一个参数。

```
<?php``$foo``->bar(``  ``$longArgument``,``  ``$longerArgument``,``  ``$muchLongerArgument``);
```





## 5. 控制结构



下面是对于控制结构代码风格的概括：

- 控制结构的关键词之后`必须`有一个空格。
- 控制结构的左括号之后`不可`有空格。
- 控制结构的右括号之前`不可`有空格。
- 控制结构的右括号和左花括号之间`必须`有一个空格。
- 控制结构的代码主体`必须`进行一次缩进。
- 控制结构的右花括号`必须`主体的下一行。

每个控制结构的代码主体`必须`被括在花括号里。这样可是使代码看上去更加标准化，并且加入新代码的时候还可以因此而减少引入错误的可能性。



### 5.1. if，elseif，else



下面是一个`if`条件控制结构的示例，注意其中括号，空格和花括号的位置。同时注意`else`和`elseif`要和前一个条件控制结构的右花括号在同一行。

```
<?php``if` `(``$expr1``) {``  ``// if body``} ``elseif` `(``$expr2``) {``  ``// elseif body``} ``else` `{``  ``// else body;``}
```





### 5.2. switch，case



下面是一个`switch`条件控制结构的示例，注意其中括号，空格和花括号的位置。`case`语句`必须`要缩进一级，而`break`关键字（或其他中止关键字）`必须`和`case`结构的代码主体在同一个缩进层级。如果一个有主体代码的`case`结构故意的继续向下执行则`必须`要有一个类似于`// no break`的注释。

```
<?php``switch` `(``$expr``) {``  ``case` `0:``    ``echo` `'First case, with a break'``;``    ``break``;``  ``case` `1:``    ``echo` `'Second case, which falls through'``;``    ``// no break``  ``case` `2:``  ``case` `3:``  ``case` `4:``    ``echo` `'Third case, return instead of break'``;``    ``return``;``  ``default``:``    ``echo` `'Default case'``;``    ``break``;``}
```





### 5.3. while，do while



下面是一个`while`循环控制结构的示例，注意其中括号，空格和花括号的位置。

```
<?php``while` `(``$expr``) {``  ``// structure body``}
```



下面是一个`do while`循环控制结构的示例，注意其中括号，空格和花括号的位置。

```
<?php``do` `{``  ``// structure body;``} ``while` `(``$expr``);
```



### 5.4. for



下面是一个`for`循环控制结构的示例，注意其中括号，空格和花括号的位置。

```
<?php``for` `(``$i` `= 0; ``$i` `< 10; ``$i``++) {``  ``// for body``}
```





### 5.5. foreach



下面是一个`foreach`循环控制结构的示例，注意其中括号，空格和花括号的位置。

```
<?php``foreach` `(``$iterable` `as` `$key` `=> ``$value``) {``  ``// foreach body``}
```





### 5.6. try, catch



下面是一个`try catch`异常处理控制结构的示例，注意其中括号，空格和花括号的位置。

```
<?php``try` `{``  ``// try body``} ``catch` `(FirstExceptionType ``$e``) {``  ``// catch body``} ``catch` `(OtherExceptionType ``$e``) {``  ``// catch body``}
```





## 6. 闭包



声明闭包时所用的`function`关键字之后`必须`要有一个空格，而`use`关键字的前后都要有一个空格。

闭包的左花括号`必须`跟其在同一行，而右花括号`必须`在闭包主体的下一行。

闭包的参数列表和变量列表的左括号后面`不可`有空格，右括号的前面也`不可`有空格。

闭包的参数列表和变量列表中逗号前面`不可`有空格，而逗号后面则`必须`有空格。

闭包的参数列表中带默认值的参数`必须`放在参数列表的结尾部分。

下面是一个闭包的示例。注意括号，空格和花括号的位置。

```
<?php``$foo``->bar(``  ``$arg1``,``  ``function` `(``$arg2``) ``use` `(``$var1``) {``    ``// body``  ``},``  ``$arg3``);
```



参数列表和变量列表`可以`被拆分成多个缩进了一级的子行。如果要拆分成多个子行，列表中的第一项`必须`放在下一行，并且每一行`必须`只放一个参数或变量。

当列表（不管是参数还是变量）最终被拆分成多个子行，右括号和左花括号之间`必须`要有一个空格并且自成一行。

下面是一个参数列表和变量列表被拆分成多个子行的示例。

```
<?php``$longArgs_noVars` `= ``function` `(``  ``$longArgument``,``  ``$longerArgument``,``  ``$muchLongerArgument``) {``  ``// body``};` `$noArgs_longVars` `= ``function` `() ``use` `(``  ``$longVar1``,``  ``$longerVar2``,``  ``$muchLongerVar3``) {``  ``// body``};` `$longArgs_longVars` `= ``function` `(``  ``$longArgument``,``  ``$longerArgument``,``  ``$muchLongerArgument``) ``use` `(``  ``$longVar1``,``  ``$longerVar2``,``  ``$muchLongerVar3``) {``  ``// body``};` `$longArgs_shortVars` `= ``function` `(``  ``$longArgument``,``  ``$longerArgument``,``  ``$muchLongerArgument``) ``use` `(``$var1``) {``  ``// body``};` `$shortArgs_longVars` `= ``function` `(``$arg``) ``use` `(``  ``$longVar1``,``  ``$longerVar2``,``  ``$muchLongerVar3``) {``  ``// body``};

```

把闭包作为一个参数在函数或者方法中调用时，依然要遵守上述规则。

```
<?php``$foo``->bar(``  ``$arg1``,``  ``function` `(``$arg2``) ``use` `(``$var1``) {``    ``// body``  ``},``  ``$arg3``);
```







## 7. 其他

记录一些杂项。



- 不使用 AND， OR 关键字，使用 && ， || 替换。
- 字符串使用`单引号` **'** 包裹。