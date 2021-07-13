---
sidebar_position: 1
---

# Add Plugins

This page helps you to add your own plugin

## Add a new language

- Need to add new language plugin folder 
- In `code-analyser/static/supportedLanguages.yaml` 
```
languages:
  - name: "Go"
    path: "./plugin/go"
```
- If folder already exists can continue to the next step

## Add a new plugin
- I you want your folder recognized as a plugin you need to have `pluginDetails.yaml`
- Structure of `pluginDetails.yaml`
  ```
version: "0.1"
author:
  name: "author name"
  email: "author email"
date: "dd-mm-yyyy"
plugindetails:
  type: "pluginType"
  name: "name of plugin"
  version: "v1.x"
  description: "plugin description"
  icon: ""
  semver:
    - ">=0.x,<1.x"
    - ">=2.3,<4.x"
  command: "node alfredPlugin/server.js" (plugin run command)
  ```

###  Plugin Type
- Language Specific
  - detectRuntime
  - commands
  - env
  - preDetectCommand
  - staticAssets
  - buildDirectory
  - testCasesCommands
  - framework
  - orm
  - library
  - database
  - getDependencies
- Global
  - dockerFile
  - procFile
  - makeFile

- Every plugin show follow their respective proto for rpc calls. 
- Server shoould start by chacking available ports.
- Plugin directory must be in it respective language folder. 
