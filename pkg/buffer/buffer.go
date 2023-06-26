package buffer

import (
	"assisment/internal/models"
	"assisment/pkg/database"
)

type LogBuffer struct {
	logs chan models.Log
}

func NewLogBuffer() *LogBuffer {
	return &LogBuffer{
		logs: make(chan models.Log, 1000), // Define the buffer to include number of request
	}
}

func (lb *LogBuffer) StartProcessing(db *database.DatabaseImpl) {
	for log := range lb.logs {
		err := db.StoreLogInDatabase(&log)
		if err != nil {
			panic("DB Insertion Failed")
		}
	}
}

func (lb *LogBuffer) AddToBuffer(log *models.Log) {
	lb.logs <- *log
}
