package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/k8s-go-demo/app/api"
	"github.com/k8s-go-demo/config"
	"github.com/k8s-go-demo/model"
	"log"
	"net/http"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}


// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	//Post Digital Sign
	//a.Post("/digital", a.handleRequest(api.CreateDigitalSign))
	//a.Post("/sendData", a.handleDBRequest(api.SendPIIData))
	//a.Post("/updateStatus", a.handleDBRequest(api.UpdateStatus))
	//a.Get("/getClassification/{user}", a.handleDBRequest(api.GetClassification))
	//a.Get("/getAllClassification", a.handleDBRequest(api.GetAllClassification))
	a.Get("/helloworld", a.handleDBRequest(api.GetHelloWorld))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

type RequestHandlerDBFunction func(a *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}

func (a *App) handleDBRequest(handler RequestHandlerDBFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
