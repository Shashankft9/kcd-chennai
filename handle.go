package function

import (
        "context"
        "fmt"
        "net/http"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
        fmt.Println("received a request")
        fmt.Fprintln(res, "Hello from KCD Chennai! Read more about us here: https://community.cncf.io/events/details/cncf-kcd-chennai-presents-kubernetes-community-days-chennai-2022/") // Send KCD community link back to the client
}
