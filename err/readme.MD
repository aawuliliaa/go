### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
```
应该Wrap这个error。
因为我们在程序处理异常时，经常会使用fmt.Errorf()携带一些上下文信息，这样当使用==判断原始错误的时候，就会判断失败。
在这种情况下，只能与err.Error()进行字符串判断或strings.Contains(err.Error(),"not found")来判断特殊类型的错误。
但这种基于字符串的判断是很不友好的，当依赖包版本升级，err中的报错信息变化，我们就要重新测试代码，验证错误内容，带来了不必要的工作。
例如：
当我们提供一个"新增应用"的接口给前端，我们后端逻辑处理为
app,err := findApp()
if err!=nil {
   if err==sql.ErrNoRows{  --- 这里要判断错误类型
       addApp()
   }else{
     return fmt.Errorf("get app:%s err,err is:%v",app,err)
   }
}else{
   return fmt.Errorf("app:%s already exist",app)
}
```
