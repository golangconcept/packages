# gorilla/mux

`gorilla/mux` is a popular and flexible HTTP router and URL matcher for Go.

It is often used for building REST `APIs`, `web services`, or any kind of web application where `routing` and handling HTTP requests is needed.

## Features of mux

- `Route matching`: Allows for complex routing patterns, including variables in the URL.
- `URL variables`: Extracts parameters from the URL to be passed to handlers.
- `Method-specific routing`: Supports different HTTP methods (`GET`, `POST`, `PUT`, `DELETE`, etc.).
- `Subrouters`: Allows you to create groups of routes with shared middleware or other configurations.
- `Middleware support`: You can easily add custom middleware to routes.
- `Path matching`: Supports `wildcard` and path parameter matching.
