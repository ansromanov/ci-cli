# Go Development Rules for Cursor

## Code Style and Formatting

- Always follow Go's official formatting standards using `gofmt` or `go fmt`
- Use 4-space indentation (tabs in Go)
- Keep line length under 120 characters
- Use camelCase for variable and function names
- Use PascalCase for exported functions, types, and constants
- Use UPPER_CASE for package-level constants
- Use descriptive names that are clear and concise

## Project Structure

- Follow standard Go project layout:
  - `cmd/` for main applications
  - `internal/` for private application code
  - `pkg/` for public libraries
  - `api/` for API definitions
  - `web/` for web assets
  - `configs/` for configuration files
  - `test/` for additional test files
  - `docs/` for documentation
  - `scripts/` for build and deployment scripts
  - `build/` for packaging and CI/CD
  - `deployments/` for deployment configurations
  - `third_party/` for third-party tools and code

## Package Organization

- Keep packages focused and cohesive
- Use meaningful package names that describe their purpose
- Avoid package names that conflict with standard library
- Group related functionality in the same package
- Use `internal/` for code that should not be imported by other projects

## Error Handling

- Always check and handle errors explicitly
- Use `if err != nil` pattern consistently
- Return errors rather than panicking in most cases
- Use `errors.Wrap` or `fmt.Errorf` to add context to errors
- Consider using custom error types for specific error cases
- Use `errors.Is` and `errors.As` for error comparison and type assertion

## Interfaces

- Keep interfaces small and focused (interface segregation principle)
- Define interfaces where they are used, not where they are implemented
- Use interfaces for behavior, not data
- Prefer composition over inheritance
- Use `io.Reader`, `io.Writer`, and similar standard interfaces when possible

## Testing

- Write tests for all exported functions
- Use table-driven tests for multiple test cases
- Name test functions as `TestFunctionName_Scenario`
- Use `t.Run()` for subtests
- Use `testing.T` for unit tests, `testing.B` for benchmarks
- Mock external dependencies using interfaces
- Use `testify` package for assertions when needed
- Aim for high test coverage (80%+)

## Documentation

- Write godoc comments for all exported functions, types, and packages
- Start comments with the name of the thing being documented
- Use complete sentences
- Provide examples in documentation when helpful
- Keep README.md up to date with project setup and usage instructions

## Dependencies

- Use Go modules for dependency management
- Keep dependencies minimal and up to date
- Use `go mod tidy` to clean up unused dependencies
- Pin dependency versions in go.mod
- Use `go.sum` for dependency verification
- Consider using `go mod vendor` for reproducible builds

## Performance

- Use `sync.Pool` for frequently allocated objects
- Use buffered channels when appropriate
- Avoid unnecessary allocations in hot paths
- Use `strings.Builder` for string concatenation in loops
- Profile code before optimizing
- Use `pprof` for performance analysis

## Concurrency

- Use goroutines for concurrent operations
- Use channels for communication between goroutines
- Use `sync.WaitGroup` to wait for goroutines to complete
- Use `context.Context` for cancellation and timeouts
- Avoid goroutine leaks by ensuring proper cleanup
- Use `select` for non-blocking channel operations
- Use `sync.Mutex` or `sync.RWMutex` for shared state protection

## Configuration

- Use environment variables for configuration
- Use struct tags for configuration mapping
- Validate configuration on startup
- Use default values for optional configuration
- Use configuration files (YAML, JSON) for complex configurations

## Logging

- Use structured logging with fields
- Use appropriate log levels (DEBUG, INFO, WARN, ERROR)
- Include context in log messages
- Use correlation IDs for request tracing
- Avoid logging sensitive information

## Security

- Validate all input data
- Use HTTPS for all external communications
- Use secure random number generation (`crypto/rand`)
- Hash passwords using bcrypt or similar
- Use prepared statements for database queries
- Sanitize user input to prevent injection attacks

## API Design

- Use RESTful principles for HTTP APIs
- Use consistent URL patterns
- Return appropriate HTTP status codes
- Use JSON for request/response bodies
- Include error details in error responses
- Use versioning for APIs (e.g., `/v1/`, `/v2/`)

## Database

- Use prepared statements to prevent SQL injection
- Use transactions for multi-step operations
- Handle database connection errors gracefully
- Use connection pooling
- Use migrations for database schema changes
- Consider using an ORM like GORM for complex queries

## Code Generation

- Use `go generate` for code generation
- Document generated code clearly
- Keep generated code in version control
- Use tools like `stringer`, `mockgen`, or `protoc` as needed

## Common Patterns

- Use constructor functions for complex struct initialization
- Use functional options pattern for flexible configuration
- Use builder pattern for complex object construction
- Use middleware pattern for HTTP handlers
- Use dependency injection for testability

## Linting and Static Analysis

- Use `golangci-lint` for comprehensive linting
- Enable all relevant linters
- Fix all linting issues before committing
- Use `go vet` for additional static analysis
- Consider using `staticcheck` for advanced analysis

## Build and Deployment

- Use multi-stage Docker builds
- Build for multiple platforms when needed
- Use semantic versioning for releases
- Use CI/CD pipelines for automated testing and deployment
- Use health checks for deployed services
- Use graceful shutdown patterns

## Monitoring and Observability

- Add metrics using Prometheus client
- Add distributed tracing with OpenTelemetry
- Use structured logging for better observability
- Add health check endpoints
- Monitor resource usage (CPU, memory, disk)

## Code Review Guidelines

- Review for correctness, performance, and security
- Check error handling
- Verify test coverage
- Ensure proper documentation
- Look for potential race conditions
- Check for memory leaks
- Verify API compatibility

## When Writing Code

- Think about the user of your code
- Write self-documenting code
- Keep functions small and focused
- Use meaningful variable names
- Add comments for complex logic
- Consider edge cases and error conditions
- Write code that is easy to test
- Follow the principle of least surprise
