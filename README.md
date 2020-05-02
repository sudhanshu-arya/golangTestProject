## GOLANG TEST PROJECT

## Install the packages that are used
- go get -u github.com/gorilla/mux
- go get -u github.com/go-sql-driver/mysql

## Install MySQL (skip if you already have MySQL)
- sudo apt install mysql-server   
- sudo mysql_secure_installation  
- sudo mysql  
- mysql> FLUSH PRIVILEGES;    
- Now the root user is setup to connect to the MySQL shell through a secure password.

### Change user and password values to your MySQL username and password in 'trial/dao.go' and 'databaseconfig.py'

## Run following in terminal
- python3 populateDB.py
- cd trial
- go build -o server *.go
- ./server

### Now go to http://localhost:8000/

## Following are the endpoints 
- http://localhost:8000/home to get worldwide stats   
- http://localhost:8000/countrynames to get names of all countries
- http://localhost:8000/country/{countryname} to get country stats    
- http://localhost:8000/all to get all the stats

## Use 'updateDB.py' to update the database
