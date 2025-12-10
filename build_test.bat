@echo off
echo ==========================================
echo  Building Docker Image: new-api:test
echo ==========================================

REM 检查是否安装了 Docker
docker --version >nul 2>&1
IF %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Docker 未安装或未加入 PATH！
    pause
    exit /b 1
)

REM 构建镜像
docker build -t new-api:test .
IF %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Docker 构建失败！
    pause
    exit /b 1
)

echo.
echo ==========================================
echo  Saving Docker Image to new-api-image-test.tar
echo ==========================================

docker save -o new-api-image-test.tar new-api:test
IF %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Docker 镜像保存失败！
    pause
    exit /b 1
)

echo.
echo [SUCCESS] 镜像构建并保存完成！
echo 文件: new-api-image-test.tar
pause
