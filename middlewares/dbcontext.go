package middlewares

import (
	"context"
	"net/http"

	"main.go/services/dbconn"
)

const (
	db_instance string = "db_instance"
	db_ctx      string = "db_ctx"
)

func DBcontext(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), db_instance, dbconn.DB_Instance)
		ctx = context.WithValue(ctx, db_ctx, dbconn.DB_Ctx)

		// fmt.Println(ctx.Value(db_instance))
		// fmt.Println(ctx.Value(db_ctx))

		next(w, r.WithContext(ctx))
	})
}
