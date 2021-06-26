# 模糊专家系统——依恋类型人格测试
总共有18道题目，每道题有1~5五个选项，分别为：1代表完全不符合，2代表较不符合，3代表不能确定，4代表较符合，5代表完全符合。
我们根据心理学所给出的的模糊推理法则，计算出被测者的依恋类型。
添加中间类型

# 具体实现

## 前后端交互格式
### 前端传后端
```json
{
    "ans":[...]
}
```

### 后端传前端
```json
{
    'attachment styles':"...",//依恋类型
    'description':"...",//描述
    'rate of intimacy':double,//亲密维度得分
    'rate of anxiety about abandonment':double//焦虑维度得分
}
```
