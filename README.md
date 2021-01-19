# social-rss-api
social rss api service.



## API DOC

- social rss list

```go
curl -X GET http://localhost:8080/hansu

{
    "code": 200,
    "look_in": [
        {
            "nick_name": "hansu",
            "time": "2021-01-19 18:54:45",
            "URL": "https://github.com/suhan1996"
        },
        {
            "nick_name": "hansu",
            "time": "2021-01-19 18:54:45",
            "URL": "https://google.com/"
        }
    ],
    "msg": "successful.",
    "user_info": {
        "nick_name": "hansu",
        "account": "hansu@privoce.com",
        "avatar_url": "https://iconfont.alicdn.com/t/1557889264039.JPG@100h_100w.jpg"
    }
}

```
