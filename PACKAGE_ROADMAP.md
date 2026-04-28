# lx Repository - Package Expansion Roadmap

> Strategic plan for expanding the lx library with focused, idiomatic packages

**Last Updated**: February 28, 2026

---

## 📦 Current Package Inventory

| Package | Status | Purpose | Completeness |
|---------|--------|---------|--------------|
| `lxconstraints` | ✅ Stable | Generic type constraints | ~90% |
| `lxptrs` | ✅ Stable | Pointer utilities | ~60% |
| `lxslices` | ✅ Stable | Slice operations | ~85% |
| `lxstrings` | ✅ Stable | String utilities | ~90% |
| `lxsystems` | ✅ Stable | System information | ~70% |
| `lxtuples` | ✅ Stable | Tuple types | ~40% |

---

## 🎯 Proposed New Packages

### **TIER 1: High Value, Essential Utilities** (Implement First)

####  **`lxmath`** - Mathematical Utilities
**Priority**: 🔥 HIGH  
**Rationale**: Common math operations missing from stdlib

**Core Functions**:
```go
// Basic operations
func Abs[T lxconstraints.Number](n T) T
func Sign[T lxconstraints.Number](n T) int
func Clamp[T lxconstraints.Ordered](value, min, max T) T
func InRange[T lxconstraints.Ordered](value, min, max T) bool

// Rounding
func Round(f float64) int
func RoundTo(f float64, decimals int) float64
func Ceil(f float64) int
func Floor(f float64) int

// Comparisons
func Min[T lxconstraints.Ordered](a, b T) T
func Max[T lxconstraints.Ordered](a, b T) T
func MinN[T lxconstraints.Ordered](values ...T) (T, bool)
func MaxN[T lxconstraints.Ordered](values ...T) (T, bool)

// Numeric utilities
func Divmod[T lxconstraints.Integer](a, b T) (quotient, remainder T)
func GCD[T lxconstraints.Integer](a, b T) T
func LCM[T lxconstraints.Integer](a, b T) T
func Pow[T lxconstraints.Number](base, exp T) T
func IsPowerOfTwo[T lxconstraints.Integer](n T) bool

// Percentage & ratios
func Percentage(value, total float64) float64
func PercentageChange(oldValue, newValue float64) float64
func Ratio[T lxconstraints.Number](a, b T) float64
```

**File Structure**:
- `basic.go` - Abs, Sign, Clamp, InRange
- `rounding.go` - Round, RoundTo, Ceil, Floor
- `comparison.go` - Min, Max, MinN, MaxN
- `numeric.go` - Divmod, GCD, LCM, Pow
- `percentage.go` - Percentage, PercentageChange, Ratio
- `*_test.go` - Tests

---

#### **`lxerrors`** - Error Handling Utilities
**Priority**: 🔥 HIGH  
**Rationale**: Go error handling can be verbose and repetitive

**Core Functions**:
```go
// Error checking
func Must[T any](value T, err error) T
func Must0(err error)
func MustOr[T any](value T, err error, defaultValue T) T
func IgnoreError[T any](value T, err error) T

// Error aggregation
func Join(errs ...error) error
func Combine(errs ...error) error
func First(errs ...error) error

// Error wrapping
func Wrap(err error, message string) error
func Wrapf(err error, format string, args ...any) error
func WithContext(err error, context map[string]any) error

// Error checking
func IsAny(err error, targets ...error) bool
func IsAll(err error, targets ...error) bool
func As[T error](err error) (T, bool)

// Panic recovery
func Try(fn func() error) (err error)
func TryValue[T any](fn func() (T, error)) (value T, err error)
func Recover(fn func()) error
func RecoverValue[T any](fn func() T) (T, error)
```

**File Structure**:
- `must.go` - Must*, IgnoreError
- `aggregate.go` - Join, Combine, First
- `wrap.go` - Wrap, Wrapf, WithContext
- `check.go` - IsAny, IsAll, As
- `recovery.go` - Try, Recover variants
- `*_test.go` - Tests

---
### **TIER 2: Very Useful, High Impact** (Implement Second)
---

#### **`lxjson`** - JSON Utilities
**Priority**: 🟡 MEDIUM-HIGH  
**Rationale**: JSON operations are very common and can be simplified

