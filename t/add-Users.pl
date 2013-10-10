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
		url => "$baseurl/users" ),
	POST->new( 
		url => "$baseurl/user",
		query => [
			{Id=>0,Name=> "Opie"},
 			{Id=>1,Name=> "Herr Renfield"},
 			{Id=>2,Name=> "Ted"},
 			{Id=>3,Name=> "Ned"},
 			{Id=>4,Name=> "Fred"},
		]),

	TestUrl->new(
		url => "$baseurl/users" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data,5, "Added 5 users";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/user/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{Name},'Ned', "View a user, validate Name";
			is $data->{Id}, 3, "View a user, validate Id";
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

