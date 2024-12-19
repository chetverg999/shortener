package handlers

import "testing"

func Test_parseURL(t *testing.T) {
	tests := []struct {
		name    string
		args    []byte
		wantErr bool
	}{
		{"1", []byte("https://ru.stackoverflow.com/"), false},
		{"2", []byte("http://example.com"), false},
		{"3", []byte("ftp://example.com"), true},
		{"4", []byte("ttps://ru.stackoverflow.com/"), true},
		{"5", []byte("://missing.scheme.com"), true},
		{"6", []byte("invalidurl"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseURL(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
