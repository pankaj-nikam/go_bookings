package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pankaj-nikam/go_bookings/internal/config"
	"github.com/pankaj-nikam/go_bookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	//What am I going to store in session
	gob.Register(models.Reservation{})

	//Change this to true when in production
	testApp.IsProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	session.Cookie.HttpOnly = true

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}
