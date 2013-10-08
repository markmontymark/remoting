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


	## Users
	##
	GET->new( 
		url => "$baseurl/users" ),
	POST->new( 
		url => "$baseurl/user",
		query => [
			{Id=>0,Name=> 'Count Dracula'},
			{Id=>1,Name=> 'Johana Banana'},
			{Id=>2,Name=> 'Pramad'},
			{Id=>3,Name=> 'Fred'},
			{Id=>4,Name=> 'Ted'},
		]),

	TestUrl->new(
		url => "$baseurl/users" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 users";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/user/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{Name} eq 'Fred' && $data->{Id} == 3, "View a user";
		}]
		),
	DELETE->new( 
		url => "$baseurl/user/3" ),

	TestUrl->new(
		url => "$baseurl/users" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a user, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

