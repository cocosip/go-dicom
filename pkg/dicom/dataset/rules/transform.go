// Copyright (c) 2025 go-dicom contributors.
// Licensed under the Microsoft Public License (MS-PL).

package rules

import (
	"errors"
	"fmt"

	"github.com/cocosip/go-dicom/pkg/dicom/dataset"
)

// TransformStage identifies the phase that rejected a transform.
type TransformStage string

const (
	// StageCondition identifies condition evaluation.
	StageCondition TransformStage = "condition"
	// StageRule identifies transform rule execution.
	StageRule TransformStage = "rule"
	// StageCommit identifies replacement of the target Dataset.
	StageCommit TransformStage = "commit"
	// StageRollback identifies restoration after a failed commit.
	StageRollback TransformStage = "rollback"
)

// TransformError adds stage, nested rule indexes, and Dataset path context.
type TransformError struct {
	Stage    TransformStage
	RulePath []int
	Path     dataset.Path
	Cause    error
}

func (e *TransformError) Error() string {
	if e == nil {
		return "<nil>"
	}
	location := dataset.FormatPath(e.Path)
	if len(e.RulePath) > 0 {
		location = fmt.Sprintf("rules %v at %s", e.RulePath, location)
	}
	if e.Cause == nil {
		return fmt.Sprintf("DICOM transform %s failed at %s", e.Stage, location)
	}
	return fmt.Sprintf("DICOM transform %s failed at %s: %v", e.Stage, location, e.Cause)
}

// Unwrap returns the match, rule, conversion, validation, or commit cause.
func (e *TransformError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

// TransformRule is an immutable operation executed only by a Transformer.
// The unexported method prevents bypassing transaction and ChangeSet handling.
type TransformRule interface {
	apply(*dataset.Dataset, dataset.Path, *ChangeSet) error
}

// Transformer executes an optional condition and ordered rules.
type Transformer struct {
	condition MatchRule
	rules     []TransformRule
}

// NewTransformer constructs an unconditional ordered Transformer.
func NewTransformer(rules ...TransformRule) (*Transformer, error) {
	return newTransformer(nil, rules)
}

// NewConditionalTransformer constructs a Transformer whose condition is
// evaluated once against the supplied Dataset.
func NewConditionalTransformer(condition MatchRule, rules ...TransformRule) (*Transformer, error) {
	if isNilRule(condition) {
		return nil, fmt.Errorf("transform condition is nil")
	}
	return newTransformer(condition, rules)
}

func newTransformer(condition MatchRule, rules []TransformRule) (*Transformer, error) {
	copied := append([]TransformRule(nil), rules...)
	for index, rule := range copied {
		if isNilRule(rule) {
			return nil, fmt.Errorf("transform rule %d is nil", index)
		}
	}
	return &Transformer{condition: condition, rules: copied}, nil
}

// Apply transforms an independent deep clone. The source is never mutated.
func (transformer *Transformer) Apply(source *dataset.Dataset) (*dataset.Dataset, ChangeSet, error) {
	if transformer == nil {
		return nil, nil, &TransformError{Stage: StageRule, Cause: fmt.Errorf("Transformer is nil")}
	}
	if source == nil {
		return nil, nil, &TransformError{Stage: StageRule, Cause: fmt.Errorf("transform source Dataset is nil")}
	}

	result := source.DeepClone()
	changes := make(ChangeSet, 0)
	if err := transformer.apply(result, nil, &changes); err != nil {
		return nil, changes, err
	}
	return result, changes, nil
}

// ApplyInPlace commits a complete successful clone while preserving target's
// pointer and automatic-validation setting.
func (transformer *Transformer) ApplyInPlace(target *dataset.Dataset) (ChangeSet, error) {
	if target == nil {
		return nil, &TransformError{Stage: StageCommit, Cause: fmt.Errorf("transform target Dataset is nil")}
	}
	result, changes, err := transformer.Apply(target)
	if err != nil {
		return changes, err
	}
	if err := target.ReplaceFrom(result); err != nil {
		return changes, &TransformError{Stage: StageCommit, Cause: err}
	}
	return changes, nil
}

func (transformer *Transformer) apply(target *dataset.Dataset, path dataset.Path, changes *ChangeSet) error {
	if transformer == nil {
		return &TransformError{Stage: StageRule, Path: dataset.ClonePath(path), Cause: fmt.Errorf("Transformer is nil")}
	}
	if transformer.condition != nil {
		matched, err := transformer.condition.Match(target)
		if err != nil {
			return &TransformError{Stage: StageCondition, Path: dataset.ClonePath(path), Cause: err}
		}
		if !matched {
			return nil
		}
	}

	for index, rule := range transformer.rules {
		if err := rule.apply(target, path, changes); err != nil {
			return prependTransformRule(index, path, err)
		}
	}
	return nil
}

func prependTransformRule(index int, path dataset.Path, err error) error {
	var transformErr *TransformError
	if errors.As(err, &transformErr) {
		cloned := &TransformError{
			Stage:    transformErr.Stage,
			RulePath: append([]int{index}, transformErr.RulePath...),
			Path:     dataset.ClonePath(transformErr.Path),
			Cause:    transformErr.Cause,
		}
		if len(cloned.Path) == 0 {
			cloned.Path = dataset.ClonePath(path)
		}
		return cloned
	}
	return &TransformError{
		Stage:    StageRule,
		RulePath: []int{index},
		Path:     dataset.ClonePath(path),
		Cause:    err,
	}
}
