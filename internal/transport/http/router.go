package http

import (
	"encoding/json"
	"fmt"
	rep "main/internal/repository"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func InitRouter(storageInterface rep.RepositoryInterface) *Router {
	r := Router{http.NewServeMux()}

	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserCount", func(w http.ResponseWriter, r *http.Request) { fmt.Println() })
	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserData/{index}", func(w http.ResponseWriter, r *http.Request) { fmt.Println() })

	r.mux.HandleFunc("/checkbox/Z5/{device}/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w)
	})
	return &r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// router.get("/actionapi/User/GetUserData", (req, res) => GetCardsReq.request(req, res))
// router.post("/actionapi/User/UploadAlcohol", (req, res) => AddCardEventReq.request(req, res))
// router.post("/actionapi/User/UploadUser", (req, res) => { res.send("success") }) // UploadCardReq.request(req, res)
// router.get("/actionapi/User/GetUserCount", (req, res) => GetUserCount.request(req, res))

// router.head('/', function(req, res) {
// 	res.status(200).send('Success');
// });

// router.get('*', function(req, res) {
// 	res.status(404).send({error: 'Not_Found'});
// });

// router.post('*', function(req, res) {
// 	res.status(404).send({error: 'Not_Found'});
// });

// router.put('*', function(req, res) {
// 	res.status(404).send({error: 'Not_Found'});
// });

// router.delete('*', function(req, res) {
// 	res.status(404).send({error: 'Not_Found'});
// });
