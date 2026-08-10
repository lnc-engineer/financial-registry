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
