package execution

import (
	"fmt"
	"time"
	"sort"
)

type TraceNode struct {
	Context  ExecutionContext
	Children []*TraceNode
}

func BuildTraceTree(spans []ExecutionContext) []*TraceNode {
	nodeMap := make(map[string]*TraceNode)

	for _, s := range spans {
		nodeMap[s.SpanID] = &TraceNode{
			Context:  s,
			Children: []*TraceNode{},
		}
	}

	var roots []*TraceNode

	for _, node := range nodeMap {
		parentID := node.Context.ParentSpanID

		if parentID == "" {
			roots = append(roots, node)
			continue
		}

		parent, exists := nodeMap[parentID]
		if !exists {
			roots = append(roots, node)
			continue
		}

		parent.Children = append(parent.Children, node)
	}

	return roots
}

// PUBLIC ENTRY: call this to print tree
func PrintTraceTree(spans []ExecutionContext) {
	roots := BuildTraceTree(spans)

	fmt.Println("\nTRACE TREE")

	for _, r := range roots {
		printNode(r, "", true)
	}
}

// recursive renderer
func printNode(n *TraceNode, prefix string, isLast bool) {
	connector := "├── "
	nextPrefix := prefix + "│   "

	if isLast {
		connector = "└── "
		nextPrefix = prefix + "    "
	}

	duration := n.Context.Duration()

	fmt.Printf("%s%s%s (%s) [%s]\n",
		prefix,
		connector,
		n.Context.SpanName,
		n.Context.SpanID,
		duration,
	)

	if len(n.Context.Attributes) > 0 {

	keys := make([]string, 0, len(n.Context.Attributes))

	for k := range n.Context.Attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s    • %s = %s\n", nextPrefix, k, n.Context.Attributes[k])
	}
	}

	if len(n.Context.Lifecycle) > 0 {
		for _, event := range n.Context.Lifecycle {
			fmt.Printf(
				"%s    ◦ %s (%s)\n",
				nextPrefix,
				event.Name,
				event.Timestamp.Format(time.RFC3339),
			)
		}
	}

	for i, c := range n.Children {
		printNode(c, nextPrefix, i == len(n.Children)-1)
	}
}
