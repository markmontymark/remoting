use JSON::XS;
use Task;

package HttpTask;
use Mo qw'build default builder coerce is required';
extends 'Task';

has method => ();
has url => ();
has query => ();

sub run
{
	my $self = shift;

	my $process = join ' ','curl -s --request', $self->method;
	unless($self->query)
	{
		$process .= ' ' . $self->url;
		print STDERR "$process\n";
		return `$process`;
	}

	my @retval = ();
	for(@{$self->query})
	{
		my $q = JSON::XS::encode_json($_);
		$q =~ s/"/\\"/g;
		my $process_post = join ' ',$process, '--data', '"',$q ,'"', $self->url;
		print STDERR "$process_post\n";
		push @retval, `$process_post`;
	}
	return join "\n",@retval;
}

1;
