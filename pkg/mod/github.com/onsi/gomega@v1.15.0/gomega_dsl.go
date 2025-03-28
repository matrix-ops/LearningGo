/*
Gomega is the Ginkgo BDD-style testing framework's preferred matcher library.

The godoc documentation describes Gomega's API.  More comprehensive documentation (with examples!) is available at http://onsi.github.io/gomega/

Gomega on Github: http://github.com/onsi/gomega

Learn more about Ginkgo online: http://onsi.github.io/ginkgo

Ginkgo on Github: http://github.com/onsi/ginkgo

Gomega is MIT-Licensed
*/
package gomega

import (
	"errors"
	"fmt"
	"time"

	"github.com/onsi/gomega/internal"
	"github.com/onsi/gomega/types"
)

const GOMEGA_VERSION = "1.15.0"

const nilGomegaPanic = `You are trying to make an assertion, but haven't registered Gomega's fail handler.
If you're using Ginkgo then you probably forgot to put your assertion in an It().
Alternatively, you may have forgotten to register a fail handler with RegisterFailHandler() or RegisterTestingT().
Depending on your vendoring solution you may be inadvertently importing gomega and subpackages (e.g. ghhtp, gexec,...) from different locations.
`

// Gomega describes the essential Gomega DSL. This interface allows libraries
// to abstract between the standard package-level function implementations
// and alternatives like *WithT.
//
// The types in the top-level DSL have gotten a bit messy due to earlier depracations that avoid stuttering
// and due to an accidental use of a concrete type (*WithT) in an earlier release.
//
// As of 1.15 both the WithT and Ginkgo variants of Gomega are implemented by the same underlying object
// however one (the Ginkgo variant) is exported as an interface (types.Gomega) whereas the other (the withT variant)
// is shared as a concrete type (*WithT, which is aliased to *internal.Gomega).  1.15 did not clean this mess up to ensure
// that declarations of *WithT in existing code are not broken by the upgrade to 1.15.
type Gomega = types.Gomega

// DefaultGomega supplies the standard package-level implementation
var Default = Gomega(internal.NewGomega(internal.FetchDefaultDurationBundle()))

// NewGomega returns an instance of Gomega wired into the passed-in fail handler.
// You generally don't need to use this when using Ginkgo - RegisterFailHandler will wire up the global gomega
// However creating a NewGomega with a custom fail handler can be useful in contexts where you want to use Gomega's
// rich ecosystem of matchers without causing a test to fail.  For example, to aggregate a series of potential failures
// or for use in a non-test setting.
func NewGomega(fail types.GomegaFailHandler) Gomega {
	return internal.NewGomega(Default.(*internal.Gomega).DurationBundle).ConfigureWithFailHandler(fail)
}

// WithT wraps a *testing.T and provides `Expect`, `Eventually`, and `Consistently` methods.  This allows you to leverage
// Gomega's rich ecosystem of matchers in standard `testing` test suites.
//
// Use `NewWithT` to instantiate a `WithT`
//
// As of 1.15 both the WithT and Ginkgo variants of Gomega are implemented by the same underlying object
// however one (the Ginkgo variant) is exported as an interface (types.Gomega) whereas the other (the withT variant)
// is shared as a concrete type (*WithT, which is aliased to *internal.Gomega).  1.15 did not clean this mess up to ensure
// that declarations of *WithT in existing code are not broken by the upgrade to 1.15.
type WithT = internal.Gomega

// GomegaWithT is deprecated in favor of gomega.WithT, which does not stutter.
type GomegaWithT = WithT

// NewWithT takes a *testing.T and returngs a `gomega.WithT` allowing you to use `Expect`, `Eventually`, and `Consistently` along with
// Gomega's rich ecosystem of matchers in standard `testing` test suits.
//
//	func TestFarmHasCow(t *testing.T) {
//	    g := gomega.NewWithT(t)
//
//	    f := farm.New([]string{"Cow", "Horse"})
//	    g.Expect(f.HasCow()).To(BeTrue(), "Farm should have cow")
//	 }
func NewWithT(t types.GomegaTestingT) *WithT {
	return internal.NewGomega(Default.(*internal.Gomega).DurationBundle).ConfigureWithT(t)
}

// NewGomegaWithT is deprecated in favor of gomega.NewWithT, which does not stutter.
var NewGomegaWithT = NewWithT

// RegisterFailHandler connects Ginkgo to Gomega. When a matcher fails
// the fail handler passed into RegisterFailHandler is called.
func RegisterFailHandler(fail types.GomegaFailHandler) {
	Default.(*internal.Gomega).ConfigureWithFailHandler(fail)
}

