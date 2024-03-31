# Golang struct tags gotchas

URL: https://suraj.io/post/golang-struct-tags-space/
Category: Golang

---

In golang while using struct tag, the spaces make a lot of difference. For example look at the following code.

```
type PodStatus struct {
	Status string `json: ",status"`
}

```

If you run `[go vet](https://golang.org/cmd/vet/)` on this piece of code you will get following error:

```
$ go vet types.go
# command-line-arguments
./types.go:28: struct field tag `json: ",status"`
not compatible with reflect.StructTag.Get: bad syntax for struct tag value

```

Now this does not tell us what is wrong with the struct tag, `json: ",status"`. The problem with this struct tag is that the extra space can be interpreted as delimiter so provide key-value pair without space.

So if the struct changes from:

```
`json: ",status"`

```

to

So the change is just the space after `json:`, now we don’t see the error.

More information about the struct tags can be found in this elaborated blog post named: *[Tags in Golang](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e)*.