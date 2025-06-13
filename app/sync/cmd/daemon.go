package cmd

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"metaLand/app/sync/service"
	"metaLand/app/sync/service/config"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "sync contract info.",
	Long:  "sync contract info.",
	Run: func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		// rpc退出信号通知chan
		onSyncExit := make(chan error, 1)

		go func() {
			defer wg.Done()

			cfg, err := config.UnmarshalCmdConfig() // 读取和解析配置文件
			if err != nil {
				fmt.Println("Failed to unmarshal config", zap.Error(err))
				onSyncExit <- err
				return
			}

			logx.MustSetup(cfg.Log) // 初始化日志模块
			logx.Info("sync server start", zap.Any("config", cfg))

			s, err := service.New(ctx, cfg) // 初始化服务
			if err != nil {
				logx.Error("Failed to create sync server", zap.Error(err))
				onSyncExit <- err
				return
			}

			s.Start()

			if cfg.Monitor.PprofEnable { // 开启pprof，用于性能监控
				http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.Monitor.PprofPort), nil)
			}
		}()

		// 信号通知chan
		onSignal := make(chan os.Signal)
		// 优雅退出
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-onSignal:
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				cancel()
				logx.Info("Exit by signal", zap.String("signal", sig.String()))
			}
		case err := <-onSyncExit:
			cancel()
			logx.Error("Exit by error", zap.Error(err))
		}
		wg.Wait()
	},
}

func init() {
	// 将api初始化命令添加到主命令中
	rootCmd.AddCommand(DaemonCmd)
}
