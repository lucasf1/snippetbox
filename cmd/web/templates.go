package main

import "github.com/lucasf1/snippetbox/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
