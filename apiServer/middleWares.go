package apiServer

import (
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"strconv"
	"context"
	"encoding/json"
	"net/http"
)

func (server *Server) CheckAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var (
			requestParams map[string]string
			token string
			uid int
			isExist, isLogged bool
			err error
		)
		err = json.NewDecoder(r.Body).Decode(&requestParams)
		if err != nil {
			server.LogError(r, "request json decode failed - "+err.Error())
			server.error(w, errDef.InvalidRequestBody)
			return
		}
		token, isExist = requestParams["x-auth-token"]
		if !isExist {
			server.LogWarning(r, "x-auth-token not exist in request")
			server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
			return
		}

		// token, ok = item.(string)
		// if !ok {
		// 	server.LogWarning(r, "x-auth-token has wrong type")
		// 	server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		// 	return
		// }

		if token == "" {
			server.LogWarning(r, "x-auth-token is empty")
			server.error(w, errDef.UserNotLogged)
			return
		}

		uid, err = handlers.TokenUidDecode(token)
		if err != nil {
			server.LogWarning(r, "TokenUidDecode returned error - "+err.Error())
			server.error(w, errDef.UserNotLogged)
			return
		}

		isLogged = server.session.IsUserLoggedByUid(uid)
		if !isLogged {
			server.LogWarning(r, "User #"+strconv.Itoa(uid)+" is not logged")
			server.error(w, errDef.UserNotLogged)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) PanicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if ok {
					server.LogError(r, "PANIC happened - "+err.Error())
				} else { 
					server.LogError(r, "PANIC happened with unknown type of error")
				}
				server.error(w, errDef.UnknownInternalError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (server *Server) POSTMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method == "OPTIONS" {
			server.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "POST" {
			server.LogWarning(r, "wrong request method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Log(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}

func (server *Server) PATCHMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method == "OPTIONS" {
			server.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "PATCH" {
			server.LogWarning(r, "wrong request method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Log(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}

func (server *Server) DELETEMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method == "OPTIONS" {
			server.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "DELETE" {
			server.LogWarning(r, "wrong request method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Log(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}

func (server *Server) GETMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Method == "OPTIONS" {
			server.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "GET" {
			server.LogWarning(r, "wrong request method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Log(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}
