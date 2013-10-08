#!/usr/bin/env perl

use Test::More;
use lib 't/lib-perl';
use GET;
use POST;
use DELETE;
use TestUrl;
use DoneTesting;
use JSON::XS;
use Go::FindTypes;

Go::FindTypes::finder( './' );

#my $json = JSON::XS->new->pretty;
#print $json->encode( $Go::FindTypes::types );
#exit;


my $port = shift || '8080';
my $baseurl = "http://localhost:$port/simple-service";

my $tasks = [


	## Ratings
	##
	GET->new( 
		url => "$baseurl/ratings" ),
	POST->new( 
		url => "$baseurl/rating",
		query => [
			{Id=>0,Rating => 4},
			{Id=>1,Rating => 3},
			{Id=>2,Rating=> 2},
			{Id=>3,Rating=> 5},
			{Id=>4,Rating=> 1},
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
			ok $data->{Rating} == 5 && $data->{Id} == 3, "View a rating";
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

