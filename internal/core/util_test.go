package core

import "testing"

func TestHashPasswordAndCompare(t *testing.T) {
	testCases := []struct {
		title   string
		args    []byte
		want    error
		wantErr bool
	}{
		{
			title:   "1 - success case",
			args:    []byte("sample"),
			want:    nil,
			wantErr: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.title, func(t *testing.T) {
			output, err := HashPassword(tC.args)
			if err != nil {
				t.Error(err)
			}
			err = CompareHashAndPassword([]byte(output), tC.args)
			if (err != tC.want) != tC.wantErr {
				t.Errorf("Expected : %+v , got : %+v", tC.want, err)
			}
		})
	}
}
