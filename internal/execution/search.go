package execution

func FindSpans(root *TraceNode, key, value string) []*TraceNode {
	if root == nil {
		return nil
	}

	var matches []*TraceNode

	if root.Context.Attributes[key] == value {
		matches = append(matches, root)
	}

	for _, child := range root.Children {
		matches = append(matches, FindSpans(child, key, value)...)
	}

	return matches
}

func SpanNames(nodes []*TraceNode) []string {
	names := make([]string, 0, len(nodes))

	for _, node := range nodes {
		names = append(names, node.Context.SpanName)
	}

	return names
}
