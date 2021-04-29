// Copyright 2018 The life-file Authors
// This file is part of the life-file library.
//
// The life-file library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The life-file library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the life-file library. If not, see <http://www.gnu.org/licenses/>.

package accounts

import (
	"testing"
)

func TestURLParsing(t *testing.T) {
	url, err := parseURL("https://lifefile.org")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "lifefile.org" {
		t.Errorf("expected: %v, got: %v", "lifefile.org", url.Path)
	}

	_, err = parseURL("lifefile.org")
	if err == nil {
		t.Error("expected err, got: nil")
	}
}

func TestURLString(t *testing.T) {
	url := URL{Scheme: "https", Path: "lifefile.org"}
	if url.String() != "https://lifefile.org" {
		t.Errorf("expected: %v, got: %v", "https://lifefile.org", url.String())
	}

	url = URL{Scheme: "", Path: "lifefile.org"}
	if url.String() != "lifefile.org" {
		t.Errorf("expected: %v, got: %v", "lifefile.org", url.String())
	}
}

func TestURLMarshalJSON(t *testing.T) {
	url := URL{Scheme: "https", Path: "lifefile.org"}
	json, err := url.MarshalJSON()
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if string(json) != "\"https://lifefile.org\"" {
		t.Errorf("expected: %v, got: %v", "\"https://lifefile.org\"", string(json))
	}
}

func TestURLUnmarshalJSON(t *testing.T) {
	url := &URL{}
	err := url.UnmarshalJSON([]byte("\"https://lifefile.org\""))
	if err != nil {
		t.Errorf("unexpcted error: %v", err)
	}
	if url.Scheme != "https" {
		t.Errorf("expected: %v, got: %v", "https", url.Scheme)
	}
	if url.Path != "lifefile.org" {
		t.Errorf("expected: %v, got: %v", "https", url.Path)
	}
}

func TestURLComparison(t *testing.T) {
	tests := []struct {
		urlA   URL
		urlB   URL
		expect int
	}{
		{URL{"https", "lifefile.org"}, URL{"https", "lifefile.org"}, 0},
		{URL{"http", "lifefile.org"}, URL{"https", "lifefile.org"}, -1},
		{URL{"https", "lifefile.org/a"}, URL{"https", "lifefile.org"}, 1},
		{URL{"https", "abc.org"}, URL{"https", "lifefile.org"}, -1},
	}

	for i, tt := range tests {
		result := tt.urlA.Cmp(tt.urlB)
		if result != tt.expect {
			t.Errorf("test %d: cmp mismatch: expected: %d, got: %d", i, tt.expect, result)
		}
	}
}