**Core Functions**:
```go
// Marshaling
func Marshal(v any) ([]byte, error)
func MarshalString(v any) (string, error)
func MarshalIndent(v any, indent string) ([]byte, error)
func MustMarshal(v any) []byte

// Unmarshaling
func Unmarshal[T any](data []byte) (T, error)
func UnmarshalString[T any](s string) (T, error)
func MustUnmarshal[T any](data []byte) T

// Validation
func IsValid(data []byte) bool
func IsValidString(s string) bool

// Manipulation
func Get(data []byte, path string) (any, error)
func Set(data []byte, path string, value any) ([]byte, error)
func Delete(data []byte, path string) ([]byte, error)

// Pretty printing
func Pretty(data []byte) ([]byte, error)
func PrettyString(s string) (string, error)

// Conversion
func ToMap(v any) (map[string]any, error)
func FromMap[T any](m map[string]any) (T, error)
```

---

#### **`lxhttp`** - HTTP Client Utilities
**Priority**: 🟡 MEDIUM  
**Rationale**: Building HTTP requests is repetitive

**Core Functions**:
```go
// Request builders
func Get(url string, opts ...Option) (*http.Response, error)
func Post(url string, body any, opts ...Option) (*http.Response, error)
func Put(url string, body any, opts ...Option) (*http.Response, error)
func Delete(url string, opts ...Option) (*http.Response, error)
func Patch(url string, body any, opts ...Option) (*http.Response, error)

// JSON helpers
func GetJSON[T any](url string, opts ...Option) (T, error)
func PostJSON[T any](url string, body any, opts ...Option) (T, error)

// Response helpers
func ReadBody(resp *http.Response) ([]byte, error)
func ReadBodyString(resp *http.Response) (string, error)
func UnmarshalResponse[T any](resp *http.Response) (T, error)

// Client builder
type ClientBuilder
func NewClient() *ClientBuilder
func (c *ClientBuilder) WithTimeout(d time.Duration) *ClientBuilder
func (c *ClientBuilder) WithHeader(key, value string) *ClientBuilder
func (c *ClientBuilder) Build() *http.Client

// Status checking
func IsSuccess(status int) bool
func IsClientError(status int) bool
func IsServerError(status int) bool
```

---

#### **`lxregex`** - Regex Utilities
**Priority**: 🟡 MEDIUM  
**Rationale**: Regex in Go is verbose and error-prone

**Core Functions**:
```go
// Compilation
func Compile(pattern string) (*regexp.Regexp, error)
func MustCompile(pattern string) *regexp.Regexp

// Common patterns
var (
    Email     *regexp.Regexp
    URL       *regexp.Regexp
    IPv4      *regexp.Regexp
    IPv6      *regexp.Regexp
    UUID      *regexp.Regexp
    Hex       *regexp.Regexp
    Alpha     *regexp.Regexp
    Numeric   *regexp.Regexp
    AlphaNum  *regexp.Regexp
)

// Matching
func Match(pattern, s string) bool
func MatchAll(pattern, s string) []string
func FindFirst(pattern, s string) (string, bool)
func FindAll(pattern, s string) []string

// Replacement
func Replace(pattern, s, replacement string) (string, error)
func ReplaceFunc(pattern, s string, fn func(string) string) (string, error)

// Validation
func IsEmail(s string) bool
func IsURL(s string) bool
func IsIPv4(s string) bool
func IsUUID(s string) bool
func IsHex(s string) bool
func IsAlpha(s string) bool
func IsNumeric(s string) bool
func IsAlphaNumeric(s string) bool
```

---

### **TIER 3: Specialized, Context-Specific** (Implement Based on Need)

#### **`lxcontext`** - Context Utilities
**Priority**: 🟢 MEDIUM-LOW  
**Rationale**: Context operations can be simplified

**Core Functions**:
```go
// Context creation
func WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)
func WithDeadline(parent context.Context, deadline time.Time) (context.Context, context.CancelFunc)
func WithCancel(parent context.Context) (context.Context, context.CancelFunc)

// Value management
func WithValue(parent context.Context, key, value any) context.Context
func GetValue[T any](ctx context.Context, key any) (T, bool)
func GetValueOrDefault[T any](ctx context.Context, key any, defaultValue T) T
func MustGetValue[T any](ctx context.Context, key any) T

// Checking
func IsDone(ctx context.Context) bool
func TimeRemaining(ctx context.Context) (time.Duration, bool)

// Merging
func Merge(parent context.Context, contexts ...context.Context) context.Context
```

