# mu
mu is an AWS Lambda helper that allows for re-purposing HTTP API projects quickly into lambdas.

## Example

```
import (
	"net/http"

	"github.com/ModelRocket/mu"
	"github.com/gorilla/mux"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/products/{key}", productHandler)

	mu.Start(r)
}
```
