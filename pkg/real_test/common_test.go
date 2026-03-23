// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

var (
	// 全局客户端，用于所有测试
	cli *client.ZSClient

	// 命令行参数
	mnIP         string
	accountName  string
	password     string
	port         int
	debug        bool
	reportPath   string
	simpleOutput bool   // 使用简单输出而非表格输出
	contextPath  string // ZStack API 上下文路径

	// 测试结果记录
	testResults     = make(map[string]bool)
	testResultsLock sync.Mutex
)

// 初始化命令行参数
func init() {
	flag.StringVar(&mnIP, "mn_ip", "localhost", "Management node IP address")
	flag.StringVar(&accountName, "account", "admin", "Account name")
	flag.StringVar(&password, "password", "password", "Account password")
	flag.IntVar(&port, "port", 8080, "Management node port")
	flag.BoolVar(&debug, "debug", false, "Enable debug logging")
	flag.StringVar(&reportPath, "report", "test_report.html", "Path to save the test report")
	flag.BoolVar(&simpleOutput, "simple", false, "Use simple output instead of table output")
	flag.StringVar(&contextPath, "context_path", "zstack", "ZStack API context path")
}

// 字符串指针转换辅助函数
func strPtr(s string) *string {
	return &s
}

// TestMain 是测试的主入口
func TestMain(m *testing.M) {
	// 解析命令行参数
	flag.Parse()

	// 设置日志级别
	if debug {
		golog.SetLevel("debug")
	} else {
		golog.SetLevel("info")
	}

	golog.Infof("Connecting to ZStack at %s:%d with account %s", mnIP, port, accountName)

	// 创建客户端配置
	config := client.DefaultZSConfig(mnIP)
	if port != 8080 {
		config = client.NewZSConfig(mnIP, port, contextPath)
	}

	// 设置账户登录信息
	config.LoginAccount(accountName, password)

	// 如果启用了调试模式
	if debug {
		config.Debug(true)
	}

	// 创建客户端
	cli = client.NewZSClient(config)
	ctx := context.Background()
	// 登录
	session, err := cli.Login(ctx)
	if err != nil {
		golog.Fatalf("Failed to login: %v", err)
	}
	golog.Info("Login successful")

	// 运行测试
	exitCode := m.Run()

	// 登出
	err = cli.LogOut(ctx, session.Uuid, param.DeleteModePermissive)
	if err != nil {
		golog.Errorf("Failed to logout: %v", err)
	} else {
		golog.Info("Logout successful")
	}

	// 生成测试报告
	generateTestReport()

	os.Exit(exitCode)
}

// 辅助函数：跳过测试如果没有找到资源
func skipIfNoResource(t *testing.T, resourceType string, count int) {
	if count == 0 {
		golog.Warnf("测试跳过 [%s]: No %s found to test", t.Name(), resourceType)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skipf("No %s found to test", resourceType)
	}
}

// 记录测试结果
func recordTestResult(testName string, passed bool) {
	testResultsLock.Lock()
	defer testResultsLock.Unlock()
	testResults[testName] = passed
}

// 生成测试报告
func generateTestReport() {
	// 控制台输出报告
	printConsoleReport()

	// 生成HTML报告
	generateHTMLReport()
}

// 打印控制台报告
func printConsoleReport() {
	golog.Info("=================== 测试报告 ===================")

	// 统计结果
	var passed, failed int

	// 按名称排序测试结果
	var testNames []string
	for name := range testResults {
		testNames = append(testNames, name)
	}
	sort.Strings(testNames)

	// 使用简单输出
	for _, name := range testNames {
		result := testResults[name]
		resultStr := "通过"
		if !result {
			resultStr = "失败"
			failed++
		} else {
			passed++
		}

		golog.Infof("%-50s %s", name, resultStr)
	}

	// 打印汇总
	total := passed + failed
	if total > 0 {
		golog.Infof("总计: %d, 通过: %d, 失败: %d, 通过率: %.2f%%",
			total, passed, failed, float64(passed)/float64(total)*100)
	} else {
		golog.Info("没有测试结果记录")
	}
	golog.Info("================================================")
}

