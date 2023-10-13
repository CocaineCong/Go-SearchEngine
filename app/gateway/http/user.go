package http

import (
	"go.opentelemetry.io/otel"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/tangseng/app/gateway/rpc"
	pb "github.com/CocaineCong/tangseng/idl/pb/user"
	"github.com/CocaineCong/tangseng/pkg/ctl"
	"github.com/CocaineCong/tangseng/pkg/jwt"
	log "github.com/CocaineCong/tangseng/pkg/logger"
	"github.com/CocaineCong/tangseng/types"
)

var tracer = otel.Tracer("tangseng-user")

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userReq pb.UserRegisterReq
	if err := ctx.ShouldBind(&userReq); err != nil {
		log.LogrusObj.Errorf("Bind:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "绑定参数错误"))
		return
	}
	r, err := rpc.UserRegister(ctx, &userReq)
	if err != nil {
		log.LogrusObj.Errorf("UserRegister:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserRegister RPC服务调用错误"))
		return
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, r))
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var req pb.UserLoginReq
	if err := ctx.ShouldBind(&req); err != nil {
		log.LogrusObj.Errorf("Bind:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "绑定参数错误"))
		return
	}

	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		log.LogrusObj.Errorf("RPC UserLogin:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "UserLogin RPC服务调用错误"))
		return
	}

	aToken, rToken, err := jwt.GenerateToken(userResp.UserDetail.UserId, userResp.UserDetail.UserName)
	if err != nil {
		log.LogrusObj.Errorf("RPC GenerateToken:%v", err)
		ctx.JSON(http.StatusInternalServerError, ctl.RespError(ctx, err, "加密错误"))
		return
	}
	uResp := &types.UserTokenData{
		User:         userResp.UserDetail,
		AccessToken:  aToken,
		RefreshToken: rToken,
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, uResp))
}