// RegisterFailHandlerWithT is deprecated and will be removed in a future release.
// users should use RegisterFailHandler, or RegisterTestingT
func RegisterFailHandlerWithT(_ types.GomegaTestingT, fail types.GomegaFailHandler) {
	fmt.Println("RegisterFailHandlerWithT is deprecated.  Please use RegisterFailHandler or RegisterTestingT instead.")
	Default.(*internal.Gomega).ConfigureWithFailHandler(fail)
}

// RegisterTestingT connects Gomega to Golang's XUnit style
// Testing.T tests.  It is now deprecated and you should use NewWithT() instead to get a fresh instance of Gomega for each test.
func RegisterTestingT(t types.GomegaTestingT) {
	Default.(*internal.Gomega).ConfigureWithT(t)
}

// InterceptGomegaFailures runs a given callback and returns an array of
// failure messages generated by any Gomega assertions within the callback.
// Exeuction continues after the first failure allowing users to collect all failures
// in the callback.
//
// This is most useful when testing custom matchers, but can also be used to check
// on a value using a Gomega assertion without causing a test failure.
func InterceptGomegaFailures(f func()) []string {
	originalHandler := Default.(*internal.Gomega).Fail
	failures := []string{}
	Default.(*internal.Gomega).Fail = func(message string, callerSkip ...int) {
		failures = append(failures, message)
	}
	defer func() {
		Default.(*internal.Gomega).Fail = originalHandler
	}()
	f()
	return failures
}

// InterceptGomegaFailure runs a given callback and returns the first
// failure message generated by any Gomega assertions within the callback, wrapped in an error.
//
// The callback ceases execution as soon as the first failed assertion occurs, however Gomega
// does not register a failure with the FailHandler registered via RegisterFailHandler - it is up
// to the user to decide what to do with the returned error
func InterceptGomegaFailure(f func()) (err error) {
	originalHandler := Default.(*internal.Gomega).Fail
	Default.(*internal.Gomega).Fail = func(message string, callerSkip ...int) {
		err = errors.New(message)
		panic("stop execution")
	}

	defer func() {
		Default.(*internal.Gomega).Fail = originalHandler
		if e := recover(); e != nil {
			if err == nil {
				panic(e)
			}
		}
	}()

	f()
	return err
}

func ensureDefaultGomegaIsConfigured() {
	if !Default.(*internal.Gomega).IsConfigured() {
		panic(nilGomegaPanic)
	}
}

// Ω wraps an actual value allowing assertions to be made on it:
//
//	Ω("foo").Should(Equal("foo"))
//
// If Ω is passed more than one argument it will pass the *first* argument to the matcher.
// All subsequent arguments will be required to be nil/zero.
//
// This is convenient if you want to make an assertion on a method/function that returns
// a value and an error - a common patter in Go.
//
// For example, given a function with signature:
//
//	func MyAmazingThing() (int, error)
//
// Then:
//
//	Ω(MyAmazingThing()).Should(Equal(3))
//
// Will succeed only if `MyAmazingThing()` returns `(3, nil)`
//
// Ω and Expect are identical
func Ω(actual interface{}, extra ...interface{}) Assertion {
	ensureDefaultGomegaIsConfigured()
	return Default.Ω(actual, extra...)
}

// Expect wraps an actual value allowing assertions to be made on it:
//
//	Expect("foo").To(Equal("foo"))
//
// If Expect is passed more than one argument it will pass the *first* argument to the matcher.
// All subsequent arguments will be required to be nil/zero.
//
// This is convenient if you want to make an assertion on a method/function that returns
// a value and an error - a common patter in Go.
//
// For example, given a function with signature:
//
//	func MyAmazingThing() (int, error)
//
// Then:
//
//	Expect(MyAmazingThing()).Should(Equal(3))
//
// Will succeed only if `MyAmazingThing()` returns `(3, nil)`
//
// Expect and Ω are identical
func Expect(actual interface{}, extra ...interface{}) Assertion {
	ensureDefaultGomegaIsConfigured()
	return Default.Expect(actual, extra...)
}

// ExpectWithOffset wraps an actual value allowing assertions to be made on it:
//
//	ExpectWithOffset(1, "foo").To(Equal("foo"))
//
// Unlike `Expect` and `Ω`, `ExpectWithOffset` takes an additional integer argument
// that is used to modify the call-stack offset when computing line numbers.
//
// This is most useful in helper functions that make assertions.  If you want Gomega's
// error message to refer to the calling line in the test (as opposed to the line in the helper function)
// set the first argument of `ExpectWithOffset` appropriately.
func ExpectWithOffset(offset int, actual interface{}, extra ...interface{}) Assertion {
	ensureDefaultGomegaIsConfigured()
	return Default.ExpectWithOffset(offset, actual, extra...)
}

