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

		for _, rightCtx := range right {
			rightValue := resolveJoinField(rightCtx, condition.RightField)

			if leftValue == rightValue {
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
	}

	return results
}
