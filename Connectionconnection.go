package connection

import (
	"errors"
	"log"
	"time"

	"github.com/arturoverbel/microservice_compra/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// INFO - to connect mongo
var INFO = &mgo.DialInfo{
	Addrs:    []string{"127.0.0.1:27017"},
	Timeout:  60 * time.Second,
	Database: "cool_db",
	Username: "admin",
	Password: "secret_password",
}

// DBNAME the name of the DB instance
const DBNAME = "cool_db"

// DOCNAME the name of the document
const DOCNAME = "shoppings"

var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "shoppings"
)