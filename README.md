# try

### An experiment for the `if err == nil` hatter crowd

## Example (full example in `cmd/main.go`)
**ASSUMING** `func Reverse(string) (string, error)` 

simply reverses a string

**ASSUMING** `RandomValidation(try.Result[string]) try.Result[string]` 

throws an error if the input value is "hi"

## Processing example

```go
func main(){
    for _, data := []string{"44","ih"} {
        step1 := try.To(Reverse)(try.Pass(data))
        step2 := RandomValidation(step1)
        step3 := try.To(strconv.Atoi)(step2)
        
        switch res := step3.(type) {
        case try.Ok[int]:
            fmt.Println("I did it!:", res.Value)
        case try.Error:
            fmt.Println("I failed:", res.Value)
        }
    }
}
```

## Output
```
I did it!: 44
I failed: no hi. only zulu
```

## Dataset processing example

```go
func main(){
    // Normal lookin data
    data := []string{"2", "4444", "44", "44fas", "ih"}
    
    // Do a series of things to said data
    step1 := try.Map(try.To(Reverse))(try.PassSlice(data))
    step2 := try.Map(RandomValidation)(step1)
    step3 := try.Map(try.To(strconv.Atoi))(step2)
    
    // Collect outputs
    pass, fail := try.Collect(step3)
    
    // Print outptus
    fmt.Println("Pass:")
    for _, e := range pass {
        fmt.Println("  ", e)
    }
    fmt.Println("Fail:")
    for _, e := range fail {
        fmt.Println("  ", e.Error())
    }
}
```

## Output
```
PASS:
   2
   44
Fail:
   no 4 letter words
   strconv.Atoi: parsing "saf44": invalid syntax
   no hi. only zulu
```
