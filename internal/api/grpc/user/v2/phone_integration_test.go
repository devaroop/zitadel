//go:build integration

package user_test

import (
	"testing"

	"github.com/muhlemmer/gu"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zitadel/zitadel/internal/integration"
	object "github.com/zitadel/zitadel/pkg/grpc/object/v2alpha"
	user "github.com/zitadel/zitadel/pkg/grpc/user/v2alpha"
)

func TestServer_SetPhone(t *testing.T) {
	userID := Tester.CreateHumanUser(CTX).GetUserId()

	tests := []struct {
		name    string
		req     *user.SetPhoneRequest
		want    *user.SetPhoneResponse
		wantErr bool
	}{
		{
			name: "default verification",
			req: &user.SetPhoneRequest{
				UserId: userID,
				Phone:  "+41791234568",
			},
			want: &user.SetPhoneResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Tester.Organisation.ID,
				},
			},
		},
		{
			name: "send verification",
			req: &user.SetPhoneRequest{
				UserId: userID,
				Phone:  "+41791234569",
				Verification: &user.SetPhoneRequest_SendCode{
					SendCode: &user.SendPhoneVerificationCode{},
				},
			},
			want: &user.SetPhoneResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Tester.Organisation.ID,
				},
			},
		},
		{
			name: "return code",
			req: &user.SetPhoneRequest{
				UserId: userID,
				Phone:  "+41791234566",
				Verification: &user.SetPhoneRequest_ReturnCode{
					ReturnCode: &user.ReturnPhoneVerificationCode{},
				},
			},
			want: &user.SetPhoneResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Tester.Organisation.ID,
				},
				VerificationCode: gu.Ptr("xxx"),
			},
		},
		{
			name: "is verified true",
			req: &user.SetPhoneRequest{
				UserId: userID,
				Phone:  "+41791234565",
				Verification: &user.SetPhoneRequest_IsVerified{
					IsVerified: true,
				},
			},
			want: &user.SetPhoneResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Tester.Organisation.ID,
				},
			},
		},
		{
			name: "is verified false",
			req: &user.SetPhoneRequest{
				UserId: userID,
				Phone:  "+41791234564",
				Verification: &user.SetPhoneRequest_IsVerified{
					IsVerified: false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.SetPhone(CTX, tt.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			integration.AssertDetails(t, tt.want, got)
			if tt.want.GetVerificationCode() != "" {
				assert.NotEmpty(t, got.GetVerificationCode())
			}
		})
	}
}

func TestServer_VerifyPhone(t *testing.T) {
	userResp := Tester.CreateHumanUser(CTX)
	tests := []struct {
		name    string
		req     *user.VerifyPhoneRequest
		want    *user.VerifyPhoneResponse
		wantErr bool
	}{
		{
			name: "wrong code",
			req: &user.VerifyPhoneRequest{
				UserId:           userResp.GetUserId(),
				VerificationCode: "xxx",
			},
			wantErr: true,
		},
		{
			name: "wrong user",
			req: &user.VerifyPhoneRequest{
				UserId:           "xxx",
				VerificationCode: userResp.GetPhoneCode(),
			},
			wantErr: true,
		},
		{
			name: "verify user",
			req: &user.VerifyPhoneRequest{
				UserId:           userResp.GetUserId(),
				VerificationCode: userResp.GetPhoneCode(),
			},
			want: &user.VerifyPhoneResponse{
				Details: &object.Details{
					Sequence:      1,
					ChangeDate:    timestamppb.Now(),
					ResourceOwner: Tester.Organisation.ID,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Client.VerifyPhone(CTX, tt.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			integration.AssertDetails(t, tt.want, got)
		})
	}
}
