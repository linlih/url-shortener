# url-shortener

实现一个端URL服务。

测试：
- 生成短URL
    ```bash
      curl --request POST --data '{ "long_url": "https://www.baidu.com", "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10" }' http://localhost:9808/create-short-url
    ```
- 然后访问生成的短URL看是否正常跳转到长URL的页面中

参考来源：
[Let's build a URL shortener in Go](https://www.eddywm.com/lets-build-a-url-shortener-in-go)

