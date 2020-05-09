#### config.json 存放配置文件

```js
{   //数据库地址
    "DBURL":"duttt:dduutt@tcp(165.22.249.2:3306)/iot?charset=utf8&parseTime=true",
    //数据库类型
    "DBType":"mysql",
    //web服务端口号
    "APIAddr":":15001",
    //Socket服务端口号
    "TCPAddr":":15000",
    //JWT密钥
    "JwtSecret":"DDUUTT",
    //mqtt broker 地址
    "MqttAddr":"ws://mq.tongxinmao.com:18832/web",
    //mqtt 客户端ID
    "MqttClientID":"duttt_11223344"
}
```
mqtt服务器默认地址为 *ws://mq.tongxinmao.com:18832/web* ，可以在<http://www.tongxinmao.com/txm/webmqtt.php> 测试
*topic：/public/TEST/dutong/#*
#### mserver.exe 启动后端服务
访问：<http://127.0.0.1:15001/swagger/index.html> 查看API文档

#### 测试方法：
1. 设置Socket 为client 连接本机 15000端口，首次连接发送 8801开头 + 14位长度数字（如：8801123456789123） 共16位长度
的devID作为区分不同设备的设备号。
2. 访问 <http://127.0.0.1:15001/swagger/index.html> 查看API，注册用户，然后登录获取 *token* ,在 *Authorize* 处认证后可开始测试。
3. 添加数据模板
### request：
```json
{
  "dataPoints": [
    {
      "dataType": 0,
      "formula": "",
      "message": "0300000001",
      "name": "温度",
      "unit": "℃"
    },
     {
      "dataType": 0,
      "formula": "v/100",
      "message": "0300000001",
      "name": "湿度",
      "unit": "%RH"
    }
  ],
  "templateName": "温湿度二合一模板",
  "username": "123456"
}
```
### response
```json
{
  "code": 200,
  "data": {
    "id": 2,
    "templateName": "温湿度二合一模板",
    "username": "123456",
    "dataPoints": [
      {
        "id": 3,
        "name": "温度",
        "message": "0300000001",
        "dataType": 0,
        "unit": "℃",
        "formula": "",
        "templateID": 2,
        "createdAt": "2019-11-29T15:00:38.772819+08:00"
      },
      {
        "id": 4,
        "name": "湿度",
        "message": "0300000001",
        "dataType": 0,
        "unit": "%RH",
        "formula": "v/100",
        "templateID": 2,
        "createdAt": "2019-11-29T15:00:38.9202566+08:00"
      }
    ],
    "createdAt": "2019-11-29T15:00:38.6261726+08:00",
    "updatedAt": "2019-11-29T15:00:38.6261726+08:00"
  },
  "msg": "ok"
}
```
4. 添加设备

### request
```json
{
  "addr": "江西南昌",
  "cKey": "1234",
  "devID": "8801123456789124",
  "deviceName": "W610",
  "deviceType": 0,
  "frequency": 5,
  "position": "123.014,154.214",
  "slavers": [
    {
      "slaverIndex": 1,
      "slaverName": "温湿度二合一传感器",
      "templateID": 2,
      "templateName": "温湿度二合一模板"
    },
    {
      "slaverIndex": 2,
      "slaverName": "温湿度二合一传感器",
      "templateID": 2,
      "templateName": "温湿度二合一模板"
    }
  ],
  "username": "123456"
}
```
### response
```json
{
  "code": 200,
  "data": {
    "id": 2,
    "devID": "8801123456789124",
    "cKey": "1234",
    "frequency": 5,
    "deviceName": "W610",
    "deviceType": 0,
    "addr": "江西南昌",
    "position": "123.014,154.214",
    "username": "123456",
    "slavers": [
      {
        "id": 2,
        "slaverName": "温湿度二合一传感器",
        "slaverIndex": 1,
        "templateID": 2,
        "templateName": "温湿度二合一模板",
        "devID": 2,
        "createdAt": "2019-11-29T15:07:06.5161107+08:00",
        "updatedAt": "2019-11-29T15:07:06.5161107+08:00"
      },
      {
        "id": 3,
        "slaverName": "温湿度二合一传感器",
        "slaverIndex": 2,
        "templateID": 2,
        "templateName": "温湿度二合一模板",
        "devID": 2,
        "createdAt": "2019-11-29T15:07:06.6709849+08:00",
        "updatedAt": "2019-11-29T15:07:06.6709849+08:00"
      }
    ],
    "createdAt": "2019-11-29T15:07:06.352517+08:00",
    "updatedAt": "2019-11-29T15:07:06.352517+08:00"
  },
  "msg": "ok"
}
```
5. 添加定时任务

## 注意：
这里是由前端在添加设备时发起的请求，填写的query是由：

*从机地址* + *数据点的message* + *CRC16* 组成

**下面的指令仅为测试使用**
### request
```json
{
  "devID": "8801123456789124",
  "frequency": 4,
  "tasks": [
    {
      "pointID": 3,
      "query": "010300000002c40b"
    },
    {
      "pointID": 4,
      "query": "010300000002c40b"
    }
  ]
}
```

### response
```json
{
  "code": 200,
  "data": {
    "id": 2,
    "devID": "8801123456789124",
    "frequency": 4,
    "tasks": [
      {
        "id": 3,
        "pointID": 3,
        "deviceTaskID": 2,
        "query": "010300000002c40b"
      },
      {
        "id": 4,
        "pointID": 4,
        "deviceTaskID": 2,
        "query": "010300000002c40b"
      }
    ]
  },
  "msg": "ok"
}
```
6. 查看传感器数据

- 访问: <http://www.tongxinmao.com/txm/webmqtt.php>

    **订阅topic**  */public/TEST/dutong/#*
- API测试数据页面

    输入 *pointID* 要查看的开始结束时间(unix时间戳)
    ### response
    ```json
    {
        "code": 200,
        "data": [
            {
            "id": 16,
            "data": "010300000002c40b",
            "cTime": 1574855946,
            "pointID": 1
            },
            {
            "id": 17,
            "data": "010300000002c40b",
            "cTime": 1574855949,
            "pointID": 1
            }
        ],
            "msg":"ok"
    }
    ```

