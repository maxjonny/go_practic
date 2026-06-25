package transport

import (
	"encoding/json"
	rep "main/internal/repository"
	handler "main/internal/requests"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
	rep rep.RepositoryInterface
}

func InitRouter(storageInterface rep.RepositoryInterface) *Router {
	router := Router{
		mux: http.NewServeMux(),
		rep: storageInterface,
	}
	router.SetupRoutes()
	return &router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) GetUserCount(w http.ResponseWriter, req *http.Request) {
	handler.GetUserCount(w, req, r.rep)
}

func (r *Router) GetUserData(w http.ResponseWriter, req *http.Request) {
	handler.GetUserData(w, req, r.rep)
}

func (r *Router) SetupRoutes() {

	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserCount", r.GetUserCount)
	r.mux.HandleFunc("/checkbox/Z5/{device}/actionapi/User/GetUserData/{index}", r.GetUserData)

	r.mux.HandleFunc("/checkbox/Z5/{device}/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w)
	})
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
