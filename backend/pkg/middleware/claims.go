package middleware

import (
	"context"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
	"github.com/MnPutrav2/go_architecture/pkg/util"
)

func UserClaims(ctx context.Context) (*jwtEnc.Claims, bool) {
	ct, ok := ctx.Value(util.UserClaimsKey).(*jwtEnc.Claims)
	return ct, ok
}
