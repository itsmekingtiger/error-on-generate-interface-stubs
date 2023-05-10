# Summary

This issue addresses an undesired behaviour in VS Code's Go extension when implementing an interface with a generic type.



1. When the `Go: Generate Interface Stubs` command is run, it raises a parsing error.
2. When the Code Action is run, it does not infer the concrete type.



Consider the following code where we aim to implement `IUserRepo` on `UserRepo`.

```go
type IUser interface{}

type BaseRepo[T any, ID any] interface {
	Create(T) (ID, error)
	ReteriveById(ID) (T, error)
	DeleteById(ID) error
}

type IUserRepo interface {
	BaseRepo[IUser, int]

	GetFollows(IUser) ([]IUser, error)
	Follow(IUser, IUser) error
}

var _ IUserRepo = (*UserRepo)(nil)

type UserRepo struct{}
```



## Command `Go: Generate Interface Stubs`

When we run the command with `u *UserRepo IUserRepo`, we expect the following output:

```go
func (*UserRepo) Create(_ domain.IUser) (int, error) {
	panic("unimplemented")
}

func (*UserRepo) DeleteById(_ int) error {
	panic("unimplemented")
}

func (*UserRepo) ReteriveById(_ int) (domain.IUser, error) {
	panic("unimplemented")
}

func (*UserRepo) Follow(_ domain.IUser, _ domain.IUser) error {
	panic("unimplemented")
}

func (*UserRepo) GetFollows(_ domain.IUser) ([]domain.IUser, error) {
	panic("unimplemented")
}
```


However, it results in an error:

```
Cannot stub interface: couldn't parse interface: repository.CrdRepo[user.IUser, int]
```





## Code Action Type Inference

*expected*

```go
// Create implements IUserRepo
func (*UserRepo) Create(domain.IUser) (int, error) {
	panic("unimplemented")
}

// DeleteById implements IUserRepo
func (*UserRepo) DeleteById(int) error {
	panic("unimplemented")
}

// ReteriveById implements IUserRepo
func (*UserRepo) ReteriveById(int) (domain.IUser, error) {
	panic("unimplemented")
}

// Follow implements IUserRepo
func (*UserRepo) Follow(domain.IUser, domain.IUser) error {
	panic("unimplemented")
}

// GetFollows implements IUserRepo
func (*UserRepo) GetFollows(domain.IUser) ([]domain.IUser, error) {
	panic("unimplemented")
}
```



*actual*

```go
// Create implements IUserRepo
func (*UserRepo) Create(T) (ID, error) {
	panic("unimplemented")
}

// DeleteById implements IUserRepo
func (*UserRepo) DeleteById(ID) error {
	panic("unimplemented")
}

// ReteriveById implements IUserRepo
func (*UserRepo) ReteriveById(ID) (T, error) {
	panic("unimplemented")
}

// Follow implements IUserRepo
func (*UserRepo) Follow(domain.IUser, domain.IUser) error {
	panic("unimplemented")
}

// GetFollows implements IUserRepo
func (*UserRepo) GetFollows(domain.IUser) ([]domain.IUser, error) {
	panic("unimplemented")
}
```
