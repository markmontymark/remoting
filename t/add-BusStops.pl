
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
		url => "$baseurl/busstops" ),
	POST->new( 
		url => "$baseurl/busstop",
		query => [
			{Id=>"0",BusStop=> 4},
			{Id=>"1",BusStop=> 3},
			{Id=>"2",BusStop=> 2},
			{Id=>"3",BusStop=> 5},
			{Id=>"4",BusStop=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/busstops" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 busstops";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/busstop/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{BusStop} == 5 && $data->{Id} == 3, "View a busstop";
		}]
		),
	DELETE->new( 
		url => "$baseurl/busstop/3" ),

	TestUrl->new(
		url => "$baseurl/busstops" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a busstop, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

