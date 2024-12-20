package middleware

import "net/http"

// Chain combines multiple middleware into one
func Chain(middlewares ...func(http.HandlerFunc) http.HandlerFunc) func(http.HandlerFunc) http.HandlerFunc {
    return func(final http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            last := final
            for i := len(middlewares) - 1; i >= 0; i-- {
                last = middlewares[i](last)
            }
            last(w, r)
        }
    }
}
