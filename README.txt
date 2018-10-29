Application has two pages:
	"/" - the main page (https://stormy-anchorage-43452.herokuapp.com) which displays the html page (application introduction)
	"/routes" - the page which accepts the source and destination parameters and returns the json whith calculateed routes (https://stormy-anchorage-43452.herokuapp.com/routes)


Package description
 - config - it contains configurataion of html temlates and logging functionality
 - controllers - it contains web controllers. Currently there is only one controller to accepts GET request to "routes" page and responses with json
 - logging - logging wrapper to log easily with common format
 - model - it contains structs which are used by application
 - service ->
				file pathsRouter.go - iterface PathRouter to get route for selected source and destination
				file routers.go - service which calculates routes for selected source and list of destinations. It uses PathRouter interface to get route for one destination
 - service/osrm - implementation of PathRouter which uses OSRM service to calculate route for source and destination
 - template - html templates
 - test - unit test for sorting functionality
 main.go - main class to start server. I use the julienschmidt/httprouter 
 

	
