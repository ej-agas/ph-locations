package models

type Model interface {
	ID() int
	String() string
}
