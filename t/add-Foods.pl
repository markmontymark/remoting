
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
		url => "$baseurl/foods" ),
	POST->new( 
		url => "$baseurl/food",
		query => [
			{Id=>"0",Food=> 4},
			{Id=>"1",Food=> 3},
			{Id=>"2",Food=> 2},
			{Id=>"3",Food=> 5},
			{Id=>"4",Food=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/foods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 foods";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/food/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{Food} == 5 && $data->{Id} == 3, "View a food";
		}]
		),
	DELETE->new( 
		url => "$baseurl/food/3" ),

	TestUrl->new(
		url => "$baseurl/foods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a food, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

