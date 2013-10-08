
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
		url => "$baseurl/%%collectionlc%%" ),
	POST->new( 
		url => "$baseurl/%%objectlc%%",
		query => [
			{Id=>"0",%%object%%=> 4},
			{Id=>"1",%%object%%=> 3},
			{Id=>"2",%%object%%=> 2},
			{Id=>"3",%%object%%=> 5},
			{Id=>"4",%%object%%=> 1},
		]),

	TestUrl->new(
		url => "$baseurl/%%collectionlc%%" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok scalar @$data == 5, "Added 5 %%collectionlc%%";
		}]
		),

	TestUrl->new( 
		url => "$baseurl/%%objectlc%%/3",
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			ok $data->{%%object%%} == 5 && $data->{Id} == 3, "View a %%objectlc%%";
		}]
		),
	DELETE->new( 
		url => "$baseurl/%%objectlc%%/3" ),

	TestUrl->new(
		url => "$baseurl/%%collectionlc%%" ,
		decode => \&JSON::XS::decode_json,
		tests => [ sub{
			my $data = shift;
			is scalar @$data, 4, "Removed a %%objectlc%%, should have 4 now";
		}]
		),

	DoneTesting->new()

];

$_->run for @$tasks;

