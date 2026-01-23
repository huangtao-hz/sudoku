package main

import "testing"

func TestSet(t *testing.T) {
	s := NewSet(1, 4, 7)
	if s.String() != "147" {
		t.Errorf("Test NewSet、String failed!  %s != 147", s.String())
	}
	s.Add(3, 5)
	if s.String() != "13457" {
		t.Errorf("Test Add failed!  %s != 13457", s.String())
	}
	s.Remove(3, 4)
	if s.String() != "157" {
		t.Errorf("Test Add failed!  %s != 157", s.String())
	}
	s = NewSet(1, 4, 6)
	if s.Count() != 3 {
		t.Errorf("Test Count Failed! %d != 3", s.Count())
	}
	if s.GetOnlyVlaue() != 0 {
		t.Errorf("Test GetOnleyValue Failed! %d != 0", s.GetOnlyVlaue())
	}
	s.Remove(1, 4)
	if s.GetOnlyVlaue() != 6 {
		t.Errorf("Test GetOnleyValue Failed! %d != 6", s.GetOnlyVlaue())
	}

}

func TestDiff(t *testing.T) {
	s := NewSet(1, 4, 7)
	if (s.Difference(NewSet(3, 4))).String() != "17" {
		t.Errorf("Test Diff failed!  %s != 17", s.String())
	}
	if (s.Union(NewSet(3, 4))).String() != "1347" {
		t.Errorf("Test Union failed!  %s != 1347", s.String())
	}
	if (s.Intersection(NewSet(3, 4))).String() != "4" {
		t.Errorf("Test Union failed!  %s != 4", s.String())
	}
	if !s.Contains(4) {
		t.Errorf("Test Contains failed! ")
	}
	if s.Contains(1, 4, 9) {
		t.Errorf("Test Contains failed! ")
	}
}
