package AnalyserFuzzy

func init() {
	descriptions = make(map[string]string)
	styles = make(map[string]string)
	descriptions["secure"] = secure
	descriptions["preoccupied"] = preoccupied
	descriptions["fearful"] = fearful
	descriptions["dismissing"] = dismissing

	styles["secure"] = "安全型"
	styles["preoccupied"] = "痴迷型"
	styles["fearful"] = "恐惧型"
	styles["dismissing"] = "回避型"
}

var descriptions map[string]string
var styles map[string]string

type analyserFuzzy struct {
	RateOfIntimacy int
	RateOfAnxiety  int
}

func New() *analyserFuzzy {
	return new(analyserFuzzy)
}

// Analyse 处理被测者的答案，得到亲近得分和焦虑得分
func (a *analyserFuzzy) Analyse(ans []int) (rateOfIntimacy, rateOfAnxiety int, err error) {
	// 反转这部分的答案
	// 2，7，8，13，16，17，18
	// 数组下标从0开始，所以是题号减1
	toReverse := []int{1, 6, 7, 12, 15, 16, 17}
	for _, i := range toReverse {
		ans[i] = 6 - ans[i]
	}
	// 转成百分比表示
	for i := range ans {
		ans[i] *= 20
	}
	// 亲近分量表
	// 1,6,8,12,13,17
	qjArray := []int{0, 5, 7, 11, 12, 16}
	// 依赖分量表
	// 2,5,7,14,16,18
	ylArray := []int{1, 5, 6, 13, 15, 17}
	// 焦虑分量表
	// 3,4,9,10,11,15
	jlArray := []int{2, 3, 8, 9, 10, 14}
	qj := avg(ans, qjArray)
	yl := avg(ans, ylArray)
	rateOfAnxiety = avg(ans, jlArray)
	// 亲近依恋复合成一个维度, 此处使用模糊集的交运算
	rateOfIntimacy = and(qj, yl)
	a.RateOfAnxiety = rateOfAnxiety
	a.RateOfIntimacy = rateOfIntimacy
	return
}

// GetResult 根据得分获取模糊推理的结论
func (a *analyserFuzzy) GetResult() (attachmentStyle, description string) {
	var adv, style string
	// 从隶属度函数映射回自然语言。这里假定分数小于40或是大于80说明是很明确的
	if (a.RateOfAnxiety > 80 || a.RateOfAnxiety < 40) && (a.RateOfIntimacy > 80 || a.RateOfIntimacy < 40) {
		adv = "是典型的"
	} else {
		adv = "比较偏向于"
	}
	// 根据分数确定依恋风格
	if a.RateOfAnxiety < 60 && a.RateOfIntimacy > 60 {
		style = "secure"
	} else if a.RateOfAnxiety < 60 && a.RateOfIntimacy < 60 {
		style = "dismissing"
	} else if a.RateOfAnxiety > 60 && a.RateOfIntimacy > 60 {
		style = "preoccupied"
	} else {
		style = "fearful"
	}
	attachmentStyle = "您" + adv + styles[style] + "依恋风格"
	description = descriptions[style]
	return
}

// 计算a数组中下标在idx中的元素的平均值，假设idx一定是合法的
func avg(a []int, idx []int) (avg int) {
	for _, i := range idx {
		avg += a[i]
	}
	avg /= len(idx)
	return
}

func and(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const (
	secure = "安全型(secure)或安全—自主型(secure-autonomous)的人认为自己是值得爱的，他人也是值得爱和信任的。" +
		"安全型依恋是一种稳定和积极的情绪联系，以爱情关系中的关怀、亲密感、支持和理解为标志。" +
		"这种类型的人认为自己是友好、善良和可爱的人，也认为别人普遍是友好、可靠和值得信赖的人。" +
		"他们十分容易与其他人接近，总是放心地依赖他人和让别人依赖自己。一般来说，他们既不会过于担心被抛弃，" +
		"也不怕别人在感情上与自己过于亲近。无论我们自己的依恋方式属于哪一种，找到一个安全依恋型的人做自己的伴侣对于我们来说都是一件好事。"
	preoccupied = "痴迷型(preoccupied)，这是给予焦虑—矛盾型的新名称，因为这种类型的人依赖于他人的赞许来获得内心的安适坦然，所以他们过度地寻求认同，沉溺于人际关系。他们认为自己是不值得爱的和没有价值的，但是他人是可接受的，总是努力赢得他人的接纳，并以此支持消极的自我形象。" +
		"痴迷型依恋的特征是对人际关系怀着混合的情感，这就使人处于爱、恨、怀疑、拿不起、放不下的冲突情感之中，导致一种不稳定和矛盾的心理状态。通常，痴迷型的人总觉得自己被误解和不受赏识，认为自己的情人和朋友都不可靠，不愿意与自己建立持久的关系。矛盾依恋型的人担心他们的恋人并不真正爱自己，或者会离开自己。因此，他们一方面希望能与自己的恋人极为亲近，另一方面又对恋人是否可靠和可信满腹猜疑。"
	fearful    = "恐惧型(fearful)的人对自己和他人的态度都是消极的，这种类型的成人可能出于害怕被拒绝而极力避免和他人发生亲密关系。虽然他们希望有人喜欢自己，但更担心自己因此离不开别人，而一旦建立了亲密关系，又往往会过度担心伴侣会离开自己，整天提心吊胆。有时想到与伴侣亲密相处时他们就会感到恐惧。"
	dismissing = "回避型(dismissing)的人对个人的看法相对积极（自己是有价值的），但是认为他人会拒绝自己，和他人发生亲密关系得不偿失。这种类型的成人会以避免与他人发生联系来作为保护自己不受伤害的手段。他们拒绝和他人相互依赖，因为他们相信自己能自力更生，也不在乎他人是否喜欢自己。他们会更关注替代选择，会留心任何可能的其他爱情选择，更容易被新结识的人所吸引。同时，他们往往希望将来的伴侣不给他们提供帮助，因为他们不打算反过来做任何报答。"
)
