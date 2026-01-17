#!/bin/bash

# ----------------------------------------------------------------------------
# Copyright (c) 2026 lamaper
# 创建日期: 2026-01-17
# 最后修改: 2026-01-17
# ----------------------------------------------------------------------------

# B2CTF平台一键启动脚本（Linux Debian系列）
# 用于在Linux Debian/Ubuntu平台上一键启动B2CTF平台的前后端服务

# 配置颜色输出
GREEN="\033[32m"
YELLOW="\033[33m"
RED="\033[31m"
RESET="\033[0m"

# 错误处理函数
error_exit() {
    echo -e "${RED}[错误] $1${RESET}"
    exit 1
}

# 信息输出函数
info_msg() {
    echo -e "${GREEN}[信息] $1${RESET}"
}

# 警告输出函数
warn_msg() {
    echo -e "${YELLOW}[警告] $1${RESET}"
}

echo -e "${GREEN}=== B2CTF平台一键启动脚本（Linux Debian系列）===${RESET}"

# 检查是否为root用户
if [ "$(id -u)" == "0" ]; then
    warn_msg "不建议使用root用户运行此脚本"
fi

# 检查是否安装了Go
if ! command -v go &> /dev/null; then
    error_exit "未安装Go环境，请先安装Go 1.25+"
    echo -e "${YELLOW}[提示] 安装命令：sudo apt update && sudo apt install golang-go${RESET}"
    exit 1
fi

# 检查是否安装了Node.js
if ! command -v npm &> /dev/null; then
    error_exit "未安装Node.js环境，请先安装Node.js"
    echo -e "${YELLOW}[提示] 安装命令：sudo apt update && sudo apt install nodejs npm${RESET}"
    exit 1
fi

# 获取当前脚本所在目录
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR"

info_msg "当前工作目录：$(pwd)"

# 检查是否存在配置文件
if [ ! -f "backend/configs/config.example.yaml" ]; then
    error_exit "配置文件不存在：backend/configs/config.example.yaml"
    exit 1
fi

# 启动后端服务
info_msg "步骤1：启动后端服务..."

# 启动后端服务（后台运行）
cd backend
go run ./cmd/server -config ./configs/config.example.yaml > ../backend.log 2>&1 &
BACKEND_PID=$!
cd ..

# 等待后端启动
sleep 3

# 检查后端进程是否存在
if ps -p $BACKEND_PID > /dev/null 2>&1; then
    info_msg "后端服务已启动，进程ID：$BACKEND_PID"
    info_msg "后端日志：$(pwd)/backend.log"
else
    error_exit "后端服务启动失败，请查看日志：$(pwd)/backend.log"
    exit 1
fi

# 启动前端服务
info_msg "步骤2：启动前端服务..."

# 进入前端目录并启动服务（后台运行）
cd frontend
npm run dev > ../frontend.log 2>&1 &
FRONTEND_PID=$!
cd ..

# 等待前端启动
sleep 5

# 检查前端进程是否存在
if ps -p $FRONTEND_PID > /dev/null 2>&1; then
    info_msg "前端服务已启动，进程ID：$FRONTEND_PID"
    info_msg "前端日志：$(pwd)/frontend.log"
else
    error_exit "前端服务启动失败，请查看日志：$(pwd)/frontend.log"
    exit 1
fi

# 测试服务状态
info_msg "步骤3：测试服务状态..."

BACKEND_URL="http://localhost:8080/api/ping"
FRONTEND_URL="http://localhost:5173"

# 测试后端服务
if curl -s $BACKEND_URL | grep -q "pong"; then
    info_msg "后端服务已就绪：$BACKEND_URL"
else
    warn_msg "后端服务响应异常，请检查日志"
fi

# 显示访问信息
echo -e "${GREEN}=== 服务启动完成 ===${RESET}"
echo -e "${GREEN}前端访问地址：${RESET}$FRONTEND_URL"
echo -e "${GREEN}后端API地址：${RESET}$BACKEND_URL"
echo -e "${YELLOW}[提示] 请在浏览器中打开前端地址开始使用B2CTF平台${RESET}"
echo -e "${YELLOW}[提示] 服务正在后台运行...${RESET}"
echo -e "${YELLOW}[提示] 停止服务命令：${RESET}kill $BACKEND_PID $FRONTEND_PID"
echo -e "${YELLOW}[提示] 查看后端日志：${RESET}tail -f backend.log"
echo -e "${YELLOW}[提示] 查看前端日志：${RESET}tail -f frontend.log"

echo -e "${GREEN}=== B2CTF平台已成功启动！===${RESET}"
