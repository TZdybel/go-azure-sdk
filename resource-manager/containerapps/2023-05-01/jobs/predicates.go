package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p JobOperationPredicate) Matches(input Job) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type JobExecutionOperationPredicate struct {
	EndTime   *string
	Id        *string
	Name      *string
	StartTime *string
	Type      *string
}

func (p JobExecutionOperationPredicate) Matches(input JobExecution) bool {

	if p.EndTime != nil && (input.EndTime == nil || *p.EndTime != *input.EndTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.StartTime != nil && (input.StartTime == nil || *p.StartTime != *input.StartTime) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
