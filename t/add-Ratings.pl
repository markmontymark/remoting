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
		url => "$baseurl/ratings" ),
	POST->new( 
		url => "$baseurl/rating",
		query => [
			{Id=>1,Rating=> 4},
			{Id=>2,Rating=> 3},
			{Id=>3,Rating=> 2},
			{Id=>4,Rating=> 5},
			{Id=>5,Rating=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/ratings" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 ratings";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/rating/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Rating} , 2 , "View a rating, validate Rating";
			is $data->{Id} , 3, "View a rating, validate Id";
		}]
		),
	DELETE->new( 
		url => "$baseurl/rating/3" ),

	TestUrl->new(
		url => "$baseurl/ratings" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a rating, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

