package transport

import (
	"context"
	"encoding/json"
	"github.com/lr2021/recruit-backend/user/model"
	"github.com/lr2021/recruit-backend/general/errors"
	"net/http"
)

func DecodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, errors.BadRequest("para:001", "missing or invalid parameters")
	}
	return req, nil
}
