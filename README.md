# go-cli-tasks

A simple task manager application inspired by [Roadmap.sh](https://roadmap.sh/projects/task-tracker).

## Building the application

```bash
go build -o build/task-cli *.go
```

## Running the application

### Development

```bash
go run *.go
```

### Production

```bash
// Build the application
go build -o build/task-cli *.go

// Run the application
./build/task-cli
```