---

#### **`lxrand`** - Random Utilities
**Priority**: 🟢 MEDIUM-LOW  
**Rationale**: Random generation needs are common

**Core Functions**:
```go
// Number generation
func Int(min, max int) int
func Int64(min, max int64) int64
func Float64(min, max float64) float64
func Bool() bool
func BoolP(probability float64) bool

// String generation
func String(length int) string
func StringAlpha(length int) string
func StringNumeric(length int) string
func StringAlphaNumeric(length int) string
func StringCustom(length int, charset string) string

// UUID & ID generation
func UUID() string
func ULID() string
func NanoID() string

// Selection
func Choice[T any](slice []T) T
func Choices[T any](slice []T, n int) []T
func WeightedChoice[T any](items []T, weights []float64) T

// Shuffle
func Shuffle[T any](slice []T)
```

---

#### **`lxvalidate`** - Validation Utilities
**Priority**: 🟢 MEDIUM-LOW  
**Rationale**: Common validation patterns

**Core Functions**:
```go
// Basic validators
func IsZero[T comparable](v T) bool
func IsNonZero[T comparable](v T) bool
func IsNil(v any) bool
func IsEmpty(v any) bool

// String validators
func IsEmail(s string) bool
func IsURL(s string) bool
func IsAlpha(s string) bool
func IsNumeric(s string) bool
func IsAlphaNumeric(s string) bool
func MinLength(s string, min int) bool
func MaxLength(s string, max int) bool

// Number validators
func InRange[T lxconstraints.Ordered](value, min, max T) bool
func IsPositive[T lxconstraints.Number](n T) bool
func IsNegative[T lxconstraints.Number](n T) bool
func IsEven[T lxconstraints.Integer](n T) bool
func IsOdd[T lxconstraints.Integer](n T) bool

// Slice validators
func AllMatch[T any](slice []T, predicate func(T) bool) bool
func AnyMatch[T any](slice []T, predicate func(T) bool) bool
func NoneMatch[T any](slice []T, predicate func(T) bool) bool

// Builder pattern
type Validator[T any]
func New[T any]() *Validator[T]
func (v *Validator[T]) Required() *Validator[T]
func (v *Validator[T]) Min(min T) *Validator[T]
func (v *Validator[T]) Max(max T) *Validator[T]
func (v *Validator[T]) Custom(fn func(T) bool) *Validator[T]
func (v *Validator[T]) Validate(value T) error
```

---

### **TIER 4: Nice to Have** (Future Consideration)
---

#### **`lxconcurrent`** - Concurrency Utilities
**Priority**: 🔵 LOW

**Core Functions**:
```go
// Goroutine management
func Go(fn func())
func GoWithContext(ctx context.Context, fn func(context.Context))
func GoGroup(fns ...func()) *sync.WaitGroup

// Parallel execution
func Parallel[T any](fns ...func() T) []T
func ParallelMap[T, U any](slice []T, fn func(T) U) []U
func ParallelFilter[T any](slice []T, predicate func(T) bool) []T

// Rate limiting
type RateLimiter
func NewRateLimiter(rate int, per time.Duration) *RateLimiter
func (r *RateLimiter) Wait()
func (r *RateLimiter) WaitContext(ctx context.Context) error

// Worker pools
type Pool
func NewPool(workers int) *Pool
func (p *Pool) Submit(fn func())
func (p *Pool) Wait()
```

---

#### **`lxtest`** - Testing Utilities
**Priority**: 🔵 LOW

**Core Functions**:
```go
// Assertions
func Equal[T comparable](t *testing.T, expected, actual T)
func NotEqual[T comparable](t *testing.T, expected, actual T)
func DeepEqual(t *testing.T, expected, actual any)
func Nil(t *testing.T, value any)
func NotNil(t *testing.T, value any)
func True(t *testing.T, value bool)
func False(t *testing.T, value bool)
func NoError(t *testing.T, err error)
func Error(t *testing.T, err error)

// Helpers
func TempFile(t *testing.T) string
func TempDir(t *testing.T) string
func Cleanup(t *testing.T, fn func())

// Mocking
type Mock
func NewMock() *Mock
func (m *Mock) On(method string, args ...any) *Mock
func (m *Mock) Return(values ...any) *Mock
func (m *Mock) Called(args ...any) []any
```

