
# Simple API

Simple api untuk Movie dimana terdapat beberapa endpoint di antaranya:

- List Movie
- Detail Movie
- Create Movie
- Update Movie
- Delete Movie

### API Reference

#### Authorization
|   username   |  password  |
|  :---------  | :--------- |
| `simple-api` |    `xmu`   |

#### Get all Movie

```http
  GET /api/v1/movies
```

#### Detail Movie

```http
  GET /api/v1/movies/${id}
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :---------------------------------- |
| `id`      | `numeric` | **Required**. Id of movie to fetch |

#### Create Movie

```http
  POST /api/v1/movies
```

|    Body       | Type     | Description   |
|  :-------     | :------- | :------------ |
| `title`       | `string` | **Required**  |
| `description` | `string` |               |
| `rating`      | `float`  |               |
| `image`       | `string` |               |

#### Update Movie

```http
  PATCH /api/v1/movies/${id}
```
| Parameter | Type     | Description                         |
| :-------- | :------- | :---------------------------------- |
| `id`      | `numeric` | **Required**. Id of movie to fetch |

|    Body       | Type     | Description   |
|  :-------     | :------- | :------------ |
| `title`       | `string` | **Required**  |
| `description` | `string` |               |
| `rating`      | `float`  |               |
| `image`       | `string` |               |


#### Delete Movie

```http
  DELETE /api/v1/movies/${id}
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :---------------------------------- |
| `id`      | `numeric` | **Required**. Id of movie to fetch |
