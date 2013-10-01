package Go::FindTypes;

use strict;
use warnings;

use File::Find;
use File::Slurp;
use Cwd;

our $types;

sub wanted 
{
	my $file = $File::Find::name;
	return unless $file =~ /\.go$/;
	#print "found $File::Find::name\n";

	my $state_out = 0;
	my $state_in = 1;

	my $state = $state_out;

	my $lines = read_file($_);
	my $lineno = 0;
	my @current_typedef = ();
	my $current_typename;
	my $current_typetype;

	my $line;	
	for( split /\n/,$lines)
	{
		$lineno++;
		if( $state == $state_out)
		{
			if(/^\s*type\s+/)
			{
				my($typename,$typetype) = $_ =~ m/^\s*type\s+(\S+)\s+(.*)\s*\{/;
				next unless defined $typename && defined $typetype;
				($current_typename,$current_typetype) = ($typename,$typetype);
				if(exists $types->{$typename})
				{
					print "Error: already found type, $typename, at $file, line $lineno\n";
					next;
				}
				$state = $state_in;
			}
		}
		else 
		{
			if(/^(.*)\}/)
			{
				if( defined $1 )
				{
					$line = $1;
					$line =~ s/^\s*//;
					$line =~ s/\s*$//;
					$line =~ s/\/\/.*//;
					push @current_typedef,$line unless $line =~ /^\s*$/;
				}
				$types->{$current_typename} = {
					name => $current_typename, 
					type => $current_typetype, 
					def => [@current_typedef], };
				@current_typedef = $current_typename = $current_typetype = undef;
				$state = $state_out;
			}
			else{
				$line = $_;
				$line =~ s/^\s*//;
				$line =~ s/\s*$//;
				$line =~ s/\/\/.*//;
				push @current_typedef,$line unless $line =~ /^\s*$/;
			}
		}
	}
}

sub finder 
{
	print "args ",@_,"\n";
	File::Find::find(\&wanted,@_);
}

1;
