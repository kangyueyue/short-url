package vo

type Operator string

// 操作符
const (
	Equal    Operator = "="
	NotEqual Operator = "!="
	Gt       Operator = ">"
	Lt       Operator = "<"
	Gte      Operator = ">="
	Lte      Operator = "<="
	Like     Operator = "like"
)

// FilterVo 过滤参数
type FilterVo struct {
	Field    string `json:"field"` // 字段名
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
