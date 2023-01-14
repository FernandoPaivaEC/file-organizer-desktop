package main

type LastModified struct {
	Day   string
	Month string
	Year  string
}

type FileInfo struct {
	Name         string
	Keyword      string
	LastModified LastModified
}

type FileIndex []FileInfo
