package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dohaelsawy/bookStore/token"
	"github.com/dohaelsawy/bookStore/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)
func addAuthorization (
	t *testing.T,
	request *http.Request,
	tokenmaker token.Maker,
	authorizationType string,
	email string ,
	duration time.Duration,
){
	token ,err := tokenmaker.CreateToken(email , duration)
	require.NoError(t,err)

	authorizationHeader := fmt.Sprintf("%s %s" , authorizationType , token)

	request.Header.Add(authorizationHeaderKey,authorizationHeader)
}

func TestAuthenticationMiddleware(t *testing.T){
	email := util.RandomEmail()
	testcases := []struct{
		name string
		setupAuth func(t *testing.T , request *http.Request , tokenmaker token.Maker)
		checkResponse func(t *testing.T , recorder *httptest.ResponseRecorder)
	}{
		{
			name:"ok" ,
			setupAuth: func(t *testing.T, request *http.Request, tokenmaker token.Maker) {
				addAuthorization(t,request,tokenmaker,authorizationTypeBearer,email,time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t,http.StatusOK , recorder.Code)
			},
		},
	}



	for i := range testcases {
		tc := testcases[i]

		t.Run(tc.name , func (t *testing.T){
			server := newTestServer(t)
			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func (ctx *gin.Context)  {
					ctx.JSON(http.StatusOK , gin.H{})
				},
			)
			recorder := httptest.NewRecorder()
			request  , err := http.NewRequest(http.MethodGet,authPath,nil)
			require.NoError(t,err)

			tc.setupAuth(t,request,server.tokenMaker)
			server.router.ServeHTTP(recorder,request)
			tc.checkResponse(t,recorder)
		})
	}
}