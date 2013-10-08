
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
		url => "$baseurl/routes" ),
	POST->new( 
		url => "$baseurl/route",
		query => [
			{Id=>"0",Route=> 4},
			{Id=>"1",Route=> 3},
			{Id=>"2",Route=> 2},
			{Id=>"3",Route=> 5},
			{Id=>"4",Route=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/routes" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 routes";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/route/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{Route} == 5 && $data->{Id} == 3, "View a route";
		}]
		),
	DELETE->new( 
		url => "$baseurl/route/3" ),

	TestUrl->new(
		url => "$baseurl/routes" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a route, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

