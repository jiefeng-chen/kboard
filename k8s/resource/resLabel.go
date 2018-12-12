package resource


type Selector struct {
	MatchLabels map[string]string `yaml:"matchLabels"`
	MatchExpressions []*MatchExpressions `yaml:"matchExpressions"` // 匹配正则表达式，如：{key: tier, operator: In, values: [frontend]}
}


type MatchExpressions struct {
	Key string
	Operator string // In
	Values []string
}


