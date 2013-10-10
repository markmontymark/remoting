#!/usr/bin/env perl

use Test::More;
use lib 't/lib-perl';
use GET;
use POST;
use DELETE;
use TestUrl;
use DoneTesting;
use JSON::XS;

##use Go::FindTypes;
##Go::FindTypes::finder( './' );

my $port = shift || '8080';
my $baseurl = "http://localhost:$port/simple-service";

my $tasks = [

	GET->new( 
		url => "$baseurl/restaurants" ),
	POST->new( 
		url => "$baseurl/restaurant",
		query => [
			{Id=>0,Name=> "Specialtys"},
			{Id=>1,Name=> "Venissimo"},
			{Id=>2,Name=> "Monicas"},
			{Id=>3,Name=> "Grant Grill"},
			{Id=>4,Name=> "Brooklyn Bagel"},
		]),

	TestUrl->new(
		url => "$baseurl/restaurants" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 restaurants";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/restaurant/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Name} , 'Grant Grill','View a restaurant, name check';
			is $data->{Id} ,3, "View a restaurant, id check";
		}]
		),
	DELETE->new( 
		url => "$baseurl/restaurant/3" ),

	TestUrl->new(
		url => "$baseurl/restaurants" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a restaurant, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

