- for...range 和 for的区别

  ```go
  (1)for i := range a {} 转换成如下所示的逻辑:
    ha := a
    hv1 := 0
    hn := len(ha)
    v1 := hv1
    for ; hv1 < hn; hv1++ {
        v1 := hv1
        ...
    }
  (2)for i,v := range a {} 转换成如下所示的逻辑:
    ha := a
    hv1 := 0
    hn := len(ha)
    v1 := hv1
    for ; hv1 < hn; hv1++ {
        tmp := ha[hv1]
        v1, v2 := hv1, tmp
        ...
    }
  ```

  