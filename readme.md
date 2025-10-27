# CACA

Create 

A 

Cool 

App

*flags*

-name="<project_name>"

-template-path="<path_to_project_template>"

-template-name="<template_name_in_config>"

# working with it in a less loser way

you can make a config file (~/.config/caca/caca.json)

which allows you to have multiple templates for your projects, and with the **flag** template-name you can choose it from it


```json
{
	"default_template": "ebit",
	"templates": [
		{
			"template_name": "ebit",
			"path": "/home/marcig/Documents/projects/ebit",
			"files_to_replace": [
				"Makefile",
				"go.mod"
			]
		},
		{
			"template_name": "guml",
			"path": "/home/marcig/Documents/projects/guml"
		},
		{
			"template_name": "jshot",
			"path": "/home/marcig/Documents/projects/jshot"
		}
	]
}
```

roadmap is to add more useful regex replace in certain files, and also make a cute owo tui with bubbletea but rn im eppy :3
