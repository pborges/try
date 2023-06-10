# try

### An experiment for the `if err == nil` hatter crowd

## Example (full example in `cmd/main.go`)
**ASSUMING** `func Reverse(string) (string, error)` 

simply reverses a string

**ASSUMING** `RandomValidation(in try.Result[string]) try.Result[string]` 

throws an error if the input value is "hi"

```
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
    fmt.Println("PASS:")
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
