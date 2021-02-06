# Felstorm

![image](https://goreportcard.com/badge/github.com/smh2274/Felstorm)
![felstorm integration test](https://github.com/smh2274/Felstorm/workflows/felstorm%20integration%20test/badge.svg?branch=main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e490a3bb49ba44bca77fb9ee2340fcc3)](https://www.codacy.com/gh/smh2274/Felstorm/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=smh2274/Felstorm&amp;utm_campaign=Badge_Grade)

![image](https://img.zcool.cn/community/0175a85549a6d20000019ae941d3a4.jpg@1280w_1l_2o_100sh.jpg)

## this project used grpc frame

### generate json web token with HS256

  * step 1
> clone this repository && cd this project
 
  * step 2
> docker network -d bridge $(your network)

  * step 3
> docker build . -t felstorm

  * step 3
> docker  run --name felstorm -v $(your log path):/Azeroth/Felstorm/log --network $(your network) --network-alias felstorm -p 8800:8800 -it -d felstorm
        
#### configure self

  * you could change felstorm_conf.yaml file
  * if you want change access network gateway, you can do some change on envoy.yaml
