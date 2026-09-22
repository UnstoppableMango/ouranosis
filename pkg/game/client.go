// Package game wraps the unmango/game framework services in Go types.
package game

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/unmango/game/convert"
	"github.com/unmango/game/curve"
	gamev1alpha1 "github.com/unmango/game/gen/dev/unmango/game/v1alpha1"
	"github.com/unmango/game/gen/dev/unmango/game/v1alpha1/gamev1alpha1connect"
	"github.com/unmango/game/num"
	"github.com/unmango/game/seed"
	"google.golang.org/protobuf/types/known/durationpb"
)

// DefaultURL is where a local framework server listens.
const DefaultURL = "http://localhost:8080"

// Calculator is the subset of the framework the game uses.
// Client talks to a server; Local computes in process for tests.
type Calculator interface {
	Derive(ctx context.Context, path string) (uint64, error)
	Evaluate(ctx context.Context, c curve.Curve, n int64) (num.Number, error)
	Invert(ctx context.Context, c curve.Curve, from int64, budget num.Number) (int64, error)
	Advance(ctx context.Context, rate curve.Curve, level int64, elapsed time.Duration) (num.Number, error)
}

// Client calls a framework server over ConnectRPC.
type Client struct {
	calc gamev1alpha1connect.CalculatorServiceClient
}

var _ Calculator = (*Client)(nil)

// New returns a client for the server at baseURL.
func New(baseURL string) *Client {
	return &Client{calc: gamev1alpha1connect.NewCalculatorServiceClient(http.DefaultClient, baseURL)}
}

// Derive implements Calculator.
func (c *Client) Derive(ctx context.Context, path string) (uint64, error) {
	res, err := c.calc.Derive(ctx, connect.NewRequest(&gamev1alpha1.DeriveRequest{Path: path}))
	if err != nil {
		return 0, err
	}
	return res.Msg.GetSeed(), nil
}

// Evaluate implements Calculator.
func (c *Client) Evaluate(ctx context.Context, cv curve.Curve, n int64) (num.Number, error) {
	p, err := convert.CurveToProto(cv)
	if err != nil {
		return num.Zero(), err
	}
	res, err := c.calc.Evaluate(ctx, connect.NewRequest(&gamev1alpha1.EvaluateRequest{Curve: p, N: n}))
	if err != nil {
		return num.Zero(), err
	}
	return convert.NumberFromProto(res.Msg.GetValue()), nil
}

// Invert implements Calculator.
func (c *Client) Invert(ctx context.Context, cv curve.Curve, from int64, budget num.Number) (int64, error) {
	p, err := convert.CurveToProto(cv)
	if err != nil {
		return 0, err
	}
	res, err := c.calc.Invert(ctx, connect.NewRequest(&gamev1alpha1.InvertRequest{
		Curve: p, From: from, Budget: convert.NumberToProto(budget),
	}))
	if err != nil {
		return 0, err
	}
	return res.Msg.GetCount(), nil
}

// Advance implements Calculator.
func (c *Client) Advance(ctx context.Context, rate curve.Curve, level int64, elapsed time.Duration) (num.Number, error) {
	p, err := convert.CurveToProto(rate)
	if err != nil {
		return num.Zero(), err
	}
	res, err := c.calc.Advance(ctx, connect.NewRequest(&gamev1alpha1.AdvanceRequest{
		Rate: p, Level: level, Elapsed: durationpb.New(elapsed),
	}))
	if err != nil {
		return num.Zero(), err
	}
	return convert.NumberFromProto(res.Msg.GetValue()), nil
}

// Local computes with the framework's pure packages in process.
// It exists for tests and gives the same answers as a server with the same seed.
type Local struct {
	Seed int64
}

var _ Calculator = Local{}

// Derive implements Calculator.
func (l Local) Derive(_ context.Context, path string) (uint64, error) {
	return seed.Derive(l.Seed, path), nil
}

// Evaluate implements Calculator.
func (Local) Evaluate(_ context.Context, c curve.Curve, n int64) (num.Number, error) {
	return c.Evaluate(n), nil
}

// Invert implements Calculator.
func (Local) Invert(_ context.Context, c curve.Curve, from int64, budget num.Number) (int64, error) {
	return curve.Invert(c, from, budget), nil
}

// Advance implements Calculator.
func (Local) Advance(_ context.Context, rate curve.Curve, level int64, elapsed time.Duration) (num.Number, error) {
	return curve.Advance(rate, level, elapsed), nil
}
