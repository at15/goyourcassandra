# pgweb

https://github.com/sosedoff/pgweb

- api
  - can enable session and lock session
  - cors etc.
- bookmark, store files in folder for each bookmark, contains server, username, password etc.
  - [ ] we can add query template as well, i.e. `{name: find cat, query: select * from animals where tpe == cat}`
     - [ ] we can even support build forms so user only need to pass the changed information
- db client
  - history (in memory)
  - supports sshtunnel client/tunnel.go
- can find port
- UI
  - jQuery
  - bootstrap
  - [ace editor](https://ace.c9.io/)

````go
package history

type Record struct {
	Query     string `json:"query"`
	Timestamp string `json:"timestamp"`
}
````

````go
// client/client.go
package client

type Client struct {
	db               *sqlx.DB
	tunnel           *Tunnel
	serverVersion    string
	serverType       string
	lastQueryTime    time.Time
	External         bool
	History          []history.Record `json:"history"`
	ConnectionString string           `json:"connection_string"`
}


func (client *Client) Query(query string) (*Result, error) {
	res, err := client.query(query)

	// Save history records only if query did not fail
	if err == nil && !client.hasHistoryRecord(query) {
		client.History = append(client.History, history.NewRecord(query))
	}

	return res, err
}
````

````go
package bookmarks

type Bookmark struct {
	Url      string          `json:"url"`      // Postgres connection URL
	Host     string          `json:"host"`     // Server hostname
	Port     int             `json:"port"`     // Server port
	User     string          `json:"user"`     // Database user
	Password string          `json:"password"` // User password
	Database string          `json:"database"` // Database name
	Ssl      string          `json:"ssl"`      // Connection SSL mode
	Ssh      *shared.SSHInfo `json:"ssh"`      // SSH tunnel config
}

func ReadAll(path string) (map[string]Bookmark, error) {
	results := map[string]Bookmark{}

	files, err := ioutil.ReadDir(path)
}
````

routes

````go
func SetupRoutes(router *gin.Engine) {
	root := router.Group(command.Opts.Prefix)
    root.GET("/", GetHome)
	root.GET("/static/*path", GetAsset)
	root.GET("/connect/:resource", ConnectWithBackend)

	api := root.Group("/api")
	SetupMiddlewares(api)

	if command.Opts.Sessions {
		api.GET("/sessions", GetSessions)
	}

	api.GET("/info", GetInfo)
	api.POST("/connect", Connect)
	api.POST("/disconnect", Disconnect)
	api.POST("/switchdb", SwitchDb)
	api.GET("/databases", GetDatabases)
	api.GET("/connection", GetConnectionInfo)
	api.GET("/activity", GetActivity)
	api.GET("/schemas", GetSchemas)
	api.GET("/objects", GetObjects)
	api.GET("/tables/:table", GetTable)
	api.GET("/tables/:table/rows", GetTableRows)
	api.GET("/tables/:table/info", GetTableInfo)
	api.GET("/tables/:table/indexes", GetTableIndexes)
	api.GET("/tables/:table/constraints", GetTableConstraints)
	api.GET("/query", RunQuery)
	api.POST("/query", RunQuery)
	api.GET("/explain", ExplainQuery)
	api.POST("/explain", ExplainQuery)
	api.GET("/history", GetHistory)
	api.GET("/bookmarks", GetBookmarks)
	api.GET("/export", DataExport)
}
````