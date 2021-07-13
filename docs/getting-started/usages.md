---
sidebar_position: 3
---

# Usage

```go
import projectName/code-analyser

----
----
//Initialize logger
logger, _ := zap.NewProduction()
// get absolute path of code-base
var path = "/my/codebase/absolute/path"
//set configuration 
config := utils.Setting{
		Logger: logger,
	}
// instanciate Analyser & add config  
analyser := Analyser{Setting: &config}
// Start Scrape
output,error:=analyser.Analyse(ctx, path
log.Println(output,error)

```

