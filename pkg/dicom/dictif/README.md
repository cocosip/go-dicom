# Package dictif

## Overview

The `dictif` package defines interfaces for DICOM dictionary lookups. This package exists to solve the circular dependency problem between the `tag` and `dict` packages.

## Problem

Previously, there was a circular dependency issue:
- The `tag` package needs to look up dictionary entries for tags
- The `dict` package needs to work with `tag.Tag` types

The old solution used function callbacks registered via `SetDictionaryLookup`, `SetKeywordLookup`, and `SetPrivateCreatorLookup` functions in the `tag` package. While functional, this approach was:
- Not idiomatic Go
- Hard to test
- Difficult to understand the dependency flow
- Used global mutable state

## Solution

The new design uses interface-based dependency inversion:

1. **dictif package** (this package) - Defines minimal interfaces:
   - `Tag` - Basic tag interface (Group, Element, ToUint32)
   - `Entry` - Dictionary entry interface (Name, Keyword, VRs, VM)
   - `PrivateCreator` - Private creator interface (Creator)
   - `Lookup` - Main lookup interface with three methods

2. **tag package** - Depends only on `dictif`:
   - Uses `dictif.GlobalLookup()` to get the dictionary implementation
   - No direct dependency on `dict` package

3. **dict package** - Implements `dictif.Lookup`:
   - `Dictionary` implements the lookup interface via an adapter
   - Registers itself as the global lookup in `init()`

## Architecture Diagram

```
┌──────────┐
│ dictif   │  ← Defines interfaces only
└────┬─────┘
     │
     ├─────────────┐
     │             │
     ↓             ↓
┌────────┐    ┌────────┐
│  tag   │    │  dict  │
│        │    │        │
│ Uses   │    │ Impls  │
│dictif  │    │dictif  │
└────────┘    └───┬────┘
                  │
                  ↓
            Registers as
            GlobalLookup
```

## Benefits

1. **No Circular Dependencies** - Clean one-way dependency flow
2. **Testability** - Easy to mock the lookup interface
3. **Type Safety** - Interfaces provide compile-time type checking
4. **Idiomatic Go** - Uses standard interface-based design
5. **Clarity** - Explicit about what `tag` needs from `dict`

## Usage Example

### In tag package
```go
import "github.com/cocosip/go-dicom/pkg/dicom/dictif"

func (t *Tag) DictionaryEntry() interface{} {
    lookup := dictif.GlobalLookup()
    if lookup == nil {
        return nil
    }
    return lookup.LookupTag(t)
}
```

### In dict package
```go
import "github.com/cocosip/go-dicom/pkg/dicom/dictif"

func init() {
    dictif.SetGlobalLookup(&dictionaryAdapter{Dictionary: Default()})
}
```

## Implementation Notes

### Adapter Pattern

The `dict` package uses an adapter to implement the `dictif.Lookup` interface:

```go
type dictionaryAdapter struct {
    *Dictionary
}

func (da *dictionaryAdapter) LookupKeyword(keyword string) dictif.Tag {
    tag := da.Dictionary.LookupKeyword(keyword)
    return tag // *tag.Tag implements dictif.Tag
}
```

This adapter is needed because:
- Go's interface matching requires exact method signatures
- We can't change existing `Dictionary` methods without breaking compatibility
- The adapter provides the interface-required return types while delegating to existing methods

### Type Compatibility

The design leverages Go's structural typing:
- `*tag.Tag` implements `dictif.Tag` (has Group(), Element(), ToUint32())
- `*dict.Entry` implements `dictif.Entry` (has Name(), Keyword(), VRs(), VM())
- `*tag.PrivateCreator` implements `dictif.PrivateCreator` (has Creator())

No explicit type assertions or conversions are needed; the compiler verifies compatibility.

## Migration from Old Approach

### Before (Function Callbacks)
```go
// In dict package init()
tag.SetDictionaryLookup(func(t *tag.Tag) interface{} {
    return Default().Lookup(t)
})
tag.SetKeywordLookup(func(keyword string) (*tag.Tag, error) {
    t := Default().LookupKeyword(keyword)
    if t == nil {
        return nil, nil
    }
    return t, nil
})
```

### After (Interface-Based)
```go
// In dict package init()
dictif.SetGlobalLookup(&dictionaryAdapter{Dictionary: Default()})
```

The new approach is simpler, more testable, and follows Go best practices.
