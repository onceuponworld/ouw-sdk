package ouwsdk

import (
	"context"

	"github.com/go-redis/redis/v8"
)


const (
	APP_NAME									= "ouw-sdk"
	APP_VERSION								= "1.0"
)


const (
	DEFAULT_HOST						= "127.0.0.1"
	DEFAULT_PORT            = "9000"
	DEFAULT_REDIS_HOST      = "127.0.0.1"
	DEFAULT_REDIS_PORT      = "6379"
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


type ResponseErr struct {
	Msg 				string				`json:"msg"`
}


var Rds *redis.Client

var App AppConfig

var Ctx = context.Background()