---

## 📊 Implementation Priority Matrix

| Package | Priority | Complexity | Value | When to Implement |
|---------|----------|------------|-------|-------------------|
| `lxmaps` | 🔥 Highest | Low | Very High | **IMMEDIATELY** |
| `lxmath` | 🔥 High | Low | High | **Phase 1** |
| `lxerrors` | 🔥 High | Medium | High | **Phase 1** |
| `lxio` | 🔥 High | Medium | Very High | **Phase 1** |
| `lxtime` | 🟡 Med-High | Medium | High | **Phase 2** |
| `lxjson` | 🟡 Med-High | Low | High | **Phase 2** |
| `lxhttp` | 🟡 Medium | Medium | Medium | **Phase 2** |
| `lxregex` | 🟡 Medium | Low | Medium | **Phase 2** |
| `lxcontext` | 🟢 Med-Low | Low | Medium | **Phase 3** |
| `lxrand` | 🟢 Med-Low | Low | Medium | **Phase 3** |
| `lxvalidate` | 🟢 Med-Low | Medium | Medium | **Phase 3** |
| `lxcrypto` | 🟢 Low-Med | High | Medium | **Phase 3** |
| `lxenv` | 🔵 Low | Low | Low | **Phase 4** |
| `lxconcurrent` | 🔵 Low | High | Medium | **Phase 4** |
| `lxtest` | 🔵 Low | Medium | Low | **Phase 4** |

---

## 🎯 Recommended Implementation Phases

### **Phase 1: Core Essentials** (Next 1-2 months)
Focus on packages that provide maximum value with minimal dependencies:
1. ✅ Complete `lxslices` improvements (from ROADMAP.md)
2. 🆕 Implement `lxmaps` 
3. 🆕 Implement `lxmath`
4. 🆕 Implement `lxerrors`
5. 🆕 Implement `lxio`

### **Phase 2: High-Value Utilities** (Months 3-4)
Add packages that handle common patterns:
1. 🆕 Implement `lxtime`
2. 🆕 Implement `lxjson`
3. 🆕 Implement `lxhttp`
4. 🆕 Implement `lxregex`
5. ⬆️ Enhance `lxptrs` with more utilities
6. ⬆️ Enhance `lxtuples` (Triple, Quad, etc.)

### **Phase 3: Specialized Tools** (Months 5-6)
Based on community feedback and usage patterns:
1. 🆕 Implement highest-voted packages from Tier 3
2. ⬆️ Enhance existing packages based on user feedback
3. 📚 Comprehensive documentation and examples

### **Phase 4: Advanced Features** (Later)
- Implement remaining packages based on demand
- Performance optimizations
- Benchmarking suite
- Migration guides

---

## 🔧 Enhancements to Existing Packages

### **`lxptrs`** - Needs Expansion (Current: ~60%)
```go
// Add these functions:
func RefOr[V any](v *V, defaultValue V) V
func DerefOr[V any](v *V, defaultValue V) V
func IsNil[V any](v *V) bool
func IsNotNil[V any](v *V) bool
func Equal[V comparable](a, b *V) bool
func Compare[V comparable](a, b *V) int
```

### **`lxtuples`** - Needs Expansion (Current: ~40%)
```go
// Add these types:
type Triple[T, U, V any] struct {
    First  T
    Second U
    Third  V
}

type Quad[T, U, V, W any] struct {
    First  T
    Second U
    Third  V
    Fourth W
}

// Add factory functions:
func NewPair[T, U any](first T, second U) Pair[T, U]
func NewTriple[T, U, V any](first T, second U, third V) Triple[T, U, V]

// Add utility methods:
func (p Pair[T, U]) Swap() Pair[U, T]
func (p Pair[T, U]) Values() (T, U)
```

### **`lxsystems`** - Needs Expansion (Current: ~70%)
```go
// Add these functions:
func EnvVars() map[string]string
func EnvVar(key string) (string, bool)
func SetEnvVar(key, value string) error
func ProcessID() int
func ParentProcessID() int
func Executable() (string, error)
func IsRoot() bool // Unix systems
func IsAdmin() bool // Windows systems
```

