package gql

import (
	"context"
	"log"
	"main/internal/database"
	"main/internal/repository"
	graph "main/internal/transport/gql/graph"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/vektah/gqlparser/v2/ast"
)

type server struct {
	server *http.Server
}

func CreateServer(storage *database.Storage) *server {

	storageInterface := repository.InitRepositoryInterface(storage)
	resolvers := graph.InitResolvers(storageInterface)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolvers}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux := chi.NewRouter()

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return &server{
		server: &http.Server{
			Addr:    ":4010",
			Handler: mux,
		},
	}
}

func (a *server) Run() {
	go func() {
		log.Printf("Gql. Сервер запущен на порту %s", a.server.Addr)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Ошибка запуска: %s", err)
		}
	}()
}

func (a *server) Stop(ctx context.Context) {
	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("Принудительное завершение: %v", err)
	}
}
