package serve

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/psanetra/http-db/store"
	"context"
	"github.com/go-http-utils/headers"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type HttpDbServer struct {
	config     *Config
	httpServer *http.Server
	router     *mux.Router
	store      store.Store
}

func NewServer(config *Config) (*HttpDbServer) {

	ret := &HttpDbServer{
		config: config,
		store:  store.NewInMemoryStore(),
	}

	ret.httpServer = &http.Server{
		Addr:        fmt.Sprintf("%s:%d", config.address, config.port),
		ReadTimeout: config.timeout,
		Handler:     ret.createRouter(),
	}

	return ret
}

func (s *HttpDbServer) createRouter() *mux.Router {

	router := mux.NewRouter()

	router.Methods(http.MethodGet).HandlerFunc(s.Get)
	router.Methods(http.MethodPut).HandlerFunc(s.Put)
	router.Methods(http.MethodDelete).HandlerFunc(s.Delete)

	return router

}

func (s *HttpDbServer) Run() error {

	return s.httpServer.ListenAndServe()

}

func (s *HttpDbServer) Shutdown() {
	ctx, _ := context.WithTimeout(context.Background(), s.config.timeout)

	s.httpServer.Shutdown(ctx)
}

func (s *HttpDbServer) Get(w http.ResponseWriter, r *http.Request) {

	item, err := s.store.Get(r.URL.Path)

	if handle500(w, err) {
		return
	}

	if handleNotFound(w, item) {
		return
	}

	handleOK(w, item)

}

func (s *HttpDbServer) Put(w http.ResponseWriter, r *http.Request) {

	content, err := ioutil.ReadAll(r.Body)

	if handle500(w, err) {
		return
	}

	created, err := s.store.Put(r.URL.Path, &store.Item{
		ContentType: r.Header.Get(headers.ContentType),
		Content:     content,
	})

	if handle500(w, err) {
		return
	}

	if created {
		handleCreated(w, r.URL.Path)
	} else {
		handleNoContentL(w, r.URL.Path)
	}

}

func (s *HttpDbServer) Delete(w http.ResponseWriter, r *http.Request) {

	err := s.store.Delete(r.URL.Path)

	if handle500(w, err) {
		return
	}

	handleNoContent(w)

}

func handleOK(w http.ResponseWriter, item *store.Item) {

	w.Header().Set(headers.ContentType, item.ContentType)
	w.Header().Set(headers.XContentTypeOptions, "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(item.Content)

}

func handleCreated(w http.ResponseWriter, location string) {

	w.Header().Set(headers.Location, location)
	w.WriteHeader(http.StatusCreated)

}

func handleNoContentL(w http.ResponseWriter, location string) {

	w.Header().Set(headers.Location, location)
	w.WriteHeader(http.StatusNoContent)

}

func handleNoContent(w http.ResponseWriter) {

	w.WriteHeader(http.StatusNoContent)

}

func handleNotFound(w http.ResponseWriter, item *store.Item) bool {

	if item == nil {
		w.WriteHeader(http.StatusNotFound)

		return true
	}

	return false

}

func handle500(w http.ResponseWriter, err error) bool {

	if err != nil {
		logrus.Error(err)

		w.Header().Set(headers.ContentType, "text/plain")
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return true
	}

	return false

}
