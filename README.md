# ⚡ Technical Test for Deviget ⚡ 

- Dependencies

  * [Golang](https://nodejs.org/)

  * [Docker](https://www.docker.com/)

#Tecnical decision

* Layer Arquitecture
  * Routes
  * Service
  * Game_engine
  * Repository
     
* DB
  * PostgressSQL
  
the repository have a interface
to easy change DB implementation if needed
use go modules to easy manage dependencies   

## init postgress DB on Docker

```
docker run --name postgresql -e POSTGRES_PASSWORD=rootpass -v my_dbdata:/var/lib/postgresql/data -p 5432:5432 -d postgres:11
```

##Post Routes

	"http://localhost:8080/newPlayer"
	"http://localhost:8080/newGame"
    "http://localhost:8080/clickGame"
    "http://localhost:8080/pauseGame/:gameId"
    
    
##Post ClickGame example body
    
    {
      "game": {
        "gameId": "aa362a01-23c1-43ab-8589-bdb9ebbcba42",
        "rows": 6,
        "cols": 6,
        "mines": 6,
        "status": "STARTED",
        "playerId": "b36bfd67-a379-40f8-ae2a-5949d0ac8c20",
        "clicks": null,
        "timer": "2019-10-04T00:00:00.000000Z",
        "duration": null
      },
      "positionx":3,
      "positiony": 5
    }
 
##Post NewGame example body 
    {
    "playerId": "b36bfd67-a379-40f8-ae2a-5949d0ac8c20",
    "rows":6,
    "cols":6,
    "mines":6
    }

##Post NewPlayer example body
    {
      "userName": "FancyPlayerName"
    }
    
##Post NewGame example  

    http://localhost:8080/pauseGame/aa362a01-23c1-43ab-8589-bdb9ebbcba42

         