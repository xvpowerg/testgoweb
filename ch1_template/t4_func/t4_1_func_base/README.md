# Formatting with time

Since MST is GMT-0700, the reference time can be thought of as:

01/02 03:04:05PM '06 -0700

(January 02, 2006)

[godoc.org/time](https://godoc.org/time#pkg-constants)

***

## func (Time) Format

```go
func (t Time) Format(layout string) string
```

Format returns a textual representation of the time value formatted according to layout, which defines the format by showing how the reference time, defined to be

```go
Mon Jan 2 15:04:05 -0700 MST 2006
```

would be displayed if it were the value; it serves as an example of the desired output. The same display rules will then be applied to the time value.

A fractional second is represented by adding a period and zeros to the end of the seconds section of layout string, as in "15:04:05.000" to format a time stamp with millisecond precision.

Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard and convenient representations of the reference time. For more information about the formats and the definition of the reference time, see the documentation for ANSIC and the other constants defined by this package.