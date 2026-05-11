concept

structured data must be validated and converted into correct data types before being used

goal is to transform string fields into typed values and reject invalid data

How it works

input: []RawRecord
process:
- validate field count
- convert string fields into typed values (e.g. string → int)
- return error if conversion fails
output: Record (typed) or error

TEST

valid data -> converted into typed Record  
invalid field count -> rejected  
invalid data type (e.g. age not a number) -> rejected  

code pattern

age, err := strconv.Atoi(strings.TrimSpace(r.Fields[1]))
if err != nil {
    return Record{}, fmt.Errorf("invalid age")
}

record := Record{
    Name: strings.TrimSpace(r.Fields[0]),
    Age:  age,
    Role: strings.TrimSpace(r.Fields[2]),
}

System connection

structuring layer → validation layer → typed domain model

system now moves from raw structured data to safe, typed, and validated data

foundation for reliable processing and business logic