/*
Eventually enables making assertions on asynchronous behavior.

Eventually checks that an assertion *eventually* passes.  Eventually blocks when called and attempts an assertion periodically until it passes or a timeout occurs.  Both the timeout and polling interval are configurable as optional arguments.
The first optional argument is the timeout (which defaults to 1s), the second is the polling interval (which defaults to 10ms).  Both intervals can be specified as time.Duration, parsable duration strings or floats/integers (in which case they are interpreted as seconds).

Eventually works with any Gomega compatible matcher and supports making assertions against three categories of actual value:

**Category 1: Making Eventually assertions on values**

There are several examples of values that can change over time.  These can be passed in to Eventually and will be passed to the matcher repeatedly until a match occurs.  For example:

	c := make(chan bool)
	go DoStuff(c)
	Eventually(c, "50ms").Should(BeClosed())

will poll the channel repeatedly until it is closed.  In this example `Eventually` will block until either the specified timeout of 50ms has elapsed or the channel is closed, whichever comes first.

Several Gomega libraries allow you to use Eventually in this way.  For example, the gomega/gexec package allows you to block until a *gexec.Session exits successfuly via:

	Eventually(session).Should(gexec.Exit(0))

And the gomega/gbytes package allows you to monitor a streaming *gbytes.Buffer until a given string is seen:

	Eventually(buffer).Should(gbytes.Say("hello there"))

In these examples, both `session` and `buffer` are designed to be thread-safe when polled by the `Exit` and `Say` matchers.  This is not true in general of most raw values, so while it is tempting to do something like:

	// THIS IS NOT THREAD-SAFE
	var s *string
	go mutateStringEventually(s)
	Eventually(s).Should(Equal("I've changed"))

this will trigger Go's race detector as the goroutine polling via Eventually will race over the value of s with the goroutine mutating the string.  For cases like this you can use channels or introduce your own locking around s by passing Eventually a function.

**Category 2: Make Eventually assertions on functions**

Eventually can be passed functions that **take no arguments** and **return at least one value**.  When configured this way, Eventually will poll the function repeatedly and pass the first returned value to the matcher.

For example:

	   Eventually(func() int {
	   	return client.FetchCount()
	   }).Should(BeNumerically(">=", 17))

	will repeatedly poll client.FetchCount until the BeNumerically matcher is satisfied.  (Note that this example could have been written as Eventually(client.FetchCount).Should(BeNumerically(">=", 17)))

If multple values are returned by the function, Eventually will pass the first value to the matcher and require that all others are zero-valued.  This allows you to pass Eventually a function that returns a value and an error - a common patternin Go.

For example, consider a method that returns a value and an error:

	func FetchFromDB() (string, error)

Then

	Eventually(FetchFromDB).Should(Equal("got it"))

will pass only if and when the returned error is nil *and* the returned string satisfies the matcher.

It is important to note that the function passed into Eventually is invoked *synchronously* when polled.  Eventually does not (in fact, it cannot) kill the function if it takes longer to return than Eventually's configured timeout.  You should design your functions with this in mind.

**Category 3: Making assertions _in_ the function passed into Eventually**

When testing complex systems it can be valuable to assert that a _set_ of assertions passes Eventually.  Eventually supports this by accepting functions that take a single Gomega argument and return zero or more values.

Here's an example that makes some asssertions and returns a value and error:

	Eventually(func(g Gomega) (Widget, error) {
		ids, err := client.FetchIDs()
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(ids).To(ContainElement(1138))
		return client.FetchWidget(1138)
	}).Should(Equal(expectedWidget))

will pass only if all the assertions in the polled function pass and the return value satisfied the matcher.

Eventually also supports a special case polling function that takes a single Gomega argument and returns no values.  Eventually assumes such a function is making assertions and is designed to work with the Succeed matcher to validate that all assertions have passed.
For example:

	Eventually(func(g Gomega) {
		model, err := client.Find(1138)
		g.Expect(err).NotTo(HaveOccurred())
		g.Expect(model.Reticulate()).To(Succeed())
		g.Expect(model.IsReticulated()).To(BeTrue())
		g.Expect(model.Save()).To(Succeed())
	}).Should(Succeed())

will rerun the function until all assertions pass.
*/
func Eventually(actual interface{}, intervals ...interface{}) AsyncAssertion {
	ensureDefaultGomegaIsConfigured()
	return Default.Eventually(actual, intervals...)
}

// EventuallyWithOffset operates like Eventually but takes an additional
// initial argument to indicate an offset in the call stack.  This is useful when building helper
// functions that contain matchers.  To learn more, read about `ExpectWithOffset`.
func EventuallyWithOffset(offset int, actual interface{}, intervals ...interface{}) AsyncAssertion {
	ensureDefaultGomegaIsConfigured()
	return Default.EventuallyWithOffset(offset, actual, intervals...)
}

