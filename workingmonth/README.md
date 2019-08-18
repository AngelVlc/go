# workingmonth
--
    import "github.com/AngelVlc/go/workingmonth"

Package workingmonth contains methods to ge the number of working hours in a
given month

## Usage

#### type WorkingMonth

```go
type WorkingMonth struct {
	Year  int
	Month int
}
```

WorkingMonth contains methods to get the working hours in a given month

#### func (WorkingMonth) TotalDays

```go
func (m WorkingMonth) TotalDays() int
```
TotalDays returns the number of days in the month

#### func (WorkingMonth) WorkingDays

```go
func (m WorkingMonth) WorkingDays() int
```
WorkingDays returns the number of working days in the month

#### func (WorkingMonth) WorkingHours

```go
func (m WorkingMonth) WorkingHours() int
```
WorkingHours returns the number of working hours in the month

## Testing

```
go test
go test -bench .
```

```
go test -cover
go test -cover -coverprofile c.out
go tool cover -html=c.out
```

## Update README

https://github.com/robertkrimen/godocdown

```
godocdown path_to_package
```