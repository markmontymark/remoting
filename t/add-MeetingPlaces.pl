
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
		url => "$baseurl/meetingplaces" ),
	POST->new( 
		url => "$baseurl/meetingplace",
		query => [
			{Id=>"0",MeetingPlace=> 4},
			{Id=>"1",MeetingPlace=> 3},
			{Id=>"2",MeetingPlace=> 2},
			{Id=>"3",MeetingPlace=> 5},
			{Id=>"4",MeetingPlace=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/meetingplaces" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 meetingplaces";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/meetingplace/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{MeetingPlace} == 5 && $data->{Id} == 3, "View a meetingplace";
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

