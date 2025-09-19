package view

import "tonix/backend/model"

type TagView struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Usages int    `json:"usages"`
}

func ToTagView(tag *model.Tag) TagView {
	return TagView{
		Id:     tag.Id,
		Name:   tag.Name,
		Type:   tag.Type,
		Usages: tag.Usages,
	}
}
