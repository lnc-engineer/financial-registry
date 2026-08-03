package execution

import "strings"

type DistinctQuery struct {
	Enabled bool
}

func ApplyDistinct(
	records []ExecutionContext,
	projection []string,
) []ExecutionContext {
	if len(projection) == 0 {
		return records
	}

	seen := make(map[string]bool)
	var result []ExecutionContext

	for _, record := range records {
		key := buildDistinctKey(record, projection)

		if seen[key] {
			continue
		}

		seen[key] = true
		result = append(result, record)
	}

	return result
}

func buildDistinctKey(
	ctx ExecutionContext,
	projection []string,
) string {
	values := make([]string, 0, len(projection))

	for _, field := range projection {
		value := resolveField(ctx, field)
		values = append(values, value)
	}

	return strings.Join(values, "|")
}
