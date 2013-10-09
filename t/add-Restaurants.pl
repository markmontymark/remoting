
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
			{Id=>"0",Restaurant=> 4},
			{Id=>"1",Restaurant=> 3},
			{Id=>"2",Restaurant=> 2},
			{Id=>"3",Restaurant=> 5},
			{Id=>"4",Restaurant=> 1},
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
			ok $data->{Restaurant} == 5 && $data->{Id} == 3, "View a restaurant";
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

