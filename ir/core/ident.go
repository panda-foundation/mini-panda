package core

import (
	"fmt"
	"strconv"
)

// Ident is a named variable.
type Ident interface {
	Name() string
	SetName(name string)
	ID() int64
	SetID(id int64)
}

// GlobalIdent is a global identifier.
type GlobalIdent struct {
	GlobalName string
	GlobalID   int64
}

// Ident returns the identifier associated with the global identifier.
func (g *GlobalIdent) Ident() string {
	if g.GlobalName == "" {
		return globalID(g.GlobalID)
	}
	return globalName(g.GlobalName)
}

// Name returns the name of the global identifier.
//
// If unnamed, the global ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (g *GlobalIdent) Name() string {
	if g.GlobalName == "" {
		return strconv.FormatInt(g.GlobalID, 10)
	}
	if x, err := strconv.ParseInt(g.GlobalName, 10, 64); err == nil {
		// Print GlobalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return g.GlobalName
}

// SetName sets the name of the global identifier.
func (g *GlobalIdent) SetName(name string) {
	g.GlobalName = name
	g.GlobalID = 0
}

// ID returns the ID of the global identifier.
func (g *GlobalIdent) ID() int64 {
	return g.GlobalID
}

// SetID sets the ID of the global identifier.
func (g *GlobalIdent) SetID(id int64) {
	g.GlobalID = id
}

// LocalIdent is a local identifier.
type LocalIdent struct {
	LocalName string
	LocalID   int64
}

// NewLocalIdent returns a new local identifier based on the given string. An
// unnamed local ID is used if ident is an integer, and a local name otherwise.
func NewLocalIdent(ident string) LocalIdent {
	if id, err := strconv.ParseInt(ident, 10, 64); err == nil {
		return LocalIdent{LocalID: id}
	}
	return LocalIdent{LocalName: ident}
}

// Ident returns the identifier associated with the local identifier.
func (l *LocalIdent) Ident() string {
	if l.LocalName == "" {
		return localID(l.LocalID)
	}
	return localName(l.LocalName)
}

// Name returns the name of the local identifier.
//
// If unnamed, the local ID is returned. To distinguish numeric names from
// unnamed IDs, numeric names are quoted.
func (l *LocalIdent) Name() string {
	if l.LocalName == "" {
		return strconv.FormatInt(l.LocalID, 10)
	}
	if x, err := strconv.ParseInt(l.LocalName, 10, 64); err == nil {
		// Print LocalName with quotes if it is a number; e.g. "42".
		return fmt.Sprintf(`"%d"`, x)
	}
	return l.LocalName
}

// SetName sets the name of the local identifier.
func (l *LocalIdent) SetName(name string) {
	l.LocalName = name
	l.LocalID = 0
}

// ID returns the ID of the local identifier.
func (l *LocalIdent) ID() int64 {
	return l.LocalID
}

// SetID sets the ID of the local identifier.
func (i *LocalIdent) SetID(id int64) {
	i.LocalID = id
}
