package execution

import (
	"testing"
	"time"
)

func TestSortExecutionContextsBySpanNameAscending(t *testing.T) {
	contexts := []ExecutionContext{
		{SpanName: "payment"},
		{SpanName: "auth"},
		{SpanName: "validation"},
	}

	SortExecutionContexts(contexts, QuerySort{
		Field: SortBySpanName,
		Order: SortAscending,
	})

	expected := []string{"auth", "payment", "validation"}

	for i, name := range expected {
		if contexts[i].SpanName != name {
			t.Fatalf("expected %s, got %s", name, contexts[i].SpanName)
		}
	}
}

func TestSortExecutionContextsBySpanNameDescending(t *testing.T) {
	contexts := []ExecutionContext{
		{SpanName: "auth"},
		{SpanName: "validation"},
		{SpanName: "payment"},
	}

	SortExecutionContexts(contexts, QuerySort{
		Field: SortBySpanName,
		Order: SortDescending,
	})

	expected := []string{"validation", "payment", "auth"}

	for i, name := range expected {
		if contexts[i].SpanName != name {
			t.Fatalf("expected %s, got %s", name, contexts[i].SpanName)
		}
	}
}

func TestSortExecutionContextsByStartTime(t *testing.T) {
	now := time.Now()

	contexts := []ExecutionContext{
		{StartTime: now.Add(2 * time.Hour)},
		{StartTime: now},
		{StartTime: now.Add(1 * time.Hour)},
	}

	SortExecutionContexts(contexts, QuerySort{
		Field: SortByStartTime,
		Order: SortAscending,
	})

	if !contexts[0].StartTime.Equal(now) {
		t.Fatal("expected earliest start time first")
	}
}

func TestSortExecutionContextsByEndTime(t *testing.T) {
	now := time.Now()

	contexts := []ExecutionContext{
		{EndTime: now.Add(3 * time.Hour)},
		{EndTime: now},
		{EndTime: now.Add(1 * time.Hour)},
	}

	SortExecutionContexts(contexts, QuerySort{
		Field: SortByEndTime,
		Order: SortAscending,
	})

	if !contexts[0].EndTime.Equal(now) {
		t.Fatal("expected earliest end time first")
	}
}
