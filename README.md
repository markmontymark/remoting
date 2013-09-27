remoting
========

App for logging my remote office activity


My office is a remote, mobile, ever shifting arrangement of wifi availability, coffee and snack quality, quietude , meeting potential, space coordination.

Objects are
	work location	
	food location
	meeting
	scorecard
	log

Queries are
	went to which locations, filter on qualities
	locations i've not visited yet
	locations i've visited
	what's a good place for <filter on day of week and/or time of day> + food quality
	

	

Locations
 - name
 - address
 - hours of availability

  Space availability
		- number of chairs available
		- number of outlets available

	Wifi quality
		- speed
		- had to reconnect

subclasses: Work Location, Food location, Public transit locations


Public transit access quality


Meetings

	- integrate with a calendar or copy'n'paste into textarea


Scorecard


Log
	- date-in
	- date-out
	- comments
	- refs to location
	- scorecard for quality
		 - wifi
		
	


## dev

	## pull down src
	git clone git@github.com:markmontymark/remoting.git

	## edit files in remoting/
	## then, run dev_appserver.py to test locally
	dev_appserver.py remoting
	./test.sh
	## when done dev/testing, deploy
	appcfg.py update remoting

