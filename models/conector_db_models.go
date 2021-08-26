package models

import "github.com/hendytriatmoko/first/config"

//"../config"
// "github.com/gin-gonic/gin"

var (
	db  = config.DBInit()
	idb = config.InDB{DB: db}
	// idb = inDB

)
