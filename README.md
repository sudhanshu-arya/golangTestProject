
installing the packages that are used   
    go get -u github.com/gorilla/mux    
    go get -u github.com/go-sql-driver/mysql

installing MySQL (skip if you already have MySQL)   
    sudo apt install mysql-server   
    sudo mysql_secure_installation  
    sudo mysql  
    mysql> FLUSH PRIVILEGES;    
    Now the root user is setup to connect to the MySQL shell through a secure password.

change user and password values to your MySQL username and password in trial/dao.go and databaseconfig.py

run following in terminal
    python3 populateDB.py
    cd trial
    go build -o server *.go
    ./server

now go to http://localhost:8000/

following are the endpoints
    http://localhost:8000/home to get worldwide stats
    http://localhost:8000/countrynames to get names of all countries for which we have data
    http://localhost:8000/country/{countryname} to get country stats
    http://localhost:8000/all to get all the stats

use updateDB.py to update the database
