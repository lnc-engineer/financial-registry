package execution

import "testing"

func TestFindSpansByAttribute(t *testing.T) {
	root := &TraceNode{
		Context: ExecutionContext{
			SpanName:   "root",
			Attributes: map[string]string{"status": "success"},
		},
	}

	child1 := &TraceNode{
		Context: ExecutionContext{
			SpanName:   "child-success",
			Attributes: map[string]string{"status": "success"},
		},
	}

	child2 := &TraceNode{
		Context: ExecutionContext{
			SpanName:   "child-failure",
			Attributes: map[string]string{"status": "failure"},
		},
	}

	root.Children = []*TraceNode{child1, child2}

	matches := FindSpans(root, "status", "success")

	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}

func TestSpanNames(t *testing.T) {
	nodes := []*TraceNode{
		{
			Context: ExecutionContext{
				SpanName: "first",
			},
		},
		{
			Context: ExecutionContext{
				SpanName: "second",
			},
		},
	}

	names := SpanNames(nodes)

	if len(names) != 2 {
		t.Fatalf("expected 2 names")
	}

	if names[0] != "first" {
		t.Fatalf("unexpected first name")
	}

	if names[1] != "second" {
		t.Fatalf("unexpected second name")
	}
}

func TestSearchTrace(t *testing.T) {
	root := &TraceNode{
		Context: ExecutionContext{
			SpanName:   "root",
			Attributes: map[string]string{"status": "success"},
		},
	}

	child := &TraceNode{
		Context: ExecutionContext{
			SpanName:   "worker",
			Attributes: map[string]string{"status": "failure"},
		},
	}

	root.Children = []*TraceNode{child}

	names := SearchTrace(root, "status", "failure")

	if len(names) != 1 {
		t.Fatalf("expected 1 match, got %d", len(names))
	}

	if names[0] != "worker" {
		t.Fatalf("expected worker, got %s", names[0])
	}
}
