# 模糊专家系统——依恋类型人格测试
总共有18道题目，每道题有1~5五个选项，分别为：1代表完全不符合，2代表较不符合，3代表不能确定，4代表较符合，5代表完全符合。
我们根据心理学所给出的的法则，运用模糊集的知识，计算出被测者的依恋类型。

# 具体实现
## 后端
基于golang的gin框架实现，AnalyserFuzzy里是一个用于模糊分析的对象，写成了接口便于替换其它的Analyser。

## 前端页面
基于vue2.x实现，用到elementui和promise组件。