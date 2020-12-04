package lifecycle

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	config "github.com/galaxy-center/galaxy/config"
	logger "github.com/galaxy-center/galaxy/log"
)

var (
	healthCheckInfo atomic.Value
	healthCheckLock sync.Mutex

	log      = logger.Get()
	checkLog = log.WithField("prefix", "lifecycle")
)

type (
	HealthCheckStatus string

	HealthCheckComponentType string
)

type HealthCheckItem struct {
	Status        HealthCheckStatus `json:"status"`
	Output        string            `json:"output,omitempty"`
	ComponentType string            `json:"componentType,omitempty"`
	ComponentID   string            `json:"componentId,omitempty"`
	Time          string            `json:"time"`
}

func setCurrentHealthCheckInfo(h map[string]HealthCheckItem) {
	healthCheckLock.Lock()
	healthCheckInfo.Store(h)
	healthCheckLock.Unlock()
}

func initHealthCheck(ctx context.Context) {
	setCurrentHealthCheckInfo(make(map[string]HealthCheckItem, 3))

	go func(ctx context.Context) {
		var n = config.Global().LivenessCheck.CheckDuration

		if n == 0 {
			n = 10
		}

		ticker := time.NewTicker(time.Second * n)

		for {
			select {
			case <-ctx.Done():

				ticker.Stop()
				checkLog.Debug("Stopping Health checks for all components")
				return

			case <-ticker.C:
				gatherHealthChecks()
			}
		}
	}(ctx)
}

func gatherHealthChecks() {
	// allInfos := SafeHealthCheck{info: make(map[string]HealthCheckItem, 3)}

	// redisStore := storage.RedisCluster{KeyPrefix: "livenesscheck-"}

	// key := "tyk-liveness-probe"

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	var checkItem = HealthCheckItem{
	// 		Status:        Pass,
	// 		ComponentType: Datastore,
	// 		Time:          time.Now().Format(time.RFC3339),
	// 	}

	// 	err := redisStore.SetRawKey(key, key, 10)
	// 	if err != nil {
	// 		mainLog.WithField("liveness-check", true).WithError(err).Error("Redis health check failed")
	// 		checkItem.Output = err.Error()
	// 		checkItem.Status = Fail
	// 	}

	// 	allInfos.mux.Lock()
	// 	allInfos.info["redis"] = checkItem
	// 	allInfos.mux.Unlock()
	// }()

	// if config.Global().UseDBAppConfigs {
	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()

	// 		var checkItem = HealthCheckItem{
	// 			Status:        Pass,
	// 			ComponentType: Datastore,
	// 			Time:          time.Now().Format(time.RFC3339),
	// 		}

	// 		if DashService == nil {
	// 			err := errors.New("Dashboard service not initialized")
	// 			mainLog.WithField("liveness-check", true).Error(err)
	// 			checkItem.Output = err.Error()
	// 			checkItem.Status = Fail
	// 		} else if err := DashService.Ping(); err != nil {
	// 			mainLog.WithField("liveness-check", true).Error(err)
	// 			checkItem.Output = err.Error()
	// 			checkItem.Status = Fail
	// 		}

	// 		checkItem.ComponentType = System

	// 		allInfos.mux.Lock()
	// 		allInfos.info["dashboard"] = checkItem
	// 		allInfos.mux.Unlock()
	// 	}()
	// }

	// if config.Global().Policies.PolicySource == "rpc" {

	// 	wg.Add(1)

	// 	go func() {
	// 		defer wg.Done()

	// 		var checkItem = HealthCheckItem{
	// 			Status:        Pass,
	// 			ComponentType: Datastore,
	// 			Time:          time.Now().Format(time.RFC3339),
	// 		}

	// 		if !rpc.Login() {
	// 			checkItem.Output = "Could not connect to RPC"
	// 			checkItem.Status = Fail
	// 		}

	// 		checkItem.ComponentType = System

	// 		allInfos.mux.Lock()
	// 		allInfos.info["rpc"] = checkItem
	// 		allInfos.mux.Unlock()
	// 	}()
	// }

	// wg.Wait()

	// allInfos.mux.Lock()
	// setCurrentHealthCheckInfo(allInfos.info)
	// allInfos.mux.Unlock()
}
