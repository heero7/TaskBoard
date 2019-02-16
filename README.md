# TaskBoard
A board to keep track of all your things to do :)

## Description
TaskBoard aims to make your life more organized at work and prioritizing your work to keep track of things you are doing.

## Installation Prerequisites
1. Make sure Go is installed, the latest version
2. Install Postgres SQL Server
3. Install React 
## Installation

Use govendor for package management
```
go get govendor 
go get github.com/heero7/TaskBoard
```
After installing get all the needed packages
```
go govendor get-dependencies
```
Do a build to make sure everything is smooth, good? Let's keep going.

Create a config file with your database information
```
{
  "database_info" : "credentials",
  "port" : "portinfo"
}
```

Start the SQL server

Start the app

```
go run main.go server.go middleware.go
```
## Debug
VSCode
- Open from the server folder
- Configure launch.json for Go
- Start debugging

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
