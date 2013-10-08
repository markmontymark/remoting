use HttpTask;

package TestUrl;
use Mo qw'build default builder coerce is required';
extends 'HttpTask';
has tests => (default => sub{[]});
has decode =>(default => sub{sub{@_}});
sub run
{
	my $self = shift;
	my $content = GET->new(url => $self->url )->run;
	$_->($self->decode->($content)) for @{ $self->tests};
}

1;
