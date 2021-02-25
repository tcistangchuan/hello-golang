- 静态方法不能调用非静态属性。因为非静态属性需要实例化后，存放在对象里；打印出数据乱码：

```
	header('content-type:text/json;charset=utf-8');
	echo json_encode($data,JSON_UNESCAPED_UNICODE);exit;
```

- php-Redis 应用-限流：https://learnku.com/articles/30822?order_by=vote_count&

- #new self()和new static() 的区别：

```
首先，他们的区别只有在继承中才能体现出来，如果没有任何继承，那么这两者是没有区别的。
然后，new self()返回的实例是万年不变的，无论谁去调用，都返回同一个（父）类的实例，而new static()则是由调用者（调用类的实例）决定的。
```

- 获取命名空间加类名

```
$namespace=xxx::class 
```

- #抛异常和捕获异常

```php
function tsest(){
	throw new \Exception("this a error");
}

try {
    test()
}catch (\Exception $e){
    echo $e->getMessage()."\n";
    echo $e->getCode()."\n";
}
```

- 静态方法的注意事项：

  ```
  （1）调用类的静态函数时不会自动调用类的构造函数。
  （2）静态方法可以调用非静态方法，使用 self 关键词。php 里，一个方法被 self:: 后，它就自动转变为静态方法；
  （3）静态方法不能调用非静态属性。因为非静态属性需要实例化后，存放在对象里；
  ```

  

- 常用函数

  ```php
  func_get_arg() 获取函数实参，返回数组
  func_num_args() 获取函数实参的个数
  
  数组函数：
  array_get($arr,$key,"默认值") 
  获取数组的值并设置默认参数,若不设置默认值，可能返回null
    
  array_map($callback_func,$arr...Array) 
  为数组的每个元素应用回调函数 返回新数组
    
  array_walk(&$array,$callback_func) 
  使用用户自定义函数对数组中的每个元素做回调处理  返回bool
    
  array_combine ( array $keys , array $values ) : array
  用来自 keys 数组的值作为键名，来自 values 数组的值作为相应的值
    
  array_diff ( array $array1 , array $array2 [, array $... ] ) : array
  求差集：对比 array1 和其他一个或者多个数组，返回在 array1 中但是不在其他 array 里的值。
    
  array_keys(array $array [, mixed $search_value = null [, bool $strict = false]]) : array
  返回 input 数组中的数字或者字符串的键名。
  
  array_merge ( array $array1 [, array $... ] ) : array
  将一个或多个数组的单元合并起来，一个数组中的值附加在前一个数组的后面。返回作为结果的数组。
  如果输入的数组中有相同的字符串键名，则该键名后面的值将覆盖前一个值。然而，如果数组包含数字键名，后面的值将不会覆盖原来的值，而是附加到后面。
  如果只给了一个数组并且该数组是数字索引的，则键名会以连续方式重新索引。
   
  array_chunk ( array $array , int $size [, bool $preserve_keys = false ] ) : array
  将一个数组分割成多个数组，其中每个数组的单元数目由 size 决定。最后一个数组的单元数目可能会少于 size 个
    
  array_intersect ( array $array1 , array $array2 [, array $... ] ) : array
  求交集
  ```

- phpstorm的注释标记

  ```php
  参考链接：https://learnku.com/articles/30457
  
  好处：部分 IDE 可直接识别 注释标签 从而实现自动填充以及类型判断．
  
  /**
  *	@property WalletService $walletService
  * @method delCooperateMedication($id)
  */
  class A{
      /**
       * @var WalletService 
       */
      protected $walletService;//类的属性注释标记
     
    	public function delCooperateMedication($id){
        //类的方法注释标记
        //详细见上
        
        	/**
         * @var AssistantDoctorPatientDao $server
           */
          $server=app(AssistantDoctorPatientDao::class);
          $server->getPatient();
      }
  }
  ```

- laravel延迟任务实现原理-redis驱动

  ```
  文章参考地址：https://blog.csdn.net/raoxiaoya/article/details/103066624
  
  因为是需要一个常驻进程来检查的，我们使用redis存储会更好。
  
  延迟队列
  就是需要延迟一段时间后执行。Redis可通过zset来实现。我们可以将有序集合的value 设置为我们的消息任务(orderId)，把value的score设置为消息的到期时间，然后轮询获取有序集合的中的到期消息进行处理。
  
  返回有序集 key 中，所有 score 值介于 min 和 max 之间(包括等于 min 或 max )的成员。有序集成员按 score 值递增(从小到大)次序排列。
  	原始返回值
  ```

  

  