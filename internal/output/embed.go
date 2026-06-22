package output

import "embed"

//go:embed logo/*
var LogoFS embed.FS

// This file makes it so the ascii art text files are included in the binary
