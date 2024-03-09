package main

type Error struct {
	Message string `json:"message"`
}

type File struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Type     string `json:"type"`
	Children []File `json:"children"`
}
