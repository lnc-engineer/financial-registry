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

				if joined.Attributes == nil {
					joined.Attributes = map[string]string{}
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
