package main

import (
	_ "embed"
	"testing"
	"time"
)

func TestChatMessage_Validate(t *testing.T) {
	users = Users{"test", "t", "Tester mcTesterson the third"}
	type fields struct {
		Username  string
		Text      string
		Time      string
		Timestamp time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"valid message", fields{Username: "test", Text: "test"}, false},
		{"username too short", fields{Username: "t", Text: "test"}, true},
		{"username too long", fields{Username: "Tester mcTesterson the third"}, true},
		{"username invalid", fields{Username: "invalid", Text: "test"}, true},
		{"invalid text", fields{Username: "test", Text: "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatMessage{
				Username:  tt.fields.Username,
				Text:      tt.fields.Text,
				Time:      tt.fields.Time,
				Timestamp: tt.fields.Timestamp,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ChatMessage.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChatMessage_FromJson(t *testing.T) {
	type fields struct {
		Username  string
		Text      string
		Time      string
		Timestamp time.Time
	}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid message", fields{Username: "", Text: ""}, args{in: `{"username":"test","text":"test"}`}, false},
		{"invalid message", fields{Username: "", Text: ""}, args{in: `{"username":"test","text":"test}`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatMessage{
				Username:  tt.fields.Username,
				Text:      tt.fields.Text,
				Time:      tt.fields.Time,
				Timestamp: tt.fields.Timestamp,
			}
			if err := c.FromJson(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("ChatMessage.FromJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