/*
Consistently, like Eventually, enables making assertions on asynchronous behavior.

Consistently blocks when called for a specified duration.  During that duration Consistently repeatedly polls its matcher and ensures that it is satisfied.  If the matcher is consistently satisfied, then Consistently will pass.  Otherwise Consistently will fail.

Both the total waiting duration and the polling interval are configurable as optional arguments.  The first optional arugment is the duration that Consistently will run for (defaults to 100ms), and the second argument is the polling interval (defaults to 10ms).  As with Eventually, these intervals can be passed in as time.Duration, parsable duration strings or an integer or float number of seconds.

Consistently accepts the same three categories of actual as Eventually, check the Eventually docs to learn more.

Consistently is useful in cases where you want to assert that something *does not happen* for a period of time.  For example, you may want to assert that a goroutine does *not* send data down a channel.  In this case you could write:

	Consistently(channel, "200ms").ShouldNot(Receive())

This will block for 200 milliseconds and repeatedly check the channel and ensure nothing has been received.
*/
func Consistently(actual interface{}, intervals ...interface{}) AsyncAssertion {
	ensureDefaultGomegaIsConfigured()
	return Default.Consistently(actual, intervals...)
}

// ConsistentlyWithOffset operates like Consistently but takes an additional
// initial argument to indicate an offset in the call stack. This is useful when building helper
// functions that contain matchers. To learn more, read about `ExpectWithOffset`.
func ConsistentlyWithOffset(offset int, actual interface{}, intervals ...interface{}) AsyncAssertion {
	ensureDefaultGomegaIsConfigured()
	return Default.ConsistentlyWithOffset(offset, actual, intervals...)
}

// SetDefaultEventuallyTimeout sets the default timeout duration for Eventually. Eventually will repeatedly poll your condition until it succeeds, or until this timeout elapses.
func SetDefaultEventuallyTimeout(t time.Duration) {
	Default.SetDefaultEventuallyTimeout(t)
}

// SetDefaultEventuallyPollingInterval sets the default polling interval for Eventually.
func SetDefaultEventuallyPollingInterval(t time.Duration) {
	Default.SetDefaultEventuallyPollingInterval(t)
}

// SetDefaultConsistentlyDuration sets  the default duration for Consistently. Consistently will verify that your condition is satisfied for this long.
func SetDefaultConsistentlyDuration(t time.Duration) {
	Default.SetDefaultConsistentlyDuration(t)
}

// SetDefaultConsistentlyPollingInterval sets the default polling interval for Consistently.
func SetDefaultConsistentlyPollingInterval(t time.Duration) {
	Default.SetDefaultConsistentlyPollingInterval(t)
}

// AsyncAssertion is returned by Eventually and Consistently and polls the actual value passed into Eventually against
// the matcher passed to the Should and ShouldNot methods.
//
// Both Should and ShouldNot take a variadic optionalDescription argument.
// This argument allows you to make your failure messages more descriptive.
// If a single argument of type `func() string` is passed, this function will be lazily evaluated if a failure occurs
// and the returned string is used to annotate the failure message.
// Otherwise, this argument is passed on to fmt.Sprintf() and then used to annotate the failure message.
//
// Both Should and ShouldNot return a boolean that is true if the assertion passed and false if it failed.
//
// Example:
//
//	Eventually(myChannel).Should(Receive(), "Something should have come down the pipe.")
//	Consistently(myChannel).ShouldNot(Receive(), func() string { return "Nothing should have come down the pipe." })
type AsyncAssertion = types.AsyncAssertion

// GomegaAsyncAssertion is deprecated in favor of AsyncAssertion, which does not stutter.
type GomegaAsyncAssertion = types.AsyncAssertion

// Assertion is returned by Ω and Expect and compares the actual value to the matcher
// passed to the Should/ShouldNot and To/ToNot/NotTo methods.
//
// Typically Should/ShouldNot are used with Ω and To/ToNot/NotTo are used with Expect
// though this is not enforced.
//
// All methods take a variadic optionalDescription argument.
// This argument allows you to make your failure messages more descriptive.
// If a single argument of type `func() string` is passed, this function will be lazily evaluated if a failure occurs
// and the returned string is used to annotate the failure message.
// Otherwise, this argument is passed on to fmt.Sprintf() and then used to annotate the failure message.
//
// All methods return a bool that is true if the assertion passed and false if it failed.
//
// Example:
//
//	Ω(farm.HasCow()).Should(BeTrue(), "Farm %v should have a cow", farm)
type Assertion = types.Assertion

// GomegaAssertion is deprecated in favor of Assertion, which does not stutter.
type GomegaAssertion = types.Assertion

// OmegaMatcher is deprecated in favor of the better-named and better-organized types.GomegaMatcher but sticks around to support existing code that uses it
type OmegaMatcher = types.GomegaMatcher
