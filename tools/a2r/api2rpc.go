// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package a2r

import (
	"context"
	"go-micro-clean/common"
	"go-micro-clean/tools/checker"
	"go-micro-clean/tools/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Call[A, B, C any](
	rpc func(client C, ctx context.Context, req *A, options ...grpc.CallOption) (*B, error),
	client C,
	c *gin.Context,
) {
	var req A
	if err := c.BindJSON(&req); err != nil {
		log.Error("gin bind json error", err, "req", req)
		// c.JSON(http.StatusBadRequest, err)
		common.WriteErrorResponse(c, common.ErrInvalidRequest(err))
		// apiresp.GinError(c, errs.ErrArgs.WithDetail(err.Error()).Wrap()) // 参数错误
		return
	}
	if err := checker.Validate(&req); err != nil {
		// c.JSON(http.StatusBadRequest, err)
		common.WriteErrorResponse(c, common.ErrInvalidRequest(err))
		// apiresp.GinError(c, err) // 参数校验失败
		return
	}
	data, err := rpc(client, c, &req)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, err)
		common.WriteErrorResponse(c, common.ErrInternal(err))
		// apiresp.GinError(c, err) // RPC调用失败
		return
	}
	c.JSON(http.StatusOK, data)
	// apiresp.GinSuccess(c, data) // 成功
}
