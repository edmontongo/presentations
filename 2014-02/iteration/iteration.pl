# while loop

my $condition = 1;
while ($condition) {
    # ...
    $condition = !$condition;
}
# condition is False

#how many lines ?
my $count = 0;
$count++ while(<STDIN>);

my $x = 10;
while ($x > 0) {
    $x--;
}
# x is 0

# maybe you're not sure how many
# iterations you need?
my $x = 100.0;
while ($x > 1) {
    $x /= 3;
}

# for loop
my $sum = 0;
for my $i (1..9) {
    $sum += $i;
}
# sum is 45

# for loop
my $sum = 0;
for ( my $i = 1 ; $i < 100000; $i++ ) {
    $sum += $i;
}
# sum is 4 999 950 000

my $s = "";
for my $elm ("a","b","c") {
    $s .= $elm;
}
# s is abc

# recursive way
sub recsum {
    if (@_) {
        my $v = shift @_;
        return $v + recsum(@_);
    }
    return 0;
}

print recsum(1..10);

#recursive tree sum
my $tree = {};
$tree->{a}->{b}->{c} = 1; $tree->{a}->{b}->{d} = 2;
$tree->{a}->{a}->{e} = 3; $tree->{a}->{f} = 4;
$tree->{g}->{h} = 5; $tree->{j}->{k} = 6;
$tree->{l} = 7; $tree->{m} = 8; 
$tree->{n} = 9; $tree->{o} = 10;
sub treesum {
    my $node = shift;
    if (ref($node)) {
        my $sum = 0;
        foreach my $key (keys %{$node}) {
            $sum += treesum($node->{$key});
        }
        return $sum;
    } else {
        return $node;
    }
}
print treesum($tree);
#55

# The OO way
package OnlyEvens;
use Moose;
has seq => (is=>'rw', default=>sub{[]});
has index => (is => 'rw', default => 0);

sub has_next {
    my ($self) = @_;
    return $self->index < scalar(@{$self->seq});
}
sub next {
    my ($self) = @_;
    my $v = $self->seq->[$self->index];
    $self->index($self->index + 2);
    return $v;
}
1;

my $iter = OnlyEvens->new( seq => [1..10] );
while ($iter->has_next()) {
    print $iter->next().$/;
}

# add 1 to a list
my @v = (1..30);
my @u = map { $_ + 1 }  @v;
# u is now [2,3,4,..,31], v is still [1,2,3]

use File::Basename;

my @v = ("/home","/file","/usr/local");
my @u = map { basename $_ } @v;
my @u = map(uc,@v);
#['/HOME', '/FILE', '/USR/LOCAL']

use LWP::Simple;

my @urls = ("http://cbc.ca","http://gc.ca","http://alberta.ca");
my @pages = map { get $_ } @urls;

# this is why you want blocks with
# few dependencies!
# slower
use Parallel::parallel_map;
sub square { $_[0] * $_[0] }
my @u = parallel_map {square($_)}  1..1000000;
print scalar(@u);
#999999

use List::Util qw(reduce sum);
use List::MoreUtils qw(part);
use POSIX;
sub split_list {
    my ($n, @args) = @_;
    my $i = 0;
    my $total = ceil(@args / $n);
    return part { int( ($i++) / $total ) } @args;
}
sub parallel_square {
    parallel_map { $_ * $_ } @_;
}
sub parallel_sum {
    my @args = @_;
    sum( parallel_map { sum(@$_) } split_list(4, @args) );
}
my $sum = parallel_sum( parallel_square( 1 .. 100 ) );

# a good use!
my @pages = parallel_map { get $_ } @urls;

use List::Util qw(reduce sum);
# 1 + 2 + 3 + ... + 99 + 100 
print(reduce { $a + $b } 1..100);
print(sum(1..100));
#5050
