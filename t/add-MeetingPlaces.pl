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

	GET->new( url => "$baseurl/meetingplaces" ),
	POST->new( 
		url => "$baseurl/meetingplace",
		query => [
			{Id=>0,Name=> "Regus Diamond View"},
			{Id=>1,Name=> "Regus Koll"},
			{Id=>2,Name=> "Regus Emerald Plaza"},
			{Id=>3,Name=> "Rebeccas"},
			{Id=>4,Name=> "Home"},
		]),

	TestUrl->new(
		url => "$baseurl/meetingplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data , 5, "Added 5 meetingplaces";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/meetingplace/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is $data->{MeetingPlace} , 5, 'View a meetingplace, name check';
			is $data->{Id}, 3, "View a meetingplace, id check";
		}]
		),
	DELETE->new( 
		url => "$baseurl/meetingplace/3" ),

	TestUrl->new(
		url => "$baseurl/meetingplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a meetingplace, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

