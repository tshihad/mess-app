package data

import (
	"mess-app/internal/core"
	"mess-app/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertGroup(t *testing.T) {
	type args struct {
		group    models.GroupPayload
		authorID int
	}
	testCases := []struct {
		desc    string
		args    args
		given   func(mock sqlmock.Sqlmock)
		want    error
		wantErr bool
	}{
		{
			desc: "Normal test case - 1",
			args: args{
				group: models.GroupPayload{
					GroupName: returnReferences("Sample Group"),
				},
				authorID: 1201,
			},
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("(.+)").WithArgs("Sample Group", 1201, sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewResult(1, 1),
				)
			},
			want:    nil,
			wantErr: false,
		},
	}
	db, mock := core.NewMock(t)
	r := &repo{
		DB: db,
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.given(mock)
			err := r.InsertGroup(tC.args.group, tC.args.authorID)
			if (err != nil) != tC.wantErr {
				t.Error("Failed to insert group ", err)
			}
		})
	}
}

func TestInsertUserToGroup(t *testing.T) {
	type args struct {
		group models.GroupPayload
		users []models.UserPayload
	}
	testCases := []struct {
		desc    string
		args    args
		given   func(mock sqlmock.Sqlmock)
		want    error
		wantErr bool
	}{
		{
			desc: "Normal test case",
			args: args{
				group: models.GroupPayload{
					GroupName: returnReferences("Test Group"),
				},
				users: []models.UserPayload{
					{
						UserName: returnReferences("Test User"),
					},
				},
			},
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("(.+)").WithArgs(sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewResult(1001, 1),
				)
			},
			want:    nil,
			wantErr: false,
		},
	}
	db, mock := core.NewMock(t)
	r := &repo{
		DB: db,
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tC.given(mock)
			err := r.InsertUserToGroup(tC.args.group, tC.args.users)
			if (err != nil) != tC.wantErr {
				t.Error(err, "Failed the test case")
			}
		})
	}
}

func returnReferences(str string) *string {
	return &str
}
