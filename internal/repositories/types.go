package repositories

import "mono-base/pkg/types"

type Condition struct {
	Field string
	Value interface{}
	Op    string // "eq", "ne", "lt", "gt", "lte", "gte", "in", "like", ...
}

type CommonCondition struct {
	Conditions []Condition
	Sorting    []types.Sorting
	Paging     *types.Paging
}

func NewCommonCondition() *CommonCondition {
	return &CommonCondition{
		Conditions: []Condition{},
		Sorting:    []types.Sorting{},
		Paging:     &types.Paging{},
	}
}

func (cc *CommonCondition) AddCondition(field string, value interface{}, op string) {
	cc.Conditions = append(cc.Conditions, Condition{
		Field: field,
		Value: value,
		Op:    op,
	})
}

func (cc *CommonCondition) SetPaging(limit, page uint64) {
	cc.Paging.Limit = limit
	cc.Paging.Page = page
}

func (cc *CommonCondition) AddSorting(field, order string) {
	cc.Sorting = append(cc.Sorting, types.Sorting{
		Field: field,
		Order: order,
	})
}

func (cc *CommonCondition) WithPaging(limit, page uint64) *CommonCondition {
	cc.Paging = &types.Paging{
		Limit: limit,
		Page:  page,
	}
	return cc
}

func (cc *CommonCondition) WithCondition(field string, value interface{}, op string) *CommonCondition {
	condition := Condition{
		Field: field,
		Value: value,
		Op:    op,
	}
	cc.Conditions = append(cc.Conditions, condition)
	return cc
}

func (cc *CommonCondition) WithSorting(field string, order string) *CommonCondition {
	sorting := types.Sorting{
		Field: field,
		Order: order,
	}
	cc.Sorting = append(cc.Sorting, sorting)
	return cc
}
