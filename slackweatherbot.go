package slackweatherbot

import (
	owm "github.com/briandowns/openweathermap"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

const (
	weatherTemplate = `It's currently {{.Main.Temp}} °F ({{range .Weather}} {{.Description}} {{end}}) `
)

// get the current weather conditions from openweather
func getCurrent(zip int, units, lang string, ctx context.Context) *owm.CurrentWeatherData {
	// create a urlfetch http client because we're in appengine and can't use defaults
	cl := urlfetch.Client(ctx)
	// establish connection to openweather API
	cc, err := owm.NewCurrent(units, lang, owm.WithHttpClient(cl))
	if err != nil {
		log.Errorf(ctx, "ERROR handler() during owm.NewCurrent: %s", err)
		return nil
	}
	cc.CurrentByZip(zip, "US")
	return cc
}

// redirect requests to root to /weather
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/weather", 302)
}

// handle requests to /weather
func handler_weather(w http.ResponseWriter, r *http.Request) {
	// create an appengine context so we can log
	ctx := appengine.NewContext(r)
	// check the parameters
	zip := r.URL.Query().Get("zip")
	switch zip {
	// if no zip parameter given, assume zip = 90071 (downtown LA)
	case "":
		zip = "90071"
	}
	// get the current weather data
	var zipint int
	zipint, err := strconv.Atoi(zip)
	if err != nil {
		log.Errorf(ctx, "ERROR handler_weather() zip conversion problem: %s", err)
		return
	}
	wd := getCurrent(zipint, "F", "EN", ctx)
	// make the template
	tmpl, err := template.New("weather").Parse(weatherTemplate)
	if err != nil {
		log.Errorf(ctx, "ERROR handler_weather() during template.New: %s", err)
		return
	}
	// execute the template
	err = tmpl.Execute(w, wd)
	if err != nil {
		log.Errorf(ctx, "ERROR handler_weather() during template.Execute: %s", err)
		return
	}
	return
}

// because we're in appengine, there is no main()
func init() {
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/weather", handler_weather)
}

// EOF
