package execution

type JoinPredicate struct {
	LeftField  string
	RightField string
}

type JoinCondition struct {
	// LeftField and RightField preserve compatibility with single-condition joins.
	LeftField  string
	RightField string

	// Conditions supports multi-condition joins.
	Conditions []JoinPredicate
}

func resolveJoinField(ctx ExecutionContext, field string) string {
	return resolveField(ctx, field)
}

func joinPredicates(condition JoinCondition) []JoinPredicate {
	if len(condition.Conditions) > 0 {
		return condition.Conditions
	}

	return []JoinPredicate{
		{
			LeftField:  condition.LeftField,
			RightField: condition.RightField,
		},
	}
}

func joinContextsMatch(
	leftCtx ExecutionContext,
	rightCtx ExecutionContext,
	condition JoinCondition,
) bool {
	for _, predicate := range joinPredicates(condition) {
		leftValue := resolveJoinField(leftCtx, predicate.LeftField)
		rightValue := resolveJoinField(rightCtx, predicate.RightField)

		if leftValue != rightValue {
			return false
		}
	}

	return true
}

func joinContexts(leftCtx, rightCtx ExecutionContext) ExecutionContext {
	joined := leftCtx
	joined.Attributes = make(map[string]string)

	for key, value := range leftCtx.Attributes {
		joined.Attributes[key] = value
	}

	for key, value := range rightCtx.Attributes {
		joined.Attributes["right_"+key] = value
	}

	return joined
}

func joinContextsRight(leftCtx, rightCtx ExecutionContext) ExecutionContext {
	joined := rightCtx
	joined.Attributes = make(map[string]string)

	for key, value := range rightCtx.Attributes {
		joined.Attributes[key] = value
	}

	for key, value := range leftCtx.Attributes {
		joined.Attributes["left_"+key] = value
	}

	return joined
}

func ApplyJoin(
	left []ExecutionContext,
	right []ExecutionContext,
	condition JoinCondition,
) []ExecutionContext {
	var results []ExecutionContext

	for _, leftCtx := range left {
		matched := false

		for _, rightCtx := range right {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
				matched = true
				results = append(results, joinContexts(leftCtx, rightCtx))
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
		matched := false

		for _, leftCtx := range left {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
				matched = true
				results = append(results, joinContextsRight(leftCtx, rightCtx))
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
		matched := false

		for rightIndex, rightCtx := range right {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
				matched = true
				matchedRight[rightIndex] = true

				results = append(results, joinContexts(leftCtx, rightCtx))
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
			results = append(results, joinContexts(leftCtx, rightCtx))
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
				results = append(results, joinContexts(leftCtx, rightCtx))
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
		matched := false

		for _, rightCtx := range right {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
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
		matched := false

		for _, leftCtx := range left {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
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
		matched := false

		for _, rightCtx := range right {
			if joinContextsMatch(leftCtx, rightCtx, condition) {
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
