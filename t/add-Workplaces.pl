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
		url => "$baseurl/workplaces" ),
	POST->new( 
		url => "$baseurl/workplace",
		query => [
			{Id=>0,Name=> "Home"},
			{Id=>1,Name=> "Gelato Vero"},
			{Id=>2,Name=> "Specialtys"},
			{Id=>3,Name=> "Regus Koll"},
			{Id=>4,Name=> "Regus Emerald Plaza" },
		]),

	TestUrl->new(
		url => "$baseurl/workplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data , 5, "Added 5 workplaces";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/workplace/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Workplace} ,5, 'View a workplace, name check';
			is $data->{Id} , 3, "View a workplace, id check";
		}]
		),
	DELETE->new( 
		url => "$baseurl/workplace/3" ),

	TestUrl->new(
		url => "$baseurl/workplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a workplace, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

