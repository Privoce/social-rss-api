# social-rss-api
social rss api service.

## Models

### UserModel

| Key             | Type   | Required | Default |
| --------------- | ------ | -------- | ------- |
| nick_name       | String | true     | N/A     |
| account         | String | true     | N/A     |
| avatar_url      | String | true    |  N/A     |


### LookInModel

| Key             | Type   | Required | Default |
| --------------- | ------ | -------- | ------- |
| nick_name       | String | true     | N/A     |
| time            | String | true     | N/A     |
| url             | String | true    |  N/A     |

## API

### RSS LIST

#### `domain/signin`

URL: `/NickName`

Method: `GET`

Param:

```json
curl —X GET http://localhost:8080/ding
```

Response:

Status Code: `200 | 400 | 401`

Cause: `success | user not exist   | user look in no data`

```json
{
    "code": 200,
    "look_in": [
        {
            "nick_name": "ding",
            "time": "2021-01-19 18:54:45",
            "url": "https://github.com/higker"
        },
        {
            "nick_name": "ding",
            "time": "2021-01-19 18:54:45",
            "url": "https://google.com/"
        },
        {
            "nick_name": "ding",
            "time": "2021-01-19 18:54:45",
            "url": "https://google.com/"
        }
    ],
    "msg": "successful.",
    "user_info": {
        "nick_name": "ding",
        "account": "ding@privoce.com",
        "avatar_url": "https://iconfont.alicdn.com/t/1570850261949.png@100h_100w.jpg"
    }
}



