package snowflake_id

import (
	"fmt"
	"testing"
)

func TestGenerator_new(t *testing.T) {
	type fields struct {
		nodeId int64
	}
	type args struct {
		nowTimestamp int64
		randomNum    int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ID
	}{
		{
			name: "case 1",
			fields: fields{
				nodeId: 341,
			},
			args: args{
				nowTimestamp: 23456248059221,
				randomNum:    1365,
			},
			want: 6148914691236517205,
		},
		{
			name: "case 2",
			fields: fields{
				nodeId: 682,
			},
			args: args{
				nowTimestamp: 733007751850,
				randomNum:    2730,
			},
			want: 3074457345618258602,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				nodeId: tt.fields.nodeId,
			}
			got := g.new(tt.args.nowTimestamp, tt.args.randomNum)
			fmt.Printf("got result %d: %s\n", got, got.Base2())
			if got != tt.want {
				t.Errorf("new() = %v, want %v", got, tt.want)
			}
		})
	}
}
