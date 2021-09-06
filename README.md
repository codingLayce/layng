# Layng

Layng is a language written in Go, based on the book "Writing An Interpreter In Go" by Thorsten Ball.

I use it to keep playing with Go and learn how a simple language could work.

For now the language is only interpretable in a REPL.

## Run REPL
```
go run main.go
```

## Features Implemented

### Variables
You can store values inside variables.
```bash
>> let a = 5
>> a
5
>> let b = a
>> b
5
```
For now, when you do `let b = a` b isn't referencing a, it copies the value of a into b.
```bash
>> let a = 1
>> let b = a
>> b
1
>> let a = 2
>> a
2
>> b
1
```

For now, you cannot change the value of a variable `a = 6` will throw an error. But you can override the previous variable:
```
>> let a = 6
>> a
6
>> a = 5
parser errors:
    no prefix parse function for = found
>> let a = 5
>> a
5
```


### Integers
You can manipulate integers as you please. The supported operators are the following:
- \+ (Add)
- \- (Substract)
- \* (Multiply)
- / (Divide)

```
>> 5 + 5
10
>> 5 - 4
1
>> 5 * 5
25
>> 5 / 4
1
```
Notice that for now the division between 2 integers returns the floor value.

Layng know how to prioritize operators between intergers, example:
```
>> 5 + 2 * 10
25
>> (5 + 10 * 2 + 15 / 3) * 2 + -10
50
```

### Booleans
You can manipulate booleans.

```
>> true
true
>> false
false
```

### Comparators
In Layng you can compare the built-in types (booleans and integers):

```
>> 1 == 1
true
>> 1 != 1
false
>> 1 > 2
false
>> 1 < 2
true
>> true != false
true
>> !false
true
>> !!true
true
```

### Conditions
In Layng you can make conditions (for now only if/else statements are supported, not if/elseif/else).

Such as some modern languages, the if/else statement returns a value, example:
```bash
>> if (5 > 5) {1} else {2}
2
```

### Return
In Layng you can return values inside functions (not yet implemented) but also anywhere in your statements.
```bash
>> return 9; 10
9
```
Notice that the evaluator is stopping his evaluation when it encounters a return statement.

For more readability, you can explicitly indicates what a if/else statement is returning
```bash
>> let a = 5
>> if (a == 5) { return a } else { return false }
5
>> if (a == 4) { return a } else { return false }
false
```

## Features To Come
- Support floating point numbers
- Support modulo operator
- Support changing the value of a variable
- Support and/or in comparators
- Create better errors:
    - Display line number where the error occures
    - Display stack trace
- Create switch statement
- Create lambda expression
- Create more readable big integers (100000000 could be write 100_000_000)
- Create a simple standard library:
    - See what's missing in the Array/Hashes/Strings implementation