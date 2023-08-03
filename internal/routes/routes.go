package routes

import (
	serviceConfig "github.com/Asad2730/DynamoDB_CRUD_App/config"
	"github.com/Asad2730/DynamoDB_CRUD_App/internal/repositories/adapter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeOut(serviceConfig.GetConfig().TimeOut),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repositories adapter.Interface) *chi.Mux {
	r.setConfigRouters()
	r.RouterHealth(repositories)
	r.RouterProduct(repositories)
	return r.router
}

func (r *Router) setConfigRouters() {
	r.EnableCors()
	r.EnableLogger()
	r.EnableTimeOut()
	r.EnableRecover()
	r.EnableRequestID()
	r.EnableRealIP()
}

func (r *Router) RouterHealth(repositories adapter.Interface) {
	handler := HealthHandler.NewHandler(repositories)

	r.router.Route("/health", func(r chi.Router) {
		r.Get("/", handler.Get)
		r.Post("/", handler.Post)
		r.Put("/", handler.Put)
		r.Delete("/", handler.Delete)
		r.Options("/", handler.Options)
	})
}

func (r *Router) RouterProduct(repositories adapter.Interface) {

	handler := ProductHandler.NewHandler(repositories)

	r.router.Route("/product", func(r chi.Router) {
		r.Get("/", handler.Get)
		r.Post("/", handler.Post)
		r.Put("/{ID}", handler.Put)
		r.Delete("/{ID}", handler.Delete)
		r.Options("/", handler.Options)
	})

}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}

func (r *Router) EnableTimeOut() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeOut()))
	return r
}

func (r *Router) EnableCors() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecover() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestID() *Router {
	r.router.Use(middleware.RequestID)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}
