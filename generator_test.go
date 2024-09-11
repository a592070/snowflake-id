package snowflake

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGenerator_NewId(t *testing.T) {
	type fields struct {
		nodeId      int64
		timeNowFunc func() time.Time
		randNumFunc func(int64) int64
	}
	tests := []struct {
		name   string
		fields fields
		want   func(ID) bool
	}{
		{
			name: "case 1",
			fields: fields{
				nodeId:      341,
				timeNowFunc: time.Now,
				randNumFunc: rand.Int63n,
			},
			want: func(id ID) bool {
				return id.Int64() > 0
			},
		},
		{
			name: "case 2",
			fields: fields{
				nodeId: 341,
				timeNowFunc: func() time.Time {
					return time.UnixMilli(23456248059221)
				},
				randNumFunc: func(int64) int64 {
					return 1365
				},
			},
			want: func(id ID) bool {
				return id.Int64() == 6148914691236517205
			},
		},
		{
			name: "case 3",
			fields: fields{
				nodeId: 682,
				timeNowFunc: func() time.Time {
					return time.UnixMilli(733007751850)
				},
				randNumFunc: func(int64) int64 {
					return 2730
				},
			},
			want: func(id ID) bool {
				return id.Int64() == 3074457345618258602
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				nodeId:      tt.fields.nodeId,
				timeNowFunc: tt.fields.timeNowFunc,
				randNumFunc: tt.fields.randNumFunc,
			}
			got := g.NewId()
			fmt.Printf("got result %d: %s\n", got, got.Base2())
			if assert := tt.want(got); !assert {
				t.Errorf("NewId() = %v, assert %v", got, assert)
			}
		})
	}
}
