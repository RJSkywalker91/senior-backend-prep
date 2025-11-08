package auth

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ctxKey string

const PlayerIDKey ctxKey = "playerID"

type Verifier struct{ signingKey []byte }

func NewVerifier(signingKey []byte) *Verifier { return &Verifier{signingKey: signingKey} }

func (v *Verifier) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if strings.Contains(info.FullMethod, "AuthService/Login") {
			return handler(ctx, req) // login is public
		}
		pid, err := v.playerIDFromMD(ctx)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, PlayerIDKey, pid)
		return handler(ctx, req)
	}
}

func (v *Verifier) Stream() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if strings.Contains(info.FullMethod, "AuthService/Login") {
			return handler(srv, ss)
		}
		md, _ := metadata.FromIncomingContext(ss.Context())
		pid, err := v.playerIDFromToken(md)
		if err != nil {
			return err
		}
		wrapped := &wrappedStream{ServerStream: ss, ctx: context.WithValue(ss.Context(), PlayerIDKey, pid)}
		return handler(srv, wrapped)
	}
}

func (v *Verifier) playerIDFromMD(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	return v.playerIDFromToken(md)
}

func (v *Verifier) playerIDFromToken(md metadata.MD) (string, error) {
	vals := md.Get("authorization")
	if len(vals) == 0 {
		return "", grpc.Errorf(16, "unauthenticated")
	} // codes.Unauthenticated
	parts := strings.SplitN(vals[0], " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", grpc.Errorf(16, "unauthenticated")
	}
	tokenStr := parts[1]
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, grpc.Errorf(16, "unauthenticated")
		}
		return v.signingKey, nil
	})
	if err != nil || !token.Valid {
		return "", grpc.Errorf(16, "unauthenticated")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", grpc.Errorf(16, "unauthenticated")
	}

	pid, _ := claims["sub"].(string) // store player_id as JWT subject
	if pid == "" {
		return "", grpc.Errorf(16, "unauthenticated")
	}
	return pid, nil
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context { return w.ctx }
