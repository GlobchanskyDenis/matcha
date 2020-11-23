package apiServer

import (
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (server *Server) CheckAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams         map[string]interface{}
			item                  interface{}
			ctx                   context.Context
			token                 string
			uid                   int
			isExist, isLogged, ok bool
			err                   error
		)
		ctx = r.Context()
		requestParams = ctx.Value("requestParams").(map[string]interface{})

		item, isExist = requestParams["x-auth-token"]
		if !isExist {
			server.Logger.LogWarning(r, "x-auth-token not exist in request")
			server.error(w, errors.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
			return
		}

		token, ok = item.(string)
		if !ok {
			server.Logger.LogWarning(r, "x-auth-token has wrong type")
			server.error(w, errors.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
			return
		}

		if token == "" {
			server.Logger.LogWarning(r, "x-auth-token is empty")
			server.error(w, errors.UserNotLogged)
			return
		}

		uid, err = handlers.TokenUidDecode(token)
		if err != nil {
			server.Logger.LogWarning(r, "TokenUidDecode returned error - "+err.Error())
			server.error(w, errors.UserNotLogged)
			return
		}

		isLogged = server.Session.IsUserLoggedByUid(uid)
		if !isLogged {
			server.Logger.LogWarning(r, "User #"+strconv.Itoa(uid)+" is not logged")
			server.error(w, errors.UserNotLogged)
			return
		}

		ctx = context.WithValue(r.Context(), "uid", uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) PanicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if ok {
					server.Logger.LogError(r, "PANIC happened - "+err.Error())
				} else {
					server.Logger.LogError(r, "PANIC happened with unknown type of error")
				}
				server.error(w, errors.UnknownInternalError)
				return
			}
		}()
		next.ServeHTTP(w, r)
		server.Logger.TimeLog(r, time.Since(t))
	})
}

func (server *Server) PutMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams map[string]interface{}
			err           error
			ctx           context.Context
		)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PUT,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Content-Length")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "PUT" {
			server.Logger.LogWarning(r, "wrong request method. Should be PUT method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Logger.Log(r, "request from client was received")
		err = json.NewDecoder(r.Body).Decode(&requestParams)
		if err != nil {
			server.Logger.LogError(r, "request body json decode failed - "+err.Error())
			server.error(w, errors.InvalidRequestBody)
			return
		}
		ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) PostMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams map[string]interface{}
			err           error
			ctx           context.Context
		)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Content-Length")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "POST" {
			server.Logger.LogWarning(r, "wrong request method. Should be POST method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Logger.Log(r, "request from client was received")
		err = json.NewDecoder(r.Body).Decode(&requestParams)
		if err != nil {
			server.Logger.LogError(r, "request body json decode failed - "+err.Error())
			server.error(w, errors.InvalidRequestBody)
			return
		}
		ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) PatchMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams map[string]interface{}
			err           error
			ctx           context.Context
		)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Content-Length")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "PATCH" {
			server.Logger.LogWarning(r, "wrong request method. Should be PATCH method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Logger.Log(r, "request from client was received")
		err = json.NewDecoder(r.Body).Decode(&requestParams)
		if err != nil {
			server.Logger.LogError(r, "request body json decode failed - "+err.Error())
			server.error(w, errors.InvalidRequestBody)
			return
		}
		ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) PatchPostGetMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams = map[string]interface{}{}
			err           error
			ctx           context.Context
		)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Content-Length")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "PATCH" && r.Method != "POST" && r.Method != "GET" {
			server.Logger.LogWarning(r, "wrong request method. Should be PATCH or POST or GET method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}

		server.Logger.Log(r, "request from client was received")

		if r.Method == "PATCH" {
			err = json.NewDecoder(r.Body).Decode(&requestParams)
			if err != nil {
				server.Logger.LogError(r, "request body json decode failed - "+err.Error())
				server.error(w, errors.InvalidRequestBody)
				return
			}
			ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				server.Logger.LogError(r, "parse form failed - "+err.Error())
				server.error(w, errors.InvalidRequestBody)
				return
			}
			for key, value := range r.PostForm {
				requestParams[key] = value
			}
			ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		} else if r.Method == "GET" {
			ctx = r.Context()
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) DeleteMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestParams map[string]interface{}
			err           error
			ctx           context.Context
		)

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Content-Length")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "DELETE" {
			server.Logger.LogWarning(r, "wrong request method. Should be DELETE method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Logger.Log(r, "request from client was received")
		err = json.NewDecoder(r.Body).Decode(&requestParams)
		if err != nil {
			server.Logger.LogError(r, "request body json decode failed - "+err.Error())
			server.error(w, errors.InvalidRequestBody)
			return
		}
		ctx = context.WithValue(r.Context(), "requestParams", requestParams)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (server *Server) GetMethodMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			server.Logger.Log(r, "client wants to know what methods are allowed")
			return
		} else if r.Method != "GET" {
			server.Logger.LogWarning(r, "wrong request method. Should be GET method")
			w.WriteHeader(http.StatusMethodNotAllowed) // 405
			return
		}
		server.Logger.Log(r, "request from client was received")
		next.ServeHTTP(w, r)
	})
}
