# 2 - Web server

Web servers are a key part of modern applications, they are responsible for handling requests and responses, and they are the bridge between the user and the application.

## Objective

This project aims to create a web server that will expose the A* algorithm as a service.


## Results

- The web server is implemented with the [Gin](https://gin-gonic.com/) framework. 
- The server exposes a single endpoint `/paths` that handles GET request. It receives as query params the start and end points of the path and returns the shortest path between them.
- The result is returned as a JSON object with the path, which includes the points in an ordered list.

```http request
GET localhost:8080/paths?sx=2.0&sy=2.0&sz=0.0&ex=9.0&ey=40.0&ez=0.0
```

```json
{  "foundPath": [{"X": 2,"Y": 2,"Z": 0},{"X": 2,"Y": 3,"Z": 0},{"X": 2,"Y": 4,"Z": 0},{"X": 2,"Y": 5,"Z": 0},{"X": 2,"Y": 6,"Z": 0},{"X": 2,"Y": 7,"Z": 0},{"X": 2,"Y": 8,"Z": 0},{"X": 2,"Y": 9,"Z": 0},{"X": 2,"Y": 10,"Z": 0},{"X": 2,"Y": 11,"Z": 0},{"X": 2,"Y": 12,"Z": 0},{"X": 2,"Y": 13,"Z": 0},{"X": 2,"Y": 14,"Z": 0},{"X": 2,"Y": 15,"Z": 0},{"X": 2,"Y": 16,"Z": 0},{"X": 2,"Y": 17,"Z": 0},{"X": 2,"Y": 18,"Z": 0},{"X": 2,"Y": 19,"Z": 0},{"X": 2,"Y": 20,"Z": 0},{"X": 2,"Y": 21,"Z": 0},{"X": 2,"Y": 22,"Z": 0},{"X": 2,"Y": 23,"Z": 0},{"X": 2,"Y": 24,"Z": 0},{"X": 2,"Y": 25,"Z": 0},{"X": 2,"Y": 26,"Z": 0},{"X": 2,"Y": 27,"Z": 0},{"X": 2,"Y": 28,"Z": 0},{"X": 2,"Y": 29,"Z": 0},{"X": 2,"Y": 30,"Z": 0},{"X": 2,"Y": 31,"Z": 0},{"X": 2,"Y": 32,"Z": 0},{"X": 2,"Y": 33,"Z": 0},{"X": 3,"Y": 33,"Z": 0},{"X": 3,"Y": 34,"Z": 0},{"X": 4,"Y": 34,"Z": 0},{"X": 4,"Y": 35,"Z": 0},{"X": 5,"Y": 35,"Z": 0},{"X": 5,"Y": 36,"Z": 0},{"X": 6,"Y": 36,"Z": 0},{"X": 6,"Y": 37,"Z": 0},{"X": 7,"Y": 37,"Z": 0},{"X": 7,"Y": 38,"Z": 0},{"X": 7,"Y": 39,"Z": 0},{"X": 8,"Y": 39,"Z": 0},{"X": 9,"Y": 39,"Z": 0},{"X": 9,"Y": 40,"Z": 0}  ]}
```
## Next steps

Next project will focus on improving the architecture of the web service to allow more complex scenarios and better handling of the data. 