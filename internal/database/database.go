package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jvbenetti/go-financial-core/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
