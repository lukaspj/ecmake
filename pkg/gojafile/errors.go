package gojafile

import "fmt"

type Error struct {
	Message string
}

func (g Error) Error() string {
	return g.Message
}

type TargetNotFoundError struct {
	Target string
}

func (t TargetNotFoundError) Error() string {
	return fmt.Sprintf("target %s was not found", t.Target)
}

type TargetInvalidError struct {
	Target string
}

func (t TargetInvalidError) Error() string {
	return fmt.Sprintf("target %s was not a valid function", t.Target)
}

type TargetExecutionError struct {
	Target string
	Args   []string
	Cause  error
}

func (t TargetExecutionError) Error() string {
	return fmt.Sprintf("target %s failed to execute with args: %v, caused by: %v", t.Target, t.Args, t.Cause)
}
