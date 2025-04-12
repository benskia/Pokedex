package customErr

import "errors"

var ErrMissingCommand = errors.New("missing command")
var ErrMissingLocation = errors.New("missing location name")
var ErrMissingPokemon = errors.New("missing pokemon name")

var ErrClientError = errors.New("invalid location name")
var ErrServerError = errors.New("PokeAPI server error")
