package src

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	type args struct {
		c         MySQLConfig
		db        string
		tableHint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get struct definition successfully",
			args: args{
				c: MySQLConfig{
					User:     "root",
					Host:     "127.0.0.1",
					Port:     3306,
					Password: "",
					Database: "information_schema",
					Encoding: "utf8mb4",
				},
				db:        "test_db",
				tableHint: "test",
			},
			want: `type Test struct {
	ID uint64
	TinyInt *int8
	SmallInt *int16
	VarcharString *string
	TextString *string
	LongtextString *string
	CreatedAt time.Time
}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tables, err := Parse(tt.args.c, tt.args.db, tt.args.tableHint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := GetStructDefinitionString(tables)
			if diff := cmp.Diff(tt.want, strings.TrimSpace(got)); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
