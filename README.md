# Refresh Repositories

## Introduction
This library helps you truncate your RDBM tables, and also delete your elasticsearch indices.

## Installation
Using this library is easy. First, use go get to install the latest version of the library.
```bash
go get -u github.com/majidalaeinia/refresh-repositories@latest
```
Next, include it in your application:
```bash
import rr "github.com/majidalaeinia/refresh-repositories"
```

## Usage
You can use this library at the end of your test functions (make sure you are migrating your repositories at the start 
of your test functions).  

**Important Note:** Make sure you are using this library in development mode with local repositories, since it truncates 
your mysql tables and deletes your elasticsearch indices.

#### MySql

```go
package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	
	rr "github.com/majidalaeinia/refresh-repositories"
)

func TestClientBasketCtrl_AddItem_Status_200(t *testing.T) {
	tests.Init() //On this function you can migrate your repositories
	asr := assert.New(t)
	db := mysql.Get()

	//...
	
	// your test goes here
	
	//...
	
	asr.Equal(http.StatusOK, w.Code)
	tables := []string{"first_table", "second_table"}
	err = rr.TruncateRdbm(db, tables)
	if err != nil {
		asr.Error(err)
	}
}
```

#### Elasticsearch

```go
package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	
	rr "github.com/majidalaeinia/refresh-repositories"
)

func TestClientBasketCtrl_AddItem_Status_200(t *testing.T) {
	tests.Init() //On this function you can migrate your repositories
	asr := assert.New(t)
	el := elastic.Get()

	//...
	
	// your test goes here
	
	//...
	
	asr.Equal(http.StatusOK, w.Code)
	indices := []string{"first_index", "second_index"}
	err = rr.TruncateNoSql(el, indices)
	if err != nil {
		asr.Error(err)
	}
}
```

#### CQRS

```go
package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	
	rr "github.com/majidalaeinia/refresh-repositories"
)

func TestClientBasketCtrl_AddItem_Status_200(t *testing.T) {
	tests.Init() //On this function you can migrate your repositories
	asr := assert.New(t)
	db := mysql.Get()
	el := elastic.Get()

	//...
	
	// your test goes here
	
	//...
	
	asr.Equal(http.StatusOK, w.Code)
	tables := []string{"first_table", "second_table"}
	indices := []string{"first_index", "second_index"}
	err = rr.TruncateRepositories(db, el, tables, indices)
	if err != nil {
		asr.Error(err)
	}
}
```

### TODO
- [ ] Add tests.
- [ ] Add other repositories than mysql and elasticsearch.

