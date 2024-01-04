package graph

import "src/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	baseModels []*model.BaseModel
	campaigns []*model.Campaign
	events []*model.Event
	locations []*model.Location
	projects []*model.Project
	tasks []*model.Task
	task_assignments []*model.TaskAssignment
	users []*model.User


}
