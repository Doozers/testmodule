package testmodule

import (
	"context"
)

type Service struct {
	UnimplementedTestmoduleSvcServer
}

func (Service) TheTest(ctx context.Context, req *TheTestReq) (*TheTestRes, error) {
	if b := req.GetIsTest(); b {
		return &TheTestRes{Res: &TheTestRes_Res{
			Response:   "test",
			Statuscode: 200,
		}}, nil
	}

	return &TheTestRes{Res: &TheTestRes_Res{
		Response:   "not test",
		Statuscode: 400,
	}}, nil
}

func (Service) SpamMe(ctx context.Context, empty *Empty) (*SpamMeRes, error) {
	return &SpamMeRes{Spam: "spam"}, nil
}
