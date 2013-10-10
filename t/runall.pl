#!/bin/sh

./t/add-BusStops.pl
./t/add-Foods.pl
./t/add-MeetingPlaces.pl
./t/add-Neighborhoods.pl
./t/add-Ratings.pl
./t/add-Restaurants.pl
./t/add-Routes.pl
./t/add-Users.pl
./t/add-Workplaces.pl

curl http://localhost:8080/dump
