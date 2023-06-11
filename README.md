# golang-course

GO Lang course from Nana

Link: https://www.youtube.com/watch?v=yyUHQIec83I&ab_channel=TechWorldwithNana

## How to use

To run the application use:

```bash
go run main.go
```

### Notes

#### Packaging

- To run an application with more packages (e.g. `shared.go`) use:

```bash
go run main.go shared.go
```

, otherwise the application does not know, from where the functions from `shared.go` are imported.

- To trick this you cau also run the app directory-wide so you don't have to specify the files at all

```bash
go run .
```
