# `fesl-codec`

**FESL Codec** is serialization library dedicated for handling exchange of messages between BFHeroes backend and game/server clients.

## Data Types

* `string`:

```txt
key=value
otherKey="value with spaces"
```

* `integer`:

```txt
key=123
```

* `boolean` is always represented by an integer value:

```txt
truish=1
false=0
```

* `float` is usually represented as 4-digits precision after the decimal point:

```txt
foo=1.0000
```

* empty values (_nil_ / _null_):

```txt
key=
```

* Arrays always adds additonal key (`%s.[]`) which contains number of declared items (array length).

```txt
keys.[]=4
keys.0=foo
keys.1=bar
keys.2=baz
keys.3=qux
```

* Maps (dictionaries) uses brackets to define the keys and also uses special key (`%s.{}`) which contains number of declared items (map capacity).

```txt
props.{foo}=0
props.{bar}=baz
props.{}=2
```

* Structs:

```txt
foo.id=1
foo.bar=baz
```
