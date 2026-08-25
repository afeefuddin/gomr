package gomr

import (
	"gomr/internal/worker"
	"gomr/pkg/gomr/api"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
)

func Run(mrConfig *api.MapReduceConfig) {
	ip := os.Getenv("POD_IP")
	addr := ip + ":50051"
	rawReducerCount := os.Getenv("REDUCERS_COUNT")
	reducerCount, err := strconv.Atoi(rawReducerCount)
	if err != nil || reducerCount <= 0 {
		log.Fatalf("invalid REDUCERS_COUNT %q", rawReducerCount)
	}

	worker.StartWorker(&worker.WorkerConfig{
		Mapper:       mrConfig.Mapper,
		Reducer:      mrConfig.Reducer,
		ReducerCount: reducerCount,
		WorkerAddr:   addr,
		WorkerId:     uuid.New().String(),
		MasterAddr:   "master:9001",
	})
}
