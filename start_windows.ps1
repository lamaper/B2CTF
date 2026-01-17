<#
.SYNOPSIS
一键启动B2CTF平台（Windows）

.DESCRIPTION
该脚本用于在Windows平台上一键启动B2CTF平台的前后端服务

.EXAMPLE
.tart_windows.ps1

.NOTES
Copyright (c) 2026 lamaper
创建日期: 2026-01-17
最后修改: 2026-01-17
Author: lamaper
#>

# 配置颜色输出
$Green = "[32m"
$Yellow = "[33m"
$Red = "[31m"
$Reset = "[0m"

Write-Host "${Green}=== B2CTF平台一键启动脚本（Windows）===${Reset}"

# 检查是否安装了Go
if (-not (Get-Command "go" -ErrorAction SilentlyContinue)) {
    Write-Host "${Red}[错误] 未安装Go环境，请先安装Go 1.25+${Reset}"
    Write-Host "${Yellow}[提示] 下载地址：https://go.dev/dl/${Reset}"
    exit 1
}

# 检查是否安装了Node.js
if (-not (Get-Command "npm" -ErrorAction SilentlyContinue)) {
    Write-Host "${Red}[错误] 未安装Node.js环境，请先安装Node.js${Reset}"
    Write-Host "${Yellow}[提示] 下载地址：https://nodejs.org/zh-cn/download/${Reset}"
    exit 1
}

# 获取当前脚本所在目录
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
cd $ScriptDir

Write-Host "${Green}[信息] 当前工作目录：$PWD${Reset}"

# 启动后端服务
Write-Host "${Green}[步骤1] 启动后端服务...${Reset}"

# 创建后端启动脚本
$BackendScript = Join-Path $ScriptDir "backend_start.ps1"
@'
cd backend
go run ./cmd/server -config .\configs\config.example.yaml
'@ | Set-Content $BackendScript

# 启动后端服务（新窗口）
Start-Process powershell.exe "-NoProfile -ExecutionPolicy Bypass -File '$BackendScript'" -WindowStyle Normal -WorkingDirectory $ScriptDir

# 等待后端启动
Start-Sleep -Seconds 3

# 启动前端服务
Write-Host "${Green}[步骤2] 启动前端服务...${Reset}"

# 创建前端启动脚本
$FrontendScript = Join-Path $ScriptDir "frontend_start.ps1"
@'
cd frontend
npm run dev
'@ | Set-Content $FrontendScript

# 启动前端服务（新窗口）
Start-Process powershell.exe "-NoProfile -ExecutionPolicy Bypass -File '$FrontendScript'" -WindowStyle Normal -WorkingDirectory $ScriptDir

# 等待前端启动
Start-Sleep -Seconds 5

# 测试服务状态
Write-Host "${Green}[步骤3] 测试服务状态...${Reset}"

$BackendUrl = "http://localhost:8080/api/ping"
$FrontendUrl = "http://localhost:5173"

# 测试后端服务
try {
    $BackendResponse = Invoke-RestMethod -Uri $BackendUrl -TimeoutSec 5
    if ($BackendResponse.msg -eq "pong") {
        Write-Host "${Green}[成功] 后端服务已启动：$BackendUrl${Reset}"
    } else {
        Write-Host "${Yellow}[警告] 后端服务响应异常${Reset}"
    }
} catch {
    Write-Host "${Red}[错误] 后端服务连接失败：$BackendUrl${Reset}"
    Write-Host "${Yellow}[提示] 请检查后端窗口中的错误信息${Reset}"
}

# 清理临时脚本
Remove-Item $BackendScript, $FrontendScript -Force -ErrorAction SilentlyContinue

# 显示访问信息
Write-Host "${Green}=== 服务启动完成 ===${Reset}"
Write-Host "${Green}前端访问地址：${Reset}$FrontendUrl"
Write-Host "${Green}后端API地址：${Reset}$BackendUrl"
Write-Host "${Yellow}[提示] 请在浏览器中打开前端地址开始使用B2CTF平台${Reset}"
Write-Host "${Yellow}[提示] 按Ctrl+C关闭此窗口，服务仍会在后台运行${Reset}"
Write-Host "${Yellow}[提示] 若要停止服务，请关闭对应的PowerShell窗口${Reset}"

# 等待用户输入
Read-Host "${Green}按Enter键退出...${Reset}"
