package execution

type JoinCondition struct {
	LeftField  string
	RightField string
}

func resolveJoinField(ctx ExecutionContext, field string) string {
	return resolveField(ctx, field)
}

func ApplyJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	for _, leftCtx := range left {
		leftValue := resolveJoinField(leftCtx, condition.LeftField)
		matched := false

		for _, rightCtx := range right {
			rightValue := resolveJoinField(rightCtx, condition.RightField)

			if leftValue == rightValue {
				matched = true

				joined := leftCtx

				joined.Attributes = make(map[string]string)

				for key, value := range leftCtx.Attributes {
					joined.Attributes[key] = value
				}

				for key, value := range rightCtx.Attributes {
					joined.Attributes["right_"+key] = value
				}

				results = append(results, joined)
			}
		}

		if !matched {
			results = append(results, leftCtx)
		}
	}

	return results
}

func ApplyRightJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	for _, rightCtx := range right {
		rightValue := resolveJoinField(rightCtx, condition.RightField)
		matched := false

		for _, leftCtx := range left {
			leftValue := resolveJoinField(leftCtx, condition.LeftField)

			if rightValue == leftValue {
				matched = true

				joined := rightCtx

				joined.Attributes = make(map[string]string)

				for key, value := range rightCtx.Attributes {
					joined.Attributes[key] = value
				}

				for key, value := range leftCtx.Attributes {
					joined.Attributes["left_"+key] = value
				}

				results = append(results, joined)
			}
		}

		if !matched {
			results = append(results, rightCtx)
		}
	}

	return results
}

func ApplyFullOuterJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	matchedRight := make(map[int]bool)

	for _, leftCtx := range left {
		leftValue := resolveJoinField(leftCtx, condition.LeftField)
		matched := false

		for rightIndex, rightCtx := range right {
			rightValue := resolveJoinField(rightCtx, condition.RightField)

			if leftValue == rightValue {
				matched = true
				matchedRight[rightIndex] = true

				joined := leftCtx

				joined.Attributes = make(map[string]string)

				for key, value := range leftCtx.Attributes {
					joined.Attributes[key] = value
				}

				for key, value := range rightCtx.Attributes {
					joined.Attributes["right_"+key] = value
				}

				results = append(results, joined)
			}
		}

		if !matched {
			results = append(results, leftCtx)
		}
	}

	for rightIndex, rightCtx := range right {
		if !matchedRight[rightIndex] {
			results = append(results, rightCtx)
		}
	}

	return results
}

func ApplyCrossJoin(
	left []ExecutionContext,
	right []ExecutionContext,
) []ExecutionContext {
	var results []ExecutionContext

	for _, leftCtx := range left {
		for _, rightCtx := range right {
			joined := leftCtx

			joined.Attributes = make(map[string]string)

			for key, value := range leftCtx.Attributes {
				joined.Attributes[key] = value
			}

			for key, value := range rightCtx.Attributes {
				joined.Attributes["right_"+key] = value
			}

			results = append(results, joined)
		}
	}

	return results
}

func ApplySelfJoin(
	contexts []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	return ApplyJoin(contexts, contexts, condition)
}

func ApplyNaturalJoin(
	left []ExecutionContext,
	right []ExecutionContext,
) []ExecutionContext {
	var results []ExecutionContext

	for _, leftCtx := range left {
		for _, rightCtx := range right {
			matches := true
			hasCommonField := false

			for key, leftValue := range leftCtx.Attributes {
				rightValue, exists := rightCtx.Attributes[key]

				if !exists {
					continue
				}

				hasCommonField = true

				if leftValue != rightValue {
					matches = false
					break
				}
			}

			if !hasCommonField {
				matches = true
			}

			if matches {
				joined := leftCtx

				joined.Attributes = make(map[string]string)

				for key, value := range leftCtx.Attributes {
					joined.Attributes[key] = value
				}

				for key, value := range rightCtx.Attributes {
					joined.Attributes["right_"+key] = value
				}

				results = append(results, joined)
			}
		}
	}

	return results
}

// ApplyLeftAntiJoin returns left contexts that have no match in right.
func ApplyLeftAntiJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext
	for _, leftCtx := range left {
		leftValue := resolveJoinField(leftCtx, condition.LeftField)
		matched := false
		for _, rightCtx := range right {
			rightValue := resolveJoinField(rightCtx, condition.RightField)
			if leftValue == rightValue {
				matched = true
				break
			}
		}
		if !matched {
			results = append(results, leftCtx)
		}
	}
	return results
}

// ApplyRightAntiJoin returns right contexts that have no match in left.
func ApplyRightAntiJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	for _, rightCtx := range right {
		rightValue := resolveJoinField(rightCtx, condition.RightField)
		matched := false

		for _, leftCtx := range left {
			leftValue := resolveJoinField(leftCtx, condition.LeftField)

			if rightValue == leftValue {
				matched = true
				break
			}
		}

		if !matched {
			results = append(results, rightCtx)
		}
	}

	return results
}

// ApplySemiJoin returns left contexts that have at least one match in right.
func ApplySemiJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	for _, leftCtx := range left {
		leftValue := resolveJoinField(leftCtx, condition.LeftField)
		matched := false

		for _, rightCtx := range right {
			rightValue := resolveJoinField(rightCtx, condition.RightField)

			if leftValue == rightValue {
				matched = true
				break
			}
		}

		if matched {
			results = append(results, leftCtx)
		}
	}

	return results
}
