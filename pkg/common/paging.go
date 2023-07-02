package common

type Paging struct {
	Total int64 `json:"total" form:"total"`
	Limit int   `json:"pageSize" form:"pageSize"`
	Page  int   `json:"page,omitempty" form:"page"`
}

func (p *Paging) FullFill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 50
	}
}
