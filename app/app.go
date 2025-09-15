package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
	"ucl-epreuve-technique/app/adapters"
	"ucl-epreuve-technique/app/controllers"
	"ucl-epreuve-technique/app/middlewares"
	"ucl-epreuve-technique/app/models"
	"ucl-epreuve-technique/app/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// Specific handler For single page application
type spaHandler struct {
	staticPath string
	indexPath  string
}

func (app *App) InitializeRoutes() {

	routesAdapters := []adapters.Adapter{adapters.Recover(false), adapters.API(true)}

	app.Router = mux.NewRouter()

	app.Router.Handle("/api/v1/liste-cours", adapters.Adapt(controllers.GetCoursHandler, routesAdapters...)).Methods("GET")
	app.Router.Handle("/api/v1/liste-inscriptions", adapters.Adapt(controllers.GetInscriptionsHandler, routesAdapters...)).Methods("GET")
	app.Router.Handle("/api/v1/liste-notes", adapters.Adapt(controllers.GetNotesHandler, routesAdapters...)).Methods("GET")
	app.Router.Handle("/api/v1/etudiants/{etudiant_nom}/annee/{annee}/bulletin", adapters.Adapt(controllers.GetBulletinHandler, routesAdapters...)).Methods("GET")

	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}
	app.Router.PathPrefix("/").Handler(spa)
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		// file does not exist or path is a directory, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func (app *App) Run() {

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	corsHeaders := handlers.AllowedHeaders([]string{"content-type", "cache-control", "x-requested-with", "Authorization"})
	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsCredentials := handlers.AllowCredentials()

	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	logger := log.New(os.Stdout, "", log.LstdFlags)

	routesMiddlewares := middlewares.RoutesMiddlewares( /*middlewares.KeycloakAuthMiddleware("USERS"),*/ middlewares.Logging(logger))

	port := utils.GetEnv("APPLICATION_PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr: "0.0.0.0:" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(corsHeaders, corsOrigins, corsCredentials, methods)(routesMiddlewares(app.Router)),
	}

	// Launch the srv in his own goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c
	log.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	defer models.GetDB().Close()
	os.Exit(0)

}
