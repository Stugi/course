package reader

import (
	"reflect"
	"testing"
)

func TestGetStringsFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				path: "/Users/arinastugireva/Documents/projects/learning/course/34_Regexp_Reader_Writer/task34.6.1/example.txt",
			},
			want:    []string{"5+4=?", "9+3=?", "13+7=?", "4-2=?"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStringsFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStringsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringsFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
