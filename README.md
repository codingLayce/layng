# Layng

Layng is a language written in Go, based on the book "Writing An Interpreter In Go" by Thorsten Ball.

I'am making it to keep playing with Go (because I really like with language) and learn how a simple language could work.

For now the language is only interpretable in a REPL.

## Install and Run REPL
```
git clone https://github.com/codingLayce/layng.git
go mod download
go build
./layng
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
The language variables are only passed as value, so if you assign a to b and then modify a, the value of b will still be the a value before the modification.
```bash
>> let a = 1
>> let b = a
>> b
1
>> a = 2
>> a
2
>> b
1
```
> Notice that you can reassign a new value to a variable with the statement `a = 2`.

### Strings
You can manipulate string as you please.
```bash
>> let a = "John"
>> a
John
>> let b = "Smith"
>> let c = a .. " " .. b
>> c
John Smith
```
> You can notice that string concatenation works with a double dot

### Integers
You can manipulate integers as you please. The supported operators are the following:
- \+ (Add)
- \- (Substract)
- \* (Multiply)
- / (Divide)

```bash
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
```bash
>> 5 + 2 * 10
25
>> (5 + 10 * 2 + 15 / 3) * 2 + -10
50
```

### Booleans
You can manipulate booleans.

```bash
>> true
true
>> false
false
```

### Comparators
In Layng you can compare the built-in types (booleans and integers):

```bash
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

### Functions
In Layng you can create functions and assign it to a variable in order to use it later.
```bash
>> let add = fn(x, y) {return x+y}
>> add(1, 2)
3
```

You can also use expression as argument:
```bash
>> let add = fn(x, y) {return x+y}
>> add((5+5), 2)
12
```

Or more complex one:
```bash
>> let add = fn(x, y) {return x+y}
>> let sub = fn(x ,y) {return x-y}
>> sub(add(1, 2), sub(10, 5))
-2
```

It can handle recursivity:
```bash
>> let loop = fn(x) {if (x == 0) {return -1} else {return loop(x-1)}}
>> loop(10)
-1
```
Here the function is called 10 times before it actualy returns a value without making another call to itself.

### Closures
In Layng functions can return function, but more importantly they have separeted scopes and keep track over time of their own arguments.
```bash
>> let createAdder = fn(x) { return fn(y) { return x + y } }
```
Here we create a function `createAdder` that takes as argument the left value of the addition. Then it returns another function that is making the addition.

```bash
>> let addTwo = createAdder(2)
```
Here we are calling the `createAdder` function previously declared and store his result into addTwo. So at this point, the value of `addTwo` is `fn(y) { return 2 + y }`. Notice that the `x` identifier is now replaced by the value given as argument of `createAdder`.

```bash
>> addTwo(5)
7
>> addTwo(3)
5
```

### Return
In Layng you can return values inside functions but also anywhere in your statements.
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

### Builtin functions
#### _len_
Description: Return the number of characters of the given string

Arguments: string

Return: int

Example: 
```bash
>> len("four")
4
>> let a = "Bonjour !"
>> len(a)
9
```

#### _string_
Description: Convert etiher a boolean or an integer into a string

Arguments: Object (BOOLEAN, INTEGER)

Return: string

Example:
```bash
>> string(42)
42
>> string(false)
false
>> let age = 21
>> string(age)
21
>> "Mon age: " + 21
Error: type mismatch: STRING + INTEGER
>> "Mon age: " + string(21)
Mon age: 21
```

#### _int_
Description: Convert a string into an int

Arguments: string

Return: int

Example:
```bash
>> int("42")
42
>> let a = "21"
>> int(a)
21
```

#### _bool_
Description: Convert a string into a boolean

Arguments: string

Return: boolean

Example:
```bash
>> bool("true")
true
>> let a = "false"
>> int(a)
false
```

#### _type_
Description: Return the type of the given expression

Arguments: Object

Return: string

Example:
```bash
>> type("test")
STRING
>> type(42)
INTEGER
>> type(false)
BOOLEAN
```

## My Own Features Implemented
- [ ] Support floating point numbers
- [ ] Support modulo operator
- [x] Support changing the value of a variable. `let a = 1; a = 2` the value of a is now 2. `b = 2` will throw an error because b isn't defined.
- [x] Pohibit changing the type of a variable. `let a = 1; a = "Bonjour"` shouldn't be possible.
- [ ] Support and/or in comparators
- [ ] Support elseif in conditions
- [ ] Support escaping character in string
- [ ] Support loop (for/while or a combined one)
- [ ] Support enum
- [ ] Support struct
- [x] Add convertion between types with builtin functions `string`, `int`, `bool`
- [x] Add builtin function to return the type of a given expression: `type`
- [ ] For now, functions can only be assigned to variables `let add = fn(x, y) {return x+y}`, allow to create global function without assigning it to a variable `fn add(x, y) {return x+y}`
- Create better errors:
    - [ ] Display line number where the error occures
    - [ ] Display stack trace
- [ ] Create switch statement
- [ ] Create more readable big integers (100000000 could be write 100_000_000)
- Create a simple standard library:
    - [ ] See what's missing in the Array/Hashes/Strings implementation