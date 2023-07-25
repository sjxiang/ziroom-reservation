echo -e "\033[33m开始 检查本地是否搭建 Go 开发环境 ... \033[0m"
	go version
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo -e "\033[32m[成功] 检测到 Go 开发环境!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the GO development environment, please refer to the manual for installation.\033[0m"
    	exit
    fi

echo -e "\033[33m开始 检查本地是否安装 docker 应用 ... \033[0m"
	docker -v
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo -e "\033[32m[成功] 检测到 docker 应用!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the docker application, please refer to the manual for installation.\033[0m"
      exit
    fi

echo -e "\033[33m开始 检查本地是否安装 docker-compose 应用 ... \033[0m"
	docker-compose version
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo -e "\033[32m[成功] 检测到 docker-compose 应用!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the docker-compose application, please refer to the manual for installation.\033[0m"
      exit
    fi

echo -e "\033[33m开始 尝试构建 ... \033[0m"
  go build -o young_engine
    if [ $? -eq  0 ]; then
        echo -e "\033[32m[成功] ✅ 构建成功!\033[0m"
        rm ./young_engine
    else
    	echo "\033[31m[错误] ❌ build failed, Please resolve the issue based on the error message\033[0m"
      exit
    fi

echo ""
echo -e "\033[32m 恭喜，已经安装所有的项目依赖！开启项目之旅！\033[0m"