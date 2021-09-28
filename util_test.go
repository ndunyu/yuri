package yuri

import "testing"

func Test_checkKenyaInternationalPhoneNumber(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test correct phone",
			args: args{
				phone: "+254798675644",
			},
			want: true,
		},

		{
			name: "test More numbers",
			args: args{
				phone: "+2547986756445",
			},
			want: false,
		},

		{
			name: "test 9 character phone",
			args: args{
				phone: "+25479867564",
			},
			want: false,
		},
		{
			name: "test 07 phone",
			args: args{
				phone: "0798675644",
			},
			want: false,
		},

		{
			name: "test Not international phone",
			args: args{
				phone: "25479867564",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkKenyaInternationalPhoneNumber(tt.args.phone); got != tt.want {
				t.Errorf("%s In checkKenyaInternationalPhoneNumber() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
