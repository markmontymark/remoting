
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
		url => "$baseurl/neighborhoods" ),
	POST->new( 
		url => "$baseurl/neighborhood",
		query => [
			{Id=>"0",Neighborhood=> 4},
			{Id=>"1",Neighborhood=> 3},
			{Id=>"2",Neighborhood=> 2},
			{Id=>"3",Neighborhood=> 5},
			{Id=>"4",Neighborhood=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/neighborhoods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 neighborhoods";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/neighborhood/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{Neighborhood} == 5 && $data->{Id} == 3, "View a neighborhood";
		}]
		),
	DELETE->new( 
		url => "$baseurl/neighborhood/3" ),

	TestUrl->new(
		url => "$baseurl/neighborhoods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a neighborhood, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