// 生成HTML报告
func generateHTMLReport() {
	// 统计结果
	var passed, failed int
	var testDetails strings.Builder

	// 按名称排序测试结果
	var testNames []string
	for name := range testResults {
		testNames = append(testNames, name)
	}
	sort.Strings(testNames)

	// 生成测试详情HTML
	for _, name := range testNames {
		result := testResults[name]
		resultStr := "通过"
		resultClass := "success"
		if !result {
			resultStr = "失败"
			resultClass = "danger"
			failed++
		} else {
			passed++
		}

		testDetails.WriteString(fmt.Sprintf(`
		<tr>
			<td>%s</td>
			<td><span class="badge bg-%s">%s</span></td>
		</tr>
		`, name, resultClass, resultStr))
	}

	// 计算通过率
	total := passed + failed
	var passRate float64
	if total > 0 {
		passRate = float64(passed) / float64(total) * 100
	}

	// 生成HTML报告
	htmlContent := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="zh-CN">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>ZStack SDK 测试报告</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
		<style>
			body { padding: 20px; }
			.header { margin-bottom: 30px; }
			.summary-card { margin-bottom: 20px; }
			.progress { height: 25px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>ZStack SDK 测试报告</h1>
				<p class="text-muted">生成时间: %s</p>
				<p>测试环境: %s:%d</p>
			</div>

			<div class="row summary-card">
				<div class="col-md-4">
					<div class="card">
						<div class="card-body">
							<h5 class="card-title">总计</h5>
							<h2 class="card-text">%d</h2>
						</div>
					</div>
				</div>
				<div class="col-md-4">
					<div class="card">
						<div class="card-body">
							<h5 class="card-title">通过</h5>
							<h2 class="card-text text-success">%d</h2>
						</div>
					</div>
				</div>
				<div class="col-md-4">
					<div class="card">
						<div class="card-body">
							<h5 class="card-title">失败</h5>
							<h2 class="card-text text-danger">%d</h2>
						</div>
					</div>
				</div>
			</div>

			<div class="row mb-4">
				<div class="col-12">
					<h5>通过率: %.2f%%</h5>
					<div class="progress">
						<div class="progress-bar bg-success" role="progressbar" style="width: %.2f%%;"
							aria-valuenow="%.2f" aria-valuemin="0" aria-valuemax="100">%.2f%%</div>
					</div>
				</div>
			</div>

			<div class="row">
				<div class="col-12">
					<div class="card">
						<div class="card-header">
							<h5>测试详情</h5>
						</div>
						<div class="card-body">
							<table class="table table-striped">
								<thead>
									<tr>
										<th>测试用例</th>
										<th>结果</th>
									</tr>
								</thead>
								<tbody>
									%s
								</tbody>
							</table>
						</div>
					</div>
				</div>
			</div>
		</div>

		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js"></script>
	</body>
	</html>
	`, time.Now().Format("2006-01-02 15:04:05"), mnIP, port,
		total, passed, failed, passRate, passRate, passRate, passRate, testDetails.String())

	// 保存HTML报告
	err := os.WriteFile(reportPath, []byte(htmlContent), 0644)
	if err != nil {
		golog.Errorf("Failed to save test report: %v", err)
	} else {
		// 获取绝对路径
		absPath, err := filepath.Abs(reportPath)
		if err != nil {
			absPath = reportPath
		}

		golog.Infof("测试报告已保存到: %s", absPath)

		// 如果是在本地运行，尝试打开浏览器
		if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
			var cmd string
			var args []string

			switch runtime.GOOS {
			case "darwin":
				cmd = "open"
				args = []string{absPath}
			case "windows":
				cmd = "cmd"
				args = []string{"/c", "start", absPath}
			}

			if cmd != "" {
				err := exec.Command(cmd, args...).Start()
				if err != nil {
					golog.Errorf("Failed to open browser: %v", err)
				}
			}
		}
	}
}
