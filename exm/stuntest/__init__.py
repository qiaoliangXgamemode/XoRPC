import requests
import time
from nonebot import on_command
from nonebot.adapters.onebot.v11.bot import Bot
from nonebot.permission import SUPERUSER

# 需要指令启动 bot，需要对应用户拥有 SUPERUSER 权限
# SUPERUSER 需要在 bot 根目录的 .env 文件中指定
# 添加如下行即可（123123123 替换为 QQ 号）
# SUPERUSERS=["123123123"]
startListening = on_command("开始获取Jenkins信息", permission=SUPERUSER)

@startListening.handle()
async def jenkinsListening():
    # 初始化构建号缓存
    buildNumberTemp=0
    
    # 60s 查询一次 Rest API
    while True:
        # 获取鹦鹉通道 Rest API 信息
        headers={
            'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188'
        }
        url='http://jk-insider.bakaxl.com:8888/job/BakaXL%20Insider%20Parrot/api/python?pretty=true'
        response = requests.get(url=url,params='',headers=headers)
        apiStr = response.text

        # 获取最新构建号 可能写的有点屎山
        apiStrSplitedFirst = apiStr.split('"number" : ', 1)
        apiStrSplitedSecond = apiStrSplitedFirst[1].split(',', 1)
        buildNumberGetted = apiStrSplitedSecond[0]

        # 与缓存构建号比对 若一致则不发送消息
        if(buildNumberGetted == buildNumberTemp):
            continue
        else:
            # 获取鹦鹉通道最新构建信息
            headers={
            'User-Agent':'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188'
            }
            url='http://jk-insider.bakaxl.com:8888/job/BakaXL%20Insider%20Parrot/' + buildNumberGetted + '/api/python?pretty=true'
            response = requests.get(url=url,params='',headers=headers)
            buildStr = response.text
            buildStrSplitedFirst = buildStr.split('"comment" : "', 1)
            buildStrSplitedSecond = buildStrSplitedFirst[1].split('"', 1)
            buildMsgRaw = buildStrSplitedSecond[0]

            # 替换字符串中 \u000a 转义为 \n 转义符
            buildMsg = buildMsgRaw.replace('\\u000a', '\n')

            # 发送消息
            await startListening.send("[下面播报 BakaXL 预览体验构建的相关信息]\nBakaXL 预览体验构建 鹦鹉通道 #" + buildNumberGetted + "\n构建描述：\n" + buildMsg)

            # 替换缓存构建号
            buildNumberTemp = buildNumberGetted
        
        # 等待 60s 再次查询信息
        time.sleep(60)