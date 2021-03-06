// generaTed by fileb0x at "2017-03-20 15:50:21.560240036 -0400 EDT" from config file "b0x.yaml"

package templates


import (
  "bytes"
  
  "io"
  "log"
  "net/http"
  "os"

  "golang.org/x/net/webdav"
  "golang.org/x/net/context"
)

var ( 
  // CTX is a context for webdav vfs
  CTX = context.Background()

  
  // FS is a virtual memory file system
  FS = webdav.NewMemFS()
  

  // Handler is used to server files through a http handler
  Handler *webdav.Handler

  // HTTP is the http file system
  HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {}



// FileTemplatesMysqlCreateDatabaseSQLTmpl is "templates/mysql/create_database.sql.tmpl"
var FileTemplatesMysqlCreateDatabaseSQLTmpl = []byte("\x2d\x2d\x0a\x2d\x2d\x20\x43\x75\x72\x72\x65\x6e\x74\x20\x44\x61\x74\x61\x62\x61\x73\x65\x3a\x20\x60\x7b\x7b\x20\x2e\x44\x61\x74\x61\x62\x61\x73\x65\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x2f\x2a\x21\x34\x30\x30\x30\x30\x20\x44\x52\x4f\x50\x20\x44\x41\x54\x41\x42\x41\x53\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x44\x61\x74\x61\x62\x61\x73\x65\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x2a\x2f\x3b\x0a\x7b\x7b\x20\x2e\x44\x61\x74\x61\x62\x61\x73\x65\x2e\x43\x72\x65\x61\x74\x65\x53\x51\x4c\x20\x7d\x7d\x3b\x0a\x0a\x55\x53\x45\x20\x60\x7b\x7b\x20\x2e\x44\x61\x74\x61\x62\x61\x73\x65\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x3b")

// FileTemplatesMysqlCreateRoutinesSQLTmpl is "templates/mysql/create_routines.sql.tmpl"
var FileTemplatesMysqlCreateRoutinesSQLTmpl = []byte("\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x52\x6f\x75\x74\x69\x6e\x65\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x52\x6f\x75\x74\x69\x6e\x65\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x44\x52\x4f\x50\x20\x7b\x7b\x20\x2e\x54\x79\x70\x65\x20\x7d\x7d\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x3d\x20\x40\x40\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x27\x7b\x7b\x20\x2e\x53\x51\x4c\x4d\x6f\x64\x65\x20\x7d\x7d\x27\x20\x2a\x2f\x20\x3b\x0a\x44\x45\x4c\x49\x4d\x49\x54\x45\x52\x20\x3b\x3b\x0a\x43\x52\x45\x41\x54\x45\x20\x44\x45\x46\x49\x4e\x45\x52\x3d\x7b\x7b\x20\x2e\x44\x65\x66\x69\x6e\x65\x72\x20\x7d\x7d\x20\x7b\x7b\x20\x2e\x54\x79\x70\x65\x20\x7d\x7d\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x28\x7b\x7b\x20\x2e\x50\x61\x72\x61\x6d\x4c\x69\x73\x74\x20\x7d\x7d\x29\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x54\x79\x70\x65\x20\x22\x46\x55\x4e\x43\x54\x49\x4f\x4e\x22\x20\x7d\x7d\x20\x52\x45\x54\x55\x52\x4e\x53\x20\x7b\x7b\x20\x2e\x52\x65\x74\x75\x72\x6e\x73\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x49\x73\x44\x65\x74\x65\x72\x6d\x69\x6e\x69\x73\x74\x69\x63\x20\x22\x59\x45\x53\x22\x20\x7d\x7d\x0a\x20\x20\x20\x20\x44\x45\x54\x45\x52\x4d\x49\x4e\x49\x53\x54\x49\x43\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x2e\x43\x72\x65\x61\x74\x65\x53\x51\x4c\x20\x7d\x7d\x20\x3b\x3b\x0a\x44\x45\x4c\x49\x4d\x49\x54\x45\x52\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x20\x3b\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d")

// FileTemplatesMysqlCreateTablesSQLTmpl is "templates/mysql/create_tables.sql.tmpl"
var FileTemplatesMysqlCreateTablesSQLTmpl = []byte("\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x54\x61\x62\x6c\x65\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x54\x61\x62\x6c\x65\x20\x73\x74\x72\x75\x63\x74\x75\x72\x65\x20\x66\x6f\x72\x20\x74\x61\x62\x6c\x65\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x44\x65\x62\x75\x67\x4d\x73\x67\x73\x20\x7d\x7d\x2d\x2d\x20\x44\x65\x62\x75\x67\x3a\x20\x7b\x7b\x20\x2e\x20\x7d\x7d\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x6e\x6f\x74\x20\x24\x2e\x41\x72\x67\x73\x2e\x53\x6b\x69\x70\x41\x64\x64\x44\x72\x6f\x70\x54\x61\x62\x6c\x65\x20\x7d\x7d\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x3b\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x7b\x7b\x20\x2e\x43\x72\x65\x61\x74\x65\x53\x51\x4c\x20\x7d\x7d\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x3b\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x52\x6f\x77\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x44\x75\x6d\x70\x69\x6e\x67\x20\x64\x61\x74\x61\x20\x66\x6f\x72\x20\x74\x61\x62\x6c\x65\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x4c\x4f\x43\x4b\x20\x54\x41\x42\x4c\x45\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x57\x52\x49\x54\x45\x3b\x0a\x2f\x2a\x21\x34\x30\x30\x30\x30\x20\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x44\x49\x53\x41\x42\x4c\x45\x20\x4b\x45\x59\x53\x20\x2a\x2f\x3b\x0a\x7b\x7b\x20\x2e\x7c\x54\x61\x62\x6c\x65\x49\x6e\x73\x65\x72\x74\x73\x20\x7d\x7d\x0a\x2f\x2a\x21\x34\x30\x30\x30\x30\x20\x41\x4c\x54\x45\x52\x20\x54\x41\x42\x4c\x45\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x45\x4e\x41\x42\x4c\x45\x20\x4b\x45\x59\x53\x20\x2a\x2f\x3b\x0a\x55\x4e\x4c\x4f\x43\x4b\x20\x54\x41\x42\x4c\x45\x53\x3b\x0a\x0a\x7b\x7b\x20\x69\x66\x20\x24\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x54\x72\x69\x67\x67\x65\x72\x73\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x74\x72\x69\x67\x67\x65\x72\x73\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d")

// FileTemplatesMysqlCreateTriggersSQLTmpl is "templates/mysql/create_triggers.sql.tmpl"
var FileTemplatesMysqlCreateTriggersSQLTmpl = []byte("\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x54\x72\x69\x67\x67\x65\x72\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x54\x72\x69\x67\x67\x65\x72\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x3d\x20\x40\x40\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x27\x7b\x7b\x20\x2e\x53\x51\x4c\x4d\x6f\x64\x65\x20\x7d\x7d\x27\x20\x2a\x2f\x20\x3b\x0a\x44\x45\x4c\x49\x4d\x49\x54\x45\x52\x20\x3b\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x43\x52\x45\x41\x54\x45\x2a\x2f\x20\x2f\x2a\x21\x35\x30\x30\x31\x37\x20\x44\x45\x46\x49\x4e\x45\x52\x3d\x7b\x7b\x20\x2e\x44\x65\x66\x69\x6e\x65\x72\x20\x7d\x7d\x2a\x2f\x20\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x54\x52\x49\x47\x47\x45\x52\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x7b\x7b\x20\x2e\x41\x63\x74\x69\x6f\x6e\x54\x69\x6d\x69\x6e\x67\x20\x7d\x7d\x20\x7b\x7b\x20\x2e\x45\x76\x65\x6e\x74\x4d\x61\x6e\x69\x70\x75\x6c\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x20\x4f\x4e\x20\x60\x7b\x7b\x20\x2e\x45\x76\x65\x6e\x74\x4f\x62\x6a\x65\x63\x74\x54\x61\x62\x6c\x65\x20\x7d\x7d\x60\x0a\x46\x4f\x52\x20\x45\x41\x43\x48\x20\x7b\x7b\x20\x2e\x41\x63\x74\x69\x6f\x6e\x4f\x72\x69\x65\x6e\x74\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x0a\x7b\x7b\x20\x2e\x43\x72\x65\x61\x74\x65\x53\x51\x4c\x20\x7d\x7d\x20\x2a\x2f\x3b\x3b\x0a\x44\x45\x4c\x49\x4d\x49\x54\x45\x52\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x73\x71\x6c\x5f\x6d\x6f\x64\x65\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x20\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x33\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x20\x3b\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d")

// FileTemplatesMysqlCreateViewsFinalSQLTmpl is "templates/mysql/create_views_final.sql.tmpl"
var FileTemplatesMysqlCreateViewsFinalSQLTmpl = []byte("\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x56\x69\x65\x77\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x46\x69\x6e\x61\x6c\x20\x76\x69\x65\x77\x20\x73\x74\x72\x75\x63\x74\x75\x72\x65\x20\x66\x6f\x72\x20\x76\x69\x65\x77\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x7b\x7b\x20\x69\x66\x20\x6e\x6f\x74\x20\x24\x2e\x41\x72\x67\x73\x2e\x53\x6b\x69\x70\x41\x64\x64\x44\x72\x6f\x70\x54\x61\x62\x6c\x65\x20\x7d\x7d\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x44\x52\x4f\x50\x20\x56\x49\x45\x57\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x2a\x2f\x3b\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x20\x20\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x20\x20\x20\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x43\x52\x45\x41\x54\x45\x20\x41\x4c\x47\x4f\x52\x49\x54\x48\x4d\x3d\x55\x4e\x44\x45\x46\x49\x4e\x45\x44\x20\x2a\x2f\x0a\x2f\x2a\x21\x35\x30\x30\x31\x33\x20\x44\x45\x46\x49\x4e\x45\x52\x3d\x7b\x7b\x20\x2e\x44\x65\x66\x69\x6e\x65\x72\x20\x7d\x7d\x20\x53\x51\x4c\x20\x53\x45\x43\x55\x52\x49\x54\x59\x20\x7b\x7b\x20\x2e\x53\x65\x63\x75\x72\x69\x74\x79\x54\x79\x70\x65\x20\x7d\x7d\x20\x2a\x2f\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x7b\x7b\x20\x2e\x43\x72\x65\x61\x74\x65\x53\x51\x4c\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x20\x20\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x72\x65\x73\x75\x6c\x74\x73\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x53\x45\x54\x20\x63\x6f\x6c\x6c\x61\x74\x69\x6f\x6e\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x20\x20\x20\x20\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x6f\x6c\x5f\x63\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x20\x2a\x2f\x3b\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d")

// FileTemplatesMysqlCreateViewsTempSQLTmpl is "templates/mysql/create_views_temp.sql.tmpl"
var FileTemplatesMysqlCreateViewsTempSQLTmpl = []byte("\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x56\x69\x65\x77\x73\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x54\x65\x6d\x70\x6f\x72\x61\x72\x79\x20\x76\x69\x65\x77\x20\x73\x74\x72\x75\x63\x74\x75\x72\x65\x20\x66\x6f\x72\x20\x76\x69\x65\x77\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x0a\x2d\x2d\x0a\x0a\x7b\x7b\x20\x69\x66\x20\x6e\x6f\x74\x20\x24\x2e\x41\x72\x67\x73\x2e\x53\x6b\x69\x70\x41\x64\x64\x44\x72\x6f\x70\x54\x61\x62\x6c\x65\x20\x7d\x7d\x44\x52\x4f\x50\x20\x54\x41\x42\x4c\x45\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x3b\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x6e\x6f\x74\x20\x24\x2e\x41\x72\x67\x73\x2e\x53\x6b\x69\x70\x41\x64\x64\x44\x72\x6f\x70\x54\x61\x62\x6c\x65\x20\x7d\x7d\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x44\x52\x4f\x50\x20\x56\x49\x45\x57\x20\x49\x46\x20\x45\x58\x49\x53\x54\x53\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x2a\x2f\x3b\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x53\x45\x54\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x20\x20\x20\x20\x20\x3d\x20\x40\x40\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x3b\x0a\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x3d\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x3b\x0a\x2f\x2a\x21\x35\x30\x30\x30\x31\x20\x43\x52\x45\x41\x54\x45\x20\x56\x49\x45\x57\x20\x60\x7b\x7b\x20\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x20\x41\x53\x20\x53\x45\x4c\x45\x43\x54\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x20\x24\x69\x2c\x20\x24\x65\x20\x3a\x3d\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x7d\x7d\x7b\x7b\x20\x69\x66\x20\x24\x69\x20\x7d\x7d\x2c\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x20\x31\x20\x41\x53\x20\x60\x7b\x7b\x20\x24\x65\x2e\x4e\x61\x6d\x65\x20\x7d\x7d\x60\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x2a\x2f\x3b\x0a\x53\x45\x54\x20\x63\x68\x61\x72\x61\x63\x74\x65\x72\x5f\x73\x65\x74\x5f\x63\x6c\x69\x65\x6e\x74\x20\x3d\x20\x40\x73\x61\x76\x65\x64\x5f\x63\x73\x5f\x63\x6c\x69\x65\x6e\x74\x3b\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d")

// FileTemplatesMysqlDumpSQLTmpl is "templates/mysql/dump.sql.tmpl"
var FileTemplatesMysqlDumpSQLTmpl = []byte("\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x68\x65\x61\x64\x65\x72\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x44\x61\x74\x61\x62\x61\x73\x65\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x64\x61\x74\x61\x62\x61\x73\x65\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x54\x61\x62\x6c\x65\x73\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x74\x61\x62\x6c\x65\x73\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x56\x69\x65\x77\x73\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x76\x69\x65\x77\x73\x5f\x74\x65\x6d\x70\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x52\x6f\x75\x74\x69\x6e\x65\x73\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x72\x6f\x75\x74\x69\x6e\x65\x73\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x69\x66\x20\x2e\x53\x68\x6f\x75\x6c\x64\x44\x75\x6d\x70\x56\x69\x65\x77\x73\x20\x7d\x7d\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x63\x72\x65\x61\x74\x65\x5f\x76\x69\x65\x77\x73\x5f\x66\x69\x6e\x61\x6c\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d\x7b\x7b\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x73\x2f\x6d\x79\x73\x71\x6c\x2f\x66\x6f\x6f\x74\x65\x72\x2e\x73\x71\x6c\x2e\x74\x6d\x70\x6c\x22\x20\x2e\x20\x7d\x7d")

// FileTemplatesMysqlFooterSQLTmpl is "templates/mysql/footer.sql.tmpl"
var FileTemplatesMysqlFooterSQLTmpl = []byte("\x0a\x2f\x2a\x21\x34\x30\x31\x30\x33\x20\x53\x45\x54\x20\x54\x49\x4d\x45\x5f\x5a\x4f\x4e\x45\x3d\x40\x4f\x4c\x44\x5f\x54\x49\x4d\x45\x5f\x5a\x4f\x4e\x45\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x53\x51\x4c\x5f\x4d\x4f\x44\x45\x3d\x40\x4f\x4c\x44\x5f\x53\x51\x4c\x5f\x4d\x4f\x44\x45\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x43\x4c\x49\x45\x4e\x54\x3d\x40\x4f\x4c\x44\x5f\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x43\x4c\x49\x45\x4e\x54\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x52\x45\x53\x55\x4c\x54\x53\x3d\x40\x4f\x4c\x44\x5f\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x52\x45\x53\x55\x4c\x54\x53\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x43\x4f\x4c\x4c\x41\x54\x49\x4f\x4e\x5f\x43\x4f\x4e\x4e\x45\x43\x54\x49\x4f\x4e\x3d\x40\x4f\x4c\x44\x5f\x43\x4f\x4c\x4c\x41\x54\x49\x4f\x4e\x5f\x43\x4f\x4e\x4e\x45\x43\x54\x49\x4f\x4e\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x31\x31\x20\x53\x45\x54\x20\x53\x51\x4c\x5f\x4e\x4f\x54\x45\x53\x3d\x40\x4f\x4c\x44\x5f\x53\x51\x4c\x5f\x4e\x4f\x54\x45\x53\x20\x2a\x2f\x3b\x0a\x0a\x2d\x2d\x20\x44\x75\x6d\x70\x20\x63\x6f\x6d\x70\x6c\x65\x74\x65\x64\x20\x6f\x6e\x20\x7b\x7b\x20\x2e\x44\x75\x6d\x70\x44\x61\x74\x65\x20\x7d\x7d\x20\x69\x6e\x20\x7b\x7b\x20\x2e\x44\x75\x6d\x70\x44\x75\x72\x61\x74\x69\x6f\x6e\x20\x7d\x7d\x0a")

// FileTemplatesMysqlHeaderSQLTmpl is "templates/mysql/header.sql.tmpl"
var FileTemplatesMysqlHeaderSQLTmpl = []byte("\x2d\x2d\x20\x7b\x7b\x20\x2e\x41\x70\x70\x4e\x61\x6d\x65\x20\x7d\x7d\x20\x76\x7b\x7b\x20\x2e\x41\x70\x70\x56\x65\x72\x73\x69\x6f\x6e\x20\x7d\x7d\x0a\x2d\x2d\x0a\x2d\x2d\x20\x48\x6f\x73\x74\x3a\x20\x7b\x7b\x20\x2e\x43\x6f\x6e\x6e\x65\x63\x74\x69\x6f\x6e\x2e\x48\x6f\x73\x74\x20\x7d\x7d\x20\x44\x61\x74\x61\x62\x61\x73\x65\x3a\x20\x7b\x7b\x20\x2e\x4f\x72\x69\x67\x69\x6e\x61\x6c\x44\x61\x74\x61\x62\x61\x73\x65\x4e\x61\x6d\x65\x20\x7d\x7d\x0a\x2d\x2d\x20\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x2d\x0a\x2d\x2d\x20\x53\x65\x72\x76\x65\x72\x20\x76\x65\x72\x73\x69\x6f\x6e\x20\x7b\x7b\x20\x2e\x53\x65\x72\x76\x65\x72\x2e\x56\x65\x72\x73\x69\x6f\x6e\x20\x7d\x7d\x0a\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x43\x4c\x49\x45\x4e\x54\x3d\x40\x40\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x43\x4c\x49\x45\x4e\x54\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x52\x45\x53\x55\x4c\x54\x53\x3d\x40\x40\x43\x48\x41\x52\x41\x43\x54\x45\x52\x5f\x53\x45\x54\x5f\x52\x45\x53\x55\x4c\x54\x53\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x43\x4f\x4c\x4c\x41\x54\x49\x4f\x4e\x5f\x43\x4f\x4e\x4e\x45\x43\x54\x49\x4f\x4e\x3d\x40\x40\x43\x4f\x4c\x4c\x41\x54\x49\x4f\x4e\x5f\x43\x4f\x4e\x4e\x45\x43\x54\x49\x4f\x4e\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x4e\x41\x4d\x45\x53\x20\x7b\x7b\x20\x2e\x43\x68\x61\x72\x53\x65\x74\x20\x7d\x7d\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x33\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x54\x49\x4d\x45\x5f\x5a\x4f\x4e\x45\x3d\x40\x40\x54\x49\x4d\x45\x5f\x5a\x4f\x4e\x45\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x33\x20\x53\x45\x54\x20\x54\x49\x4d\x45\x5f\x5a\x4f\x4e\x45\x3d\x27\x2b\x30\x30\x3a\x30\x30\x27\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x30\x31\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x53\x51\x4c\x5f\x4d\x4f\x44\x45\x3d\x40\x40\x53\x51\x4c\x5f\x4d\x4f\x44\x45\x2c\x20\x53\x51\x4c\x5f\x4d\x4f\x44\x45\x3d\x27\x4e\x4f\x5f\x41\x55\x54\x4f\x5f\x56\x41\x4c\x55\x45\x5f\x4f\x4e\x5f\x5a\x45\x52\x4f\x27\x20\x2a\x2f\x3b\x0a\x2f\x2a\x21\x34\x30\x31\x31\x31\x20\x53\x45\x54\x20\x40\x4f\x4c\x44\x5f\x53\x51\x4c\x5f\x4e\x4f\x54\x45\x53\x3d\x40\x40\x53\x51\x4c\x5f\x4e\x4f\x54\x45\x53\x2c\x20\x53\x51\x4c\x5f\x4e\x4f\x54\x45\x53\x3d\x30\x20\x2a\x2f\x3b\x0a")



func init() {
  if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}


var err error

  
  err = FS.Mkdir(CTX, "templates/", 0777)
  if err != nil {
    log.Fatal(err)
  }
  

  
  err = FS.Mkdir(CTX, "templates/mysql/", 0777)
  if err != nil {
    log.Fatal(err)
  }
  




  var f webdav.File
  

  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_database.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateDatabaseSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_routines.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateRoutinesSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_tables.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateTablesSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_triggers.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateTriggersSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_views_final.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateViewsFinalSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/create_views_temp.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlCreateViewsTempSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/dump.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlDumpSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/footer.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlFooterSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  
  

  f, err = FS.OpenFile(CTX, "templates/mysql/header.sql.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
  if err != nil {
    log.Fatal(err)
  }

  
  _, err = f.Write(FileTemplatesMysqlHeaderSQLTmpl)
  if err != nil {
    log.Fatal(err)
  }
  

  err = f.Close()
  if err != nil {
    log.Fatal(err)
  }
  


  Handler = &webdav.Handler{
    FileSystem: FS,
    LockSystem: webdav.NewMemLS(),
  }
}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
  f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
  if err != nil {
    return nil, err
  }

  buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

  // If the buffer overflows, we will get bytes.ErrTooLarge.
  // Return that as an error. Any other panic remains.
  defer func() {
    e := recover()
    if e == nil {
      return
    }
    if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
      err = panicErr
    } else {
      panic(e)
    }
  }()
  _, err = buf.ReadFrom(f)
  return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
  if err != nil {
    return err
  }
  n, err := f.Write(data)
  if err == nil && n < len(data) {
    err = io.ErrShortWrite
  }
  if err1 := f.Close(); err == nil {
    err = err1
  }
  return err
}

// FileNames is a list of files included in this filebox
var FileNames = []string {
  "templates/mysql/create_database.sql.tmpl",
  "templates/mysql/create_routines.sql.tmpl",
  "templates/mysql/create_tables.sql.tmpl",
  "templates/mysql/create_triggers.sql.tmpl",
  "templates/mysql/create_views_final.sql.tmpl",
  "templates/mysql/create_views_temp.sql.tmpl",
  "templates/mysql/dump.sql.tmpl",
  "templates/mysql/footer.sql.tmpl",
  "templates/mysql/header.sql.tmpl",
  
}

