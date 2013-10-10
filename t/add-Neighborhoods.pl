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
			{Id=>0,Name=> "Normal Heights"},
			{Id=>1,Name=> "Hillcrest"},
			{Id=>2,Name=> "North Park"},
			{Id=>3,Name=> "Downtown"},
			{Id=>4,Name=> "Univ. Heights"},
		]),

	TestUrl->new(
		url => "$baseurl/neighborhoods" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data , 5, "Added 5 neighborhoods";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/neighborhood/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Name},'Downtown', 'View a neighborhood, name check';
			is $data->{Id},3, 'View a neighborhood, id check';
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

