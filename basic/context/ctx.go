package main

import (
	"context"
	"fmt"
	"time"
)

/*
	context 上下文，作为函数参数传递控制上下文

	context 不要放在 struct 类型中，而应该作为函数的第一个参数显示传递。
	不要传递 nil context ，如果不确定使用哪个上下文请传递 context.TODO 。
	context 仅用于转换进程和API的 request-scoped 数据，而不是作为可选参数传递给函数。
	可以将相同的 context 传递给不同 goroutine 中运行的函数，context 对于多个 goroutine 同时使用时安全的。
	https://blog.golang.org/context
*/
func main() {
	/*
		func Background() Context : 返回一个空 context，作为 root context 使用
		func TODO() Context : 返回一个空 context，实际不使用或不知道使用哪个 context 时调用
		func WithValue(parent Context, key, val interface{}) Context : 上下文传值，key 必须是自定义的类型，避免与内置类型冲突
	*/
	// func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel() // 即使 context 会过期，仍建议显示调用 cancel

	select { // select 监听 channel
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done(): // context 到期
		fmt.Println(ctx.Err())
	}

	// ----------------------------------------------
	// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
	// WithTimeout 也是返回 WithDeadline 的调用
	ctx2, cancel2 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel2()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx2.Done(): // context 超时
		fmt.Println(ctx2.Err())
	}

	// ----------------------------------------------
	// func WithValue(parent Context, key, val interface{}) Context
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}
		fmt.Println("key not found: ", k)
	}

	k := favContextKey("language")
	ctx3 := context.WithValue(context.Background(), k, "Go")

	f(ctx3, k)
	f(ctx3, favContextKey("color"))
}

/*

type Context interface {
        // Deadline returns the time when work done on behalf of this context
        // should be canceled. Deadline returns ok==false when no deadline is
        // set. Successive calls to Deadline return the same results.
        Deadline() (deadline time.Time, ok bool)

        // Done returns a channel that's closed when work done on behalf of this
        // context should be canceled. Done may return nil if this context can
        // never be canceled. Successive calls to Done return the same value.
        //
        // WithCancel arranges for Done to be closed when cancel is called;
        // WithDeadline arranges for Done to be closed when the deadline
        // expires; WithTimeout arranges for Done to be closed when the timeout
        // elapses.
        //
        // Done is provided for use in select statements:
        //
        //  // Stream generates values with DoSomething and sends them to out
        //  // until DoSomething returns an error or ctx.Done is closed.
        //  func Stream(ctx context.Context, out chan<- Value) error {
        //  	for {
        //  		v, err := DoSomething(ctx)
        //  		if err != nil {
        //  			return err
        //  		}
        //  		select {
        //  		case <-ctx.Done():
        //  			return ctx.Err()
        //  		case out <- v:
        //  		}
        //  	}
        //  }
        //
        // See https://blog.golang.org/pipelines for more examples of how to use
        // a Done channel for cancelation.
        Done() <-chan struct{}

        // If Done is not yet closed, Err returns nil.
        // If Done is closed, Err returns a non-nil error explaining why:
        // Canceled if the context was canceled
        // or DeadlineExceeded if the context's deadline passed.
        // After Err returns a non-nil error, successive calls to Err return the same error.
        Err() error

        // Value returns the value associated with this context for key, or nil
        // if no value is associated with key. Successive calls to Value with
        // the same key returns the same result.
        //
        // Use context values only for request-scoped data that transits
        // processes and API boundaries, not for passing optional parameters to
        // functions.
        //
        // A key identifies a specific value in a Context. Functions that wish
        // to store values in Context typically allocate a key in a global
        // variable then use that key as the argument to context.WithValue and
        // Context.Value. A key can be any type that supports equality;
        // packages should define keys as an unexported type to avoid
        // collisions.
        //
        // Packages that define a Context key should provide type-safe accessors
        // for the values stored using that key:
        //
        // 	// Package user defines a User type that's stored in Contexts.
        // 	package user
        //
        // 	import "context"
        //
        // 	// User is the type of value stored in the Contexts.
        // 	type User struct {...}
        //
        // 	// key is an unexported type for keys defined in this package.
        // 	// This prevents collisions with keys defined in other packages.
        // 	type key int
        //
        // 	// userKey is the key for user.User values in Contexts. It is
        // 	// unexported; clients use user.NewContext and user.FromContext
        // 	// instead of using this key directly.
        // 	var userKey key
        //
        // 	// NewContext returns a new Context that carries value u.
        // 	func NewContext(ctx context.Context, u *User) context.Context {
        // 		return context.WithValue(ctx, userKey, u)
        // 	}
        //
        // 	// FromContext returns the User value stored in ctx, if any.
        // 	func FromContext(ctx context.Context) (*User, bool) {
        // 		u, ok := ctx.Value(userKey).(*User)
        // 		return u, ok
        // 	}
        Value(key interface{}) interface{}
}

*/
