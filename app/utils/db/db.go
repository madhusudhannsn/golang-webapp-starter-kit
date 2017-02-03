package db

import (
	"github.com/spf13/viper"
	"github.com/madhusudhannsn/go-web-app/app/utils/error"
	"github.com/madhusudhannsn/go-web-app/app/utils/logger"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetSession : This is the initial session we will be obtaining. All the Other sessions will be copied from this
func GetSession() (*error.AppError, *mgo.Session) {
	if session == nil {
		return createSession()
	} else {
		logger.Info.Println("Session", session)
		return nil, session
	}
}

func createSession() (*error.AppError, *mgo.Session) {
	logger.Info.Println("Initializing the mongo session")
	session, err := mgo.Dial(viper.GetString("db.url"))
	var appErr *error.AppError
	if err != nil {
		appErr = error.GetInstance("DB Initialization error", 1001, err)
	}
	return appErr, session

}
