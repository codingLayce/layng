package object

/* The Environment systemkeep track of variables.
 * An Environment can be nested in another one, this means that an Environment has inner variables and outer variables.
 * Keep in mind that when trying to retrieve a variable, it will first check for inner one and then outer one.
 *
 * REMINDER : You can create a variable with the same name than one in the parent, but the one used will be the one from the child.
 *
 * >> let a = 5;
 * >> let add = fn(a) { return a + a; }
 * >> add(2)
 * 4
 */

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Reassign(name string, val Object) bool {
	if _, ok := e.store[name]; !ok {
		return false
	}
	e.store[name] = val

	return true
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