### **`lxconstraints`** - Consider Adding
```go
// Signed integers only
type Signed interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned integers only  
type Unsigned interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Complex numbers
type Complex interface {
    ~complex64 | ~complex128
}
```

---

## 📝 Package Design Principles

Based on your current excellent work, maintain these principles:

### ✅ DO:
- Keep functions small, focused, and composable
- Provide both `Must*` and error-returning variants where appropriate
- Use `*Or` variants for default values (e.g., `GetOr`, `FirstOr`)
- Return `(value, bool)` for operations that might fail on valid input
- Provide comprehensive test coverage
- Use descriptive function names
- Document edge cases and nil handling
- Follow Go idioms and conventions

### ❌ DON'T:
- Create giant "utils" packages
- Add dependencies unless absolutely necessary
- Include framework-specific code
- Add functions that are trivial wrappers with no value
- Sacrifice readability for brevity
- Ignore error handling
- Break backward compatibility

---

## 🚀 Quick Start Guide for New Packages

### 1. **Package Structure Template**

**For small packages:**
```
lxpackagename/
├── packagename.go        # Main implementation
├── packagename_test.go   # Tests
├── doc.go               # Package documentation (optional)
├── README.md            # Package-specific docs (optional)
└── examples_test.go     # Example functions (optional)
```

**For large packages** (like lxslices):
```
lxpackagename/
├── feature1.go          # Feature group 1 (e.g., aggregation.go)
├── feature1_test.go     # Tests for feature 1
├── feature2.go          # Feature group 2 (e.g., filter.go)
├── feature2_test.go     # Tests for feature 2
├── errs.go              # Package-specific errors (if needed)
├── doc.go               # Package documentation (optional)
├── README.md            # Package-specific docs (optional)
├── ROADMAP.md           # Future enhancements (optional)
└── examples_test.go     # Example functions (optional)
```

### 2. **File Header Template**
```go
// Package lxpackagename provides [brief description].
//
// This package is part of the lx library and follows Go idioms
// for [specific domain].
package lxpackagename
```

### 3. **Test File Template**
```go
package lxpackagename

import "testing"

func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    Type
        expected Type
    }{
        // Test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

---

## 🎓 Learning from Popular Libraries

Your packages compare well with:
- **golang.org/x/exp** - experimental packages
- **samber/lo** - lodash-style utilities
- **spf13/cast** - type conversion
- **jinzhu/copier** - struct copying

**Your advantages**:
- Smaller, more focused scope
- Better nil handling
- Consistent API patterns
- Excellent documentation
- Native Go generics from the start

---

## 📈 Success Metrics

Track these to validate new packages:

1. **Usage**: GitHub stars, downloads, imports
2. **Quality**: Test coverage (>80%), no critical bugs
3. **Adoption**: Community feedback, issues, PRs
4. **Performance**: Benchmarks vs alternatives
5. **Documentation**: Examples, godoc, README completeness

---

## 🤝 Community Input

Before implementing new packages:
1. Create GitHub issues for discussion
2. Gather community feedback
3. Validate use cases
4. Consider alternatives
5. Start with minimal viable API

---

## 📅 Timeline Estimate

| Phase | Duration | Packages | Effort |
|-------|----------|----------|--------|
| Phase 1 | 1-2 months | 4 packages | ~40-60 hours |
| Phase 2 | 2-3 months | 4 packages | ~40-60 hours |
| Phase 3 | 2-3 months | 4 packages | ~40-60 hours |
| Phase 4 | Ongoing | 3 packages | ~30-40 hours |

**Total estimated**: 6-9 months for complete roadmap

---

## 🎉 Summary

**Immediate Next Steps**:
1. ✅ Complete `lxslices` enhancements (already planned in ROADMAP.md)
2. 🆕 Start with `lxmaps` - highest value, lowest complexity
3. 🆕 Follow with `lxmath` and `lxerrors`
4. 🆕 Then implement `lxio`

**Key Focus Areas**:
- Maps and collections (`lxmaps`)
- Mathematical operations (`lxmath`)
- Error handling (`lxerrors`)
- File I/O (`lxio`)

This roadmap provides a strategic path to building a comprehensive, high-quality Go utilities library while maintaining your excellent standards for code quality, testing, and documentation.

**Keep up the amazing work!** 🚀

