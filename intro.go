package main

/**
1. type parameter for functions and types
*/

type ConfigX struct {
	X string `json:"x"`
}

type ConfigY struct {
	Y bool `json:"y"`
}

type ConfigType interface {
	ConfigX | ConfigY
}

type Config[T ConfigType] struct {
	Name  string `json:"name"`
	Value T      `json:"value"`
}
