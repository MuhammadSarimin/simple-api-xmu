
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

```bash
  GET /api/v1/movies
```

#### Detail Movie

```bash
  GET /api/v1/movies/${id}
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :---------------------------------- |
| `id`      | `numeric` | **Required**. Id of movie to fetch |

#### Create Movie

```bash
  POST /api/v1/movies
```

|    Body       | Type     | Description   |
|  :-------     | :------- | :------------ |
| `title`       | `string` | **Required**  |
| `description` | `string` |               |
| `rating`      | `float`  |               |
| `image`       | `string` |               |

```bash
{
    "title": "Pengabdi Setan 2 Comunion",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan",
    "rating": 7,
    "image": "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg"
}
```

#### Update Movie

```bash
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

```bash
{
    "title": "Pengabdi Setan 3 Comunion",
    "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan",
    "rating": 7,
    "image": "https://image.tmdb.org/t/p/w500/1g0dhYtq4irTY1GPXvft6k4YLjm.jpg"
}


#### Delete Movie

```bash
  DELETE /api/v1/movies/${id}
```

| Parameter | Type     | Description                         |
| :-------- | :------- | :---------------------------------- |
| `id`      | `numeric` | **Required**. Id of movie to fetch |

### Package

 - [Gin](https://github.com/gin-gonic/gin)
 - [Gorm](https://gorm.io/)
 - [Logrus](https://github.com/sirupsen/logrus)
 - [Caarlos env](https://github.com/caarlos0/env)
 - [Caarlos env](https://github.com/caarlos0/env)
 - [Godotenv](github.com/joho/godotenv)