// Copyright 2020-2021 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package openapi

import (
	"fmt"
	"github.com/daveshanley/vacuum/model"
	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
	"strings"
)

// PostResponseSuccess is a rule that will check if a post operations contain a successful response code or not.
type PostResponseSuccess struct {
}

// GetSchema returns a model.RuleFunctionSchema defining the schema of the PostResponseSuccess rule.
func (prs PostResponseSuccess) GetSchema() model.RuleFunctionSchema {
	return model.RuleFunctionSchema{Name: "operation_response_success"}
}

// RunRule will execute the PostResponseSuccess rule, based on supplied context and a supplied []*yaml.Node slice.
func (prs PostResponseSuccess) RunRule(nodes []*yaml.Node, context model.RuleFunctionContext) []model.RuleFunctionResult {

	if len(nodes) <= 0 {
		return nil
	}

	props := utils.ExtractValueFromInterfaceMap("properties", context.Options)
	values := utils.ConvertInterfaceArrayToStringArray(props)
	found := 0
	var results []model.RuleFunctionResult

	for j, node := range nodes {

		var startNode, endNode *yaml.Node

		for _, propVal := range values {
			key, _ := utils.FindFirstKeyNode(propVal, []*yaml.Node{node}, 0)
			if key != nil {
				found++
			} else {
				startNode = node
				endNode = utils.FindLastChildNodeWithLevel(startNode, 0)
				if j+1 < len(nodes) {
					endNode = nodes[j+1]
				}
			}
		}

		if found <= 0 {

			if j+1 < len(node.Content) {
				endNode = node.Content[j+1]
			}

			results = append(results, model.RuleFunctionResult{
				Message: fmt.Sprintf("operations must define a success response with one of the following codes: '%s'",
					strings.Join(values, ", ")),
				StartNode: startNode,
				EndNode:   endNode,
				Rule:      context.Rule,
			})
		}

	}

	return results
}
