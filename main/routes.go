import (
    "github.com/gorilla/mux"
    "net/http"
)
// Struct for a Route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
// Array of Route
type Routes []Route
// initiate the router
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    // declare routes
    var routes = Routes{
        Route{
            "getIndex",
            "GET",
            "/",
            getIndex,
        },
        Route{
            "getJson",
            "GET",
            "/json",
            getJson,
        },
    }
    // bind routes
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}