package ouwsdk

import (
	"context"
	"flag"

	"github.com/go-redis/redis/v8"
)


const (
	APP_NAME									= "ouw-sdk"
	APP_VERSION								= "1.0"
)


const (
	DEFAULT_CONF_FILE       = "./config.json"
	DEFAULT_HOST						= "127.0.0.1"
	DEFAULT_PORT            = "9000"
	DEFAULT_REDIS_HOST      = "127.0.0.1"
	DEFAULT_REDIS_PORT      = "6379"
)


const (
	DEFAULT_NAME_LENGTH_MIN					= 2
	DEFAULT_CHARS_MIN       				= 1
	DEFAULT_FORM_SIZE								= 20000
)


const (
	EMPTY_STR								= ""
)


const (
	KEY_CONF_GLOBAL             = "configs:global"
	KEY_CONF_INIT             	= "configs:init"
	KEY_EMAILS                  = "emails"
	KEY_FAMILIES                = "families"
	KEY_FAMILY                	= "family"
	KEY_KINGDOM                 = "kingdom"
	KEY_KINGDOMS                = "kingdoms"	
	KEY_MUNICIPAL          			= "municipal"
	KEY_MUNICIPALS          		= "municipals"
	KEY_NAMES          					=	"names"
	KEY_PLOT                   	= "plot"
	KEY_NPC                     = "npc"
	KEY_NPCS                    = "npcs"
	KEY_USER                    = "user"
	KEY_USERS                   = "users"
)


const (
	FIELD_NAME									= "name"
	FIELD_POPULATION						= "population"
	FIELD_WEALTH								= "wealth"
	FIELD_TREES									= "trees"
	FIELD_ROCKS									= "rocks"
	FIELD_LAND									= "land"
	FIELD_COWS									= "cows"
)


type ResponseErr struct {
	Msg 				string				`json:"msg"`
}


var (
	conf = flag.String("conf", DEFAULT_CONF_FILE, "config file path")
)


var Store *redis.Client

var app AppConfig

var ctx = context.Background()
