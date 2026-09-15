# GopherCon 2021: Unbound Concurrency in Go - Johnny Boursiquot & Aaron Schlesinger
[video](https://www.youtube.com/watch?v=PAUjYyBfELk)

0:04
I was thinking right that every year we hear about how we we get a lot of new
0:10
folks in the uh go joining the go Community for the very first time um and
0:15
they're they're they're curious about go they're they're picking up you know some tutorials and picking up you know some
0:21
of the idioms of the language and things like that um and everybody comes in they're like hey go routines and things
0:26
concurrency that's one of the the draws of the language right MH and uh and a lot of times you know we end up also
0:33
seeing uh um some misuses right of the concurrency U um concurrency is great
0:38
and go um but you can easily get yourself into trouble right when you have too much concurrency there is such
0:43
a thing as too much concurrency right you know we get this toy we Engineers love toys right and we get this Gokey
0:50
word that's the toy that's all it takes one thing we could do is about introducing sort of a a simple little
0:57
program simple little concept you know um to sort of add some sophistication to it as we go uh and and to sort of
1:03
introduce some some patterns and some ways of reasoning about you know a concurrent program uh so I thought we'd
1:10
build like a quick little Port scanner together what do you think sounds great sounds great plan so before we jump into
1:15
building that let's actually quickly go over what a port scanner actually is so I have a quick little slide here um
1:21
that's going to um sort of illustrate sort of the communication that happens right between a client and and a server
1:27
right and the client could be another server right and whatever the case may maybe so one person initiates some Network contactor contact to another to
1:35
another um server maybe you know on the other side of the world wherever the case may be right so usually there's a
1:41
there's a there's a synchronization packet uh that sort of finds this way from the client over to that server if
1:46
the server wishes right to engage in our communication right maybe there's no firewall or maybe there's no rule that
1:51
says it it can't talk back um to that uh that initial uh request it sends back an
1:57
acknowledgement right of that initial synchron packet right and the client now says okay I'm going to acknowledge that
2:04
you sent me an acknowledgement to my original um synchronized packet now we can actually talk back and forth with
2:11
with each other right so that but that process of actually uh connecting over to that client that needs to happen on a
2:17
port right so the server must have a process right running locally that has
2:23
basically is bound to that port to a port right whatever Port you're trying to connect over there must be a process
2:29
list listening on that Port right for the the host to actually receive your incoming request and says okay yeah I I
2:36
do have a process running on that Port so I'm going to facilitate that communication right so when we talk about Port scanning really what we're
2:42
doing is basically saying hey um can I given a range of ports can I can I go through the list quickly right for a
2:49
given host and sort of enumerate right the list of open ports enumerate the list of ports that I actually I can
2:55
actually initiate some sort of connection on onto not that you're actually going to exchange any data you just want to see he just probing right
3:01
is this port open so we're going to be doing some heavy duty kind of lowlevel networking we're going to probably
3:07
expect a lot of times to fail to not be able to connect but we want to figure out what are those times and what are
3:13
those ports that we can connect on successfully right exactly exactly and go right exceeds at this particular kind
3:20
of application right it goes the cent of Library provides a lot of the building blocks right to to facilitate right
3:27
create the creation of these cons programs and a lot of our work is today is going to be based on the net package
3:33
and the dial function because it abstracts away a lot of this sort of you know handshake process the back and forth you know handling you know drops
3:40
and packets and all that stuff like all that stuff is sort of a lowlevel concerns that we as a developers
3:45
basically using the standard Library package don't have to worry about and we're going to be doing almost everything with the standard Library
3:50
here absolutely there maybe one or two but you know 90% of our work is standard library library no no external
3:56
dependencies y indeed indeed so let's flip back over right so the way I I usually like to do
4:02
these things is to basically you know our our our task right what we've been asked to do is to create this sport
4:08
scanner and I usually want to start with the simplest thing that could possibly work what is what is the minimal effort
4:13
I could put in to just test out the idea right kind of similar in spirit to what Ben was just talking about is you know
4:20
we don't necessarily need to go and start up another database and all that let's just take what we got run on the same machine here we're kind of
4:26
following the Same Spirit let's get the simplest thing working and then evolve it over time we're gonna see nine
4:33
evolutions of this as we go as we go yep exactly exactly so what we have here is
4:39
is really is the The Bare Bones of what could possibly work so we know we're going to be using the net packages dial
4:44
function to actually you know initiate that connection it's going to be a TCP based connection and obviously the the
4:50
St Library the net package supports all kinds of connection type CCP is only one of a handful um so we're going to B
4:57
basically um hardcode the port that we for now it's Port 5432 um um I happen to
5:03
have a postgress server running locally on that Port right uh we're going to see if we can actually establish a connection on on our Local Host here uh
5:10
and on that port and see and see what we get right and then we'll try it you know with a on a port that we know is closed
5:16
and see what the response is that way we know what what to look for for an open port and a closed Port right how does
5:21
that sound that sounds great and look how I mean this is uh well this is 18 lines including the new line including
5:27
Imports but really we've got line 10 through 16 some of that's error checking yep the ghost standard Library makes
5:34
this pretty easy makes it absolutely very very easy to do so let's actually do a quick go run here uh let's do this
5:40
is the simplest thing right and then let's just run this and it says it's open right and then what I'm going to do
5:46
is I'm going to change this I'm going to change my port here and say 5433 right and then I'm going to run this again and
5:53
then now we're going to consider that Port Clos so let's talk about the output here so when we call net. right so let's
5:59
look at the the the the the the results the the responses that come back so we're going to get back a value of type
6:05
netcon right that's going to be the connection value that we can use the object if you will um that we can use to
6:12
to to do other things with maybe we want to start exchanging data Maybe we want to send send some some some um text over
6:17
and see what we get back but that we don't particularly care about that we our job is to find out is this sport
6:23
open or not right so we do have some methods on that struct that comes back we're not going to use them today but
6:28
but we could you know in the future we want to do some communication with the server once we connect to it or what have you exactly exactly what we're
6:34
going to worry about more so today is the error value in other words for whatever reason regardless of why an
6:41
arrows return we're going to consider that Port closed for inter purposes because whether or not it is closed because we couldn't reach the host or
6:47
because you know there's nothing listening on or bound to that port on the host we don't particularly care an
6:53
error was returned maybe the we timed out trying to get a response back it doesn't matter we're going to consider
6:58
that Port closed for purposes that's why we had the different responses here 5432 again I have a a postgress server
7:04
running that we were able to sort of successfully get a connection back but in a cases where it was a different port that we didn't we couldn't get a
7:10
response back we're going to consider that close so that forms the the foundation the basis for the evolution
7:16
right of our program in in the next steps that's great this is not a lot of code I you know I've used go for a while
7:22
so I happen to think this is very understandable I think someone from Java or even python C C++ they can come in
7:30
and get a pretty good grasp on this pretty quickly absolutely absolutely but the thing is we haven't done any sort of
7:36
scanning yet right so you know remember we're typically trying to enumerate a list of ports right so let's let's
7:42
basically you know augment our previous program by simply saying hey let's just range over a couple hundred ports right
7:48
that's the that's the that's I think that's more appropriate for a scanner right so let's actually see right when
7:53
we run this version of the program what we get
8:00
all right so let's go full screen on this so here we are we've just scanned a couple hundred ports right and uh we're
8:06
seeing that the same um sort of an error result that we can expect when the port is is is we consider that Port closed
8:12
right so if we scroll down we scroll down we're scanning ports we're scanning ports we should see 5432 here it is Boom
8:18
this is a port that that was open before so now we've enumerated a list right now we have ourselves a list of all the
8:24
ports that that um we consider to be open but are you noticing uh pattern
8:30
here with scans yeah I mean we're doing these in order one after the other sequential in this case you know we're
8:36
doing 5300 to 5500 it was pretty fast on your machines pretty modern machine
8:41
right but you know what if we're getting into the the cases where we want to do tens of thousands of ports maybe I mean
8:49
a lot of ports a lot of ports hundreds of thousands might be a little out of the realm of reality yeah you know let's
8:55
say 5,000 10,000 ports you know also you've got some number of cores greater
9:01
than one on this machine we're only I can probably use you know more of those things right especially we're writing go so you know this might be a good
9:07
opportunity to explore making this concurrent and probably parallel as well yeah exactly and and the other thing is
9:13
is too is that there is there is no real dependency between one port skin and the
9:19
next right so these are these are there's a there's a term for for for these I love this term you want to you
9:25
want to take this one um how about we spit it okay I'm embarrassingly
9:31
parallel oh we're not splitting the word the word we're supposed to go together okay get it together two three
9:38
embarrassingly parallel we didn't rehearse that all yeah coup off the go right so so we we
9:46
know that we we there's no dependency from one scan to the next so there's there's an opportunity here right to
9:51
sort of make the work concurrent right enable that parallelism and because I have multiple CES in my machine we are
9:57
going to hit some parallel workloads so let's see if we actually turn this sequential program this procedural
10:02
approach to alpar scanner into something a bit more concurrent and this is where we get to use that toy that keyword the
10:07
ah yes indeed this is where we actually get to use our toy so the I've been told right
10:16
that the only thing we need to to make a a go program go faster make it concurrent right just slap some go
10:22
keywords in places and off we go let's see let's see what happens see what happens so what we did here right we
10:28
basically created ourselves a little sort of anonymous function that we are going to invoke all we did was we moved
10:33
our our Port scanning code right into that Anonymous function so that we can actually you know uh sort of a run that
10:39
concurrently right outside of the the main go routine right and then making sure not to not to have a variable
10:46
shadowing U um going on here making sure we're passing the the actual copy of the port that we want to send in so that we
10:51
don't get us into trouble with the with the loop uh variable and then basically is doing the same thing right uh if the
10:57
port if we got an error back Market has if we if we actually got got to here you
11:02
know consider it to be open and then we going to say hey we're done after the uh we we've uh we've done the work so let's
11:10
do see what happens oh uh what happened there we got no output from any of the
11:17
go routin from any of the go that was a little too fast wasn't it it was a little too fast so what the heck is going on here can you tell me what I did
11:24
wrong well when you have a go keyword that means that you have scheduled the go routine the go run time is going to
11:30
schedule the go routine doesn't mean it's going to start even start running it doesn't mean it's going to wait for it to be over and that's kind of one of
11:36
our first uh rules of the day right we want to know we always want to try to know when the go routines are done not
11:44
just schedule and not just started we don't necessarily need to know when they're started we need to know when they're scheduled and when they're done
11:50
exactly exactly and in this case here we we just scheduled them and we just kept on going I'm merry way for the best just
11:58
HED up for the best because a lot a lot of folks often don't realize that the the the moment you start running a go
12:04
program you are running go routines right in in our in most cases right if you just have the your main if you have
12:09
everything happening inside of your main girl routine that is a girl routine right so now when you fire off other
12:14
girl routines you need to somehow synchronize them you need to somehow allow one to wait for the other right or
12:21
others whatever whatever they're launching you need to synchronize them somehow so thankfully the go standard
12:28
Library makes that process a lot easier and there's a package I'm thinking that actually is going to help us add some
12:33
synchronization to to this to this problem let's go check out what that package is let's go check out what the package is and you know most of what we
12:40
do from this point on is going to be all about synchronizing these go routines we kind of got this pattern where we've got n plus one go routines where n are the
12:48
workers and then that one is the main go routine that you were talking about and we're going to do a lot of synchronization between the one and the
12:54
n and exactly exactly so the standard library has the sync package right um
13:00
Bic the same package is going to provide us all of the building blocks that we're going to build on moving forward uh for
13:05
this particular program so this is where we also start to add some sophistication right to to our to our project um
13:10
basically I'm going to parameterize the host you know the front Port the two port and other versions of this program
13:15
are going to have similar U um sort of a configurability as well as we go so I've got some defaults here so I'm basically
13:22
saying my my default front Port is going to be 8080 my default two Port is going to be um 80 8090 so if I don't override
13:27
these things I'm going to head just a small range of 10 ports of scan so and basically the rest I'm going to not
13:33
going to go through every every line here the rest I'm just doing some checks and making sure that okay you're not sending me a a letter as opposed to a
13:38
number for skinning and that kind of thing this is where we actually initialize well I I should be careful here this is where we declare right a
13:44
weight group so why why don't we have to actually initialize this weight group well if we look in the docs for the sync
13:51
package and we scroll down to the weight group type it says that the zero value is initialized right right the zero
13:56
value is useful right which is another idiom in NE community that basically you hear folks talk about make the zero value useful this is actually the stand
14:03
Library living up to that to that Mantra of of making that zero value useful so we don't have to give it a value or anything like that we can simply declare
14:09
it and when we actually start using it like when we say wg. add for example um
14:14
it's already there you're not going to get a a panic because you're referencing calling a meth a nil object or anything
14:20
like that so you're good to go right so before we talk about the ad though uh we have to think about what what are we
14:26
adding what what is it that we're going to be waiting on right so the wig group right the way that works is that you have to specify some sort of counter you
14:33
have to specify how many things you're waiting on so how do we get that value how do we how do we know ahead of time
14:39
what what what how many things to wait on yeah so we've got our list of ports so we're basically saying well we we
14:45
we've got X number of ports I should say n number of ports right we're going to want to start up for this exercise we're
14:51
going to start up a worker per port in the list so we're going to say all right well how long is the list how many ports
14:57
did you configure when you ran this code how many ports did you want to scan we're going to do our WG ad and pass the
15:04
length of that list exactly exactly and one of the empty patterns that we want to we generally want to avoid is one
15:10
where we're actually calling this ad right inside of the the the the the
15:15
invocation of the go routine right so you know you can you're going to see out there as you start to learn go and you're going to start to see some some
15:21
tutorials and some code Snippets that they don't always sort of provide they solve the problem yes um but they don't
15:27
always provide sort of the best practices so this is one of the the ones where you actually should try if you know ahead of time especially if you're
15:33
using a we group that means that you have some knowledge of how many of these things you're about to launch right just
15:39
basically call the ad method just once because you know ahead of time how many gotin you're going to be waiting on in this case it's our range of ports right
15:45
do the quick little calculation but if for some reason you don't make sure to call it before you launch that go
15:51
routine right that is if you ever call wg. add like right before the defer wg.
15:58
dun or inside the go routine somewhere else probably not correct that's a pretty good bet that there's something
16:04
wrong something yeah something smell smelly with the coat there yeah smelly is a good way to put that yeah exactly
16:10
so we're doing we're doing we're not doing anything special here again we're going to iterate over our list of ports
16:15
and for each one port we want to scan we're going to fire off a gr routine and let the work happen inside of that gr
16:20
routine and we have to ensure that when we actually the work is done with an AG
16:25
gr routine that we decrement the counter that the weight group is is using right uhuh to help us wait on on on the work
16:32
to be done right and where is that work H where is the waiting happening well that happens on line 54 right you have
16:40
to remember that you know if you can is you can you can make use of the weight group you can add the number of counters but if you don't wait what's the point
16:46
there's no point so that that's that's the that's the kudra right there this is life 54 is really what you want this is
16:52
where you telling the main go routine that you know that you've just launched a whole bunch of different uh go routines from from line 43 on on to 52
17:01
right that you need to actually wait for these things and inside of those grow routines this is what the WG that done
17:06
is doing for each one of those gotin every time the work is done subtracts from the counter and then at some point
17:12
weight stops blocking right and basically hence control back over to our main go routine so blocking what does
17:19
that actually what does that mean when you say you know weight is blocking and then it stops blocking when the when the
17:25
weight group goes back down to zero right so our scheduler right is always scheduling okay who gets to go next who
17:30
gets to actually do the next piece of work right and so our schedule is the one controlling right who who who is
17:37
next in line who gets to do some work when do they get paused when is control handed over to to to another girl
17:42
routines and whatnot right when we talk about blocking here and especially in our mango routine basically we're telling the schedule hey we want right
17:50
to to to get some work finished right and but this this thread needs to be paused right needs to be stopped in its
17:56
tracks it's to wait right here until we're ready to move forward so we are literally going to not execute anything
18:02
after the wait until the other go routine those end go routines are all done and they've called wgd each of them
18:08
and number of times that way the way group knows okay now we can proceed the co- scheder hands control back over to
18:14
the manger routine and off we go again all right all right so now let's let's now see right um the Improvement right
18:20
to this version of the program so we by default we're going to have hey there we
18:26
go that's our that's our list of 10 points before we just saw done yeah right wasn't too useful right basically
18:31
everything just you know we didn't know we didn't wait for the go to be done now we we have we have um we actually waited
18:37
on on this and this is going from 8080 to 8090 so just small range yeah small
18:43
I've got enough I've got enough cores on my machine this work was done pretty quickly right so this is the part where
18:49
uh I want to actually see what happens again this is a very small range of
18:55
ports just just just a handful right what and now I now package this ship it right
19:01
and and and all of a sudden the customer right the the the user of this program decides to throw you a curve ball or
19:08
rather they decide to actually use it an unexpected way right otherwise known as testing production right exactly exactly
19:15
so they decide okay hey you know what I'm gonna use this Nifty you know Port scanner and I'm gonna I'm gonna scan
19:21
from say one to 65,000 so if we try to do this now oh actually I need to
19:27
specify the argument right I say I think yeah I think you're right it's from for I might
19:33
even have that in my hey there we go hey there we go yeah we're going to use this one so we're going to go from 5,000 from
19:39
5,000 to um port to 65,000 let's actually see how our program behaves as
19:45
written so here we go so oh here we go that did not go well for us that's
19:51
called taking up too many resources that is called taking up too many resources right so this error is very different
19:57
from what we were getting before now this is our Host this is our operating system telling us ah you're using up too
20:03
much right I haven't allocated that many file handles for you to be able to open right the operating system is protecting
20:09
itself in this case and you know it's basically going to prevent your program from operating you know optimally right
20:15
so we're handling this particular error um we're outputting you know what went wrong but you know obviously our program
20:21
is not operating to spec and this isn't a great experience either right we would we would like to still scan all those
20:28
ports we would like to still do it concurrently and in parallel right but obviously the operating system is not
20:34
able to give us the resources we need to do it in this unbounded way unbounded exactly hence the title of this talk you
20:40
know this is this is exactly this what unbounded concurrency looks like right but you don't have you're not
20:45
constraining your program to to the context in which it is going to run in this case yes we could launch 65,000 go
20:53
routines right to do the work but that's not the point right the point is for us to actually be good C citizens within
20:59
the environment in which we are actually operating and running right so we need to constrain this program in some way we
21:05
need we need we need to somehow control the number of gr routines that we we we have doing the work while enabling right
21:12
for the concurrent work and for the Speedy work to get done without running a file of system resources that we've
21:17
been allocated yeah so I'm thinking we need to do we need we need
21:23
to introduce some some new Concepts right for for this particular program so
21:28
I'm thinking this is where we introduced the concept of a worker pool so what is a worker pool Aaron well before we had
21:35
this unbounded thing where we were just doing go and we launched a go routine and the number of go routines we
21:40
launched was dependent on what the user gave us right so we did whatever it was 60 almost 65,000 yeah 6,000 so yeah it
21:48
was it was a very wide range yeah right so we obviously cannot do that as we saw so for work with worker pools we're
21:54
goingon to launch a static number of goer teens and let's say five and we're going to feed ports to those go routines
22:01
so we're never going to have more than five but those five are going to work in parallel on all the ports that we need
22:06
to do exactly exactly so we we're we're we're we're creating a set number right
22:11
of of go routines that can go out and do the work and as they get done they come back and pick up some work right so for
22:17
in our case here we we've added a new flag right so we've actually modified our our our our flags a little bit
22:23
basically just to make sure I I don't have to type you know front port and two Port all over the place I've just you know made some improvements to to
22:29
reading of of our ports here um and uh but the important thing here is that on line 24 you see I've introduced a new a
22:36
new flag that's the that's the workers flag the number of workers flag and I'm defaulting that to the number of CPU logical CPU cors my machine right so I
22:43
think I think this is an eight car machine so that means that by default I'm going to have uh eight workers uh in
22:49
in my pool which I think is a good place to start because there's no right or wrong number for this this is all going
22:55
to be what is a sweet spot cuz you might be running in production running a a 64
23:00
core you know multi like super beefed up machine or set of machines right that is
23:05
going to be a very very different performance profile than a than a developer laptop yeah and and you know folks will have to do trade-offs they'll
23:12
have to measure how fast they want to run versus how much contention for go routines per cord do they want to run or
23:19
per OS thread and so forth so that's a whole another talk that's a whole another talk yeah that's that's that's that's something that will basically
23:25
leave to you to figure out for your particular situation or what particular workload what the lot the the the
23:31
correct number of workers is for you but the idea here is that basically what we're going to be doing and and and I've
23:37
also added some additional things here is that the for in this particular main function I'm also capturing a signal for
23:43
for when um I I basically stop and send and interrupt to the program I can still capture the result so I can actually
23:49
print them out so that's all that's happening here um and I've hidden the um sport scanning right behind a nice
23:54
little helper function that's just going to give me back a slice of integers right based on the the the arguments that I passed in when I invoked the
24:00
program so but the the real fun stuff starts on line 46 this is where I'm
24:05
actually using make to actually create myself a buffered channel right so
24:10
normally speaking right an unbuffered channel right facilitates some communication between the gr tees where
24:17
only the the the where the information is exchanged when there is somebody on the other side waiting to receive a
24:23
message how does the buffer Channel behave differently so buffer channel has kind of a queue built in M where I can
24:29
start pushing onto the channel and it'll fill up the once the gets full then I
24:34
block to S right but on the other side of course there could be someone taking off of the queue receiving from that
24:41
channel so if someone is receiving from the channel at a similar Pace as what I
24:46
sending so let me say that better if someone's receiving at the same Pace as I send then I'm not going to BL side
24:53
exactly so it's just a little bit of a cue or a buffer as as the name goes in between the sender and the receiver and
25:00
it can be used in this case of course to do our workers to do our workers exactly exactly so what we're going to do is
25:05
basically we're going to we're going to launch right an N number of grines and that's going to be sort of a driven by
25:13
the number of of uh um workers that we have in in our pool and we're going to be using the ports Chan right as the
25:20
mechanism to feed ports into and for workers to pick ports to scan out of
25:27
right so we need to feed some ports in there and others are going to be picking and picking up from there so the way we
25:32
do that is by basically we we've we've abstracted away we've tucked away all of that behavior right into this worker um
25:38
function which in basically we're calling we're calling a new worker basically we're we're preparing a new
25:44
workers start accepting work right based on the the capacity right based on the number of workers that we have this is
25:50
why we have what we're doing on line 50 here right so let's actually go take a look right and then we'll come back to line 53 let's go take a look on for the
25:57
definition of work right so what we if we take a look at the signature of this function here
26:02
we're going to take in the host yes that's that's we we know we know we need the host but the the directionality
26:08
right of our Port Chan and the directionality of our results Chan that looks that looks interesting right yeah
26:14
what's up with those little arrows can can you can explain those things to us Aaron well you said it directionality
26:19
right so we can enlist the compiler to ensure that for the first one for porch
26:24
Chan that we only read from it read only yeah and that that that directionality that arrow is on the left side of the
26:31
chin keyword so if inside of worker if we try to write to Port Chan our code won't compile which is really nice you
26:38
know you get the compiler to write a whole class of tests for us right and then similarly for results Chan well the
26:44
arrow is going into the Chan instead of out so that's the one where we can only write to we can only send on it and I
26:50
love making the compiler do work for me so this is a huge Boon for me as it should be for all of us exactly exactly
26:57
and what's happening here so basically um we're ranging over a porch Chan right
27:03
and unlike other four um ranges the when you read from a channel only one value comes back right there's no index and
27:09
value it's just a value that you're picking up from from the channel and now we're just ranging on on the port and as
27:14
long as on on the channel and as long as there's values coming in right we're gonna we're going to be doing some work
27:20
right which basically we're doing the same thing we're doing before calling atal inspecting a response and if we
27:26
notice that okay there is the we can the port close we're going to send a zero right onto the results chain indicating
27:32
that basically that that whatever Port that was we don't care about it anymore because you know the port is closed so
27:37
we're going to send a zero in there because we need to send back right an integer we need to send an integer into that port to indicate that we're done
27:43
with that work right so we send a zero in there and we continue on to the next one right so we're going to keep iterating until that that Port is is
27:51
until we receive a signal right that this port is closed and we're going to see where that signal is actually sent
27:56
back where this worker is actually is actually um is actually invoked right
28:02
but what we're going to do here is basically if we get to the point on line 112 where we manage to get past the
28:08
point where we do the error check and and we we notice a closure if we get packed to to the2 we get get pass in
28:15
down to 112 we're going to consider that port to be open so we send that P whatever P happens to be Port 5432 for
28:21
example that's going to go into the results chain right and and this is where basically we're feeding feeding
28:27
back right the the the the result of our operation of our scanning operation we're sending back the result to that
28:32
into that channel so that wherever the worker is being invoked we can actually do something and read some values from
28:37
that so we've got a consumer here worker is the consumer there's going to be n number of CPUs workers running they're
28:43
all going to consume and you know if you send a port to the ports Chan one of the N workers is going to pick it up do its
28:50
thing send back the result yep so let's look at the producer now yeah so what's going to happen so in order for us right
28:57
to steadily feed in ports into this particular um um ports CH that the
29:03
workers are actually uh um waiting on we actually need to range over right our ports to scan uh list right remember
29:10
ports to scan is is basically the what came back when we actually you know identified okay you give us this this
29:16
from range this two range like I get back a list of integer uh ports to actually scan so I need to actually I'm
29:22
going to range over this feeding in the port numbers right but what why am I doing this in a separate prot here well
29:29
that's a good question right so we're feeding stuff into porch Chan but in order for the porch to get to get
29:35
processed we need to get something back on results Chan right so if we were to in the main go routine if we were to
29:42
feed into porch Chan in the main go routine we would eventually block yes right so we need to do this in a
29:47
background go routine n plus two now background go routine y uh because every
29:53
time we feed something into ports Chan we're going to block and if we go further down on line 59 60 61 62 63
30:00
that's where we're going to receive from the is it called yeah results CH once we receive from results Chan then the
30:07
passing part on line 55 that's going to unblock and we'll be a continue to pass exactly exactly remember we have a
30:13
buffer right that we need to keep you know loose right so as as things are sending in there the moment the buffer
30:19
is full right there can be no more sends right so but we need to be popping things off right of that result chain in
30:26
order for other things to actually be able to go in right so the we we basically it's like a conveyor belt we
30:31
can't have you know a bottleneck anywhere so we need to be able to read from the other end in order for new things to go in so that's that's exactly
30:38
why we have a sort of this you know we have the feeding of the pores happening a separate go routine so that we can
30:44
actually have the the reading off of the result chain happening here as well and again the important thing is we start up
30:49
we schedule in the background the person that puts stuff on to the conveyor Bel the person who's taking stuff off is in
30:55
the main go routine right exactly exactly and at the end here just you know cleaning after ourselves in this case and then we actually print the
31:00
results so let's actually see this this version of our program in action this is the uh number four so let's actually do
31:08
go run uh this is going to be four that go and let's actually give this one the
31:14
same number of ports that sort of tripped us up right uh uh before right remember we just ran a file of the
31:20
operating system said you had too many file Handles open we just have too many grows and trying to do too many things so if we write the same range again here
31:26
so this time right we we're back to having the same kind of errs we getting before that hey the connections you know
31:32
is being refused which we expect that that's our indication that basically the the port is closed but you notice that
31:38
we haven't run a file of the operating system yet we haven't hit any errors that says hey you've got too many file Handles open you know you you you've
31:43
exhausted your your your resources and things like that so we have we have constrained our concurrency you know
31:49
using the worker B pattern we have a set number of workers doing work they're feeding in we're feeding in ports to scan we're picking things up on the
31:55
other end and actually our program now behaves a lot more um stable and it also
32:00
looks like we're actually doing this stuff concurrently not just you know like like the first one where we're just
32:06
going sequential one 2 three four yeah exactly you can look at the port numbers here you know it's completely out of
32:11
sequence right which is exactly what we want all right so let me stop this and oh actually the results and the the
32:17
signal handling we got we got our results back so we we got 4 5432 open so
32:22
now that we've got ourselves um the uh the the worko pool pattern I wanted to
32:27
intr produce a pattern that we don't often see right well so worker pool is like a legit pattern legit pattern you
32:34
can use this in production it's solid absolutely absolutely but there's another one that I I I find to be quite
32:40
elegant right that I don't often see use in a go community and that's a semaphore can can you try and explain succinctly
32:46
what an what the semaphore is I could try I haven't I haven't used semaphor in a long time so I kind of see semaphore
32:54
as similar to a worker pool except it works in batches right it's like you you batch up let's say num CPU units of work
33:02
or ports and once they're all batched up and ready to go then you kind of unlock and you start doing them you start
33:08
executing all eight in your case eight you you execute all eight of them once they're all done you start batching up
33:14
eight more execute them all at once so it's kind of like the opposite you don't stream work and the workers don't just
33:20
pick it up as soon as they can you kind of put everything on the conveyor belt the conveyor belt you stop the conveyor
33:27
belt you get everyone to pull stuff off the conveyor belt do their work and then you put all the the new stuff the new
33:32
eight ports onto the conveyor belt and just rinse and repeat until it's all done right imagine you you and your your
33:38
three other friends right go to a restaurant and and you you need a table that seats four people right um and
33:45
currently the restaurant is a capacity every table maybe it's a Friday night everybody's out having fun and uh and
33:51
and you're basically out in the the the lobby the waiting area because there is there is no table there's a table with
33:57
two people right uh um you know but two open seats right but you need a table
34:03
for four friends you can't the all four of you must be able to go in there and sit down at that table right you can't have sparse you know like you know one
34:09
one spot here you know two spots there you need a table for four people right so to me that's what the semaphor is I
34:15
need all four of uh of us right to go and sit at a table right basically that's our allocation we need all four
34:22
of us to be able to sit in to actually do the work of eating right and talking to each other and having fun with each other right um or or we simply wait
34:29
until that becomes available right so that's the whole point of the sem that's how that's how that sort of that's what that facilitation enables but seor is
34:36
not yet part of the standard Library it's part of the experimental set of packages under goang SLX right but
34:42
that's still you know don't don't let that fool you that that very much you know that I use this in production it's
34:48
it's very much part of the our ecosystem and the part of the standard library or at times these things could of um
34:55
graduate into the standard Library proper but you know you have no fear you know of of actually using this Library
35:00
so let's actually see how we get to actually um use this so I'm just scrolling past a lot of this because
35:06
we've already seen much of this already and let's focus on the uh sort of how we invoke the semore here so this semaphor
35:13
I call in se4 package and I call new weighted so this particular function will give me sort basically the the
35:20
maximum weight right of the semaphor so how many things so what's the capacity right of that restaurant how many how
35:26
many things can how many spots do I actually have and basically now from from then on we're able to say well what
35:33
is the acquisition weight meaning remember when I have my my me and my three other friends we needed to be able
35:39
to sort of you know we need four spots to be go into this um restaurant and actually sit down and eat that's the
35:45
acquisition weight right so there's the maximum weight could be the capacity of the of the restaurant right and there's
35:50
the acquisition weight I need four spots at a time to be able to go in and sit down and actually enjoy my meal right
35:57
that's what these two are and these numbers are completely arbitrary but they work in relation to each other right so from there we basically say hey
36:04
I'm going to uh uh uh iterate through the list of PS to scan right same as we've done before but this time I'm
36:11
going to be using the semaphore to acquire right on the uh using that particular acquisition weight right and
36:18
this call is a blocking call so that means this is us sitting in the lobby right waiting Wai waiting to be able to
36:24
go in to to to sit down and eat right because there's enough spot for us yet so we're waiting right here so as the
36:30
moment we' we've enough uh folks have acquired the acquisition weight right
36:35
the moment that we can no longer extract right that acquisition weight then the choir becomes a blocking call so we saw
36:42
blocking calls with the weight group if you've used mutexes and in go or other languages they're blocking calls in
36:48
there in those languages or with go mutexes as well it can be really useful here yeah absolutely absolutely and when
36:54
when as folks get done in the restaurant and they start to leave right they actually release right which is what
36:59
we're doing here we have a basically a separate goutine in that goutine the actual scanning operation gets done
37:05
where we scan the p and basically we're deferring the release right of the weight so basically imagine there was
37:10
another party of four that just got done eating now they get up they pay their bill hopefully and they they walk out
37:18
right so that's that's their release they've just released right the acquisition weight now you with the same
37:23
acquisition weight right can now go in and sit down at at that table where previously it was occupied right so that
37:29
that whole mechanism so in interest of time I'm going to skip running this one and I'm going to jump to the next um uh
37:35
concept the next uh uh pattern I want us to take a look at and that's the pipeline so when I say pipeline what
37:41
comes to mind for you um kind of like an assembly line to me you know you got one person doing something they hand it off
37:47
another person doing the next thing they hand it off but each individual person they all just do one thing at a time
37:52
they know how to get something from the previous do their thing pass it on to the next right right right so the way I
37:58
like to to sort of visualize this is basically you have all these steps right in and this sort of this conveyor belt
38:04
as as as you illustrate and each one of these steps obviously is is is handled by a go routine right so we have you
38:10
know the first one it gets done with the work you know you know through a channel right communicates to the next step in
38:16
in the process and then again through another Channel communicates with the next step in the process right so let's actually see this pattern right uh at
38:23
play in in in our code so let's flip back
38:28
oh we need [Music] to there we go all right so let's actually see how we actually put one of
38:35
these together so we've already seen much of this before we've seen report to scan up we actually need to open up the
38:43
right the right program all right here we are back in
38:49
Action all right so you know we we've we've we've uh um we got our parts to scan and we're going to this time we're
38:55
going to be writing our our the results to to a file by the way there's going to be a repository with all of this code
39:00
sample so don't you know you're not missing on anything you're going to have you're going to get to see all this and play around with it at your leisure
39:06
right and basically what we're doing here is this is where we we have a pipeline invocation right so the the way
39:12
this invocation is happening it's kind of oddl looking right so starting from the innermost function we have a a
39:18
generator right uh typically when you have a pipeline you need to actually generate there's a there's a starter
39:24
there's a starting point right that generates basically the the the work that needs to be done in the pipeline and that's and that's where this is
39:30
notice the the return type that's a that's a that's a directional channel right for a scan op what is let's take a
39:37
quick look at scan op scan op basically says hey it's an encapsulation that says hey given a port and and we're going to
39:43
find out whether it was open or not if there was an error on this we're going to capture that we're going to capture the duration right pretty simple we're
39:49
just capturing all the details about the the skin operation that's what that's about right so that's going to generate
39:54
basically a channel for us of skin operations that we can read from right to do work okay but the result of that
40:01
right is passed into our scan now our scan is the next step in the process right it's the next step in
40:08
the pipeline right so we're feeding in that that channel into the scan the scan
40:13
then is going to read from that channel as well and in whatever working needs to do right it's going to be doing but its
40:19
result is also right a a Channel of skin operations which gets fit fed into
40:24
filter right so that that conveo belt mechanism right and then the last step in the process is the store which does
40:31
the exact same thing receives that that channel so you can see how the channel becomes the conveyor belt that same
40:37
channel is what basically every step in the process is using to pass information to each other so I noticed something
40:42
here we've got you know the standard people who do the work we've got someone to not do work but just pass stuff on to
40:48
the beginning of the conveyor belt we've got someone to take it off and put it into a file but they're all waiting for
40:54
the previous person we don't have any do we have any more concurrency in this that's the thing by doing this though
40:59
we've actually we've reintroduced s of a sequential processing right of step by step by step by step so the way I'm glad
41:06
you brought this up because the way we can avoid this is actually by using a different pattern called a fin out and
41:12
fin in so again in the interest of time I'm going to let you play with with with all these code samples on your own but
41:18
I'm going to introduce the fin out and Fen in pattern instead and see how that looks a little different right from from
41:26
the pr previous version right so on the let's actually let's put this on the left side let's put the fan out on the
41:34
right side okay okay so this was our original right implementation this is the pipeline right and on the right side
41:40
this is where we actually spinning out right we have we still have a generation step that's going to give us back our
41:46
our channel of operations that we're going to read from but we're going to fin out by basically the scan call is
41:52
basically going to receive basically that channel of scan operations that we can read from but we can invoke it
41:58
multiple times in this case we're doing so manually but you can actually you know you can program around this right you can have a loop that's basically you
42:03
know driven by maybe this the number of workers that you have working on on those different steps in the pipeline
42:09
but we basically we split that work into three different um um channels right uh
42:15
uh to to actually to to actually perform the work thereby bringing back the concurrency that we lost by making
42:21
things sort of linear right with with the frost impation of the pipeline so we got this model now where we started with
42:27
but it looks like it's super easy now to make it concurrent instead of calling scan once now we're just calling it three times or end times or whatever now
42:34
we've just easily pretty easily we have three go routines exactly exactly so I'm am going to close this one right but I
42:41
do I do want to do a quick walk through what the what our scan what some of our pipeline um steps look like so if I go
42:48
to the scan right so oh I didn't mention the the the generator and the scan and
42:55
the filter so these are these are you you'll notice a pattern here which each step of our of our pipeline so the
43:02
generator this is the starting point this is when I said I reive the income the the ports that we need to scan and and returns that channel of scan
43:08
operations that we're going to be reading from so basically we create this this sort of buffer channel right and
43:14
then we feed in right the ports right to to uh the the that outbound channel that
43:20
we're going to feed into other steps that are going to read from that right and we return these things and the
43:25
reason why we are closing here right so we we always talk about basically the channel right closing a channel being a
43:31
signaling mechanism why do you think that's important here that we actually we need to return this this
43:38
this this this Channel and we need to close it here but in other cases we we're doing this sort of um because this
43:45
is the first step in the process the generator right the the the point I'm seeking to make is more more readily
43:50
Illustrated with this perhaps this the scan operation where we we need do need to send back an outbound Channel as
43:57
quickly as possible because we are in this pipeline right but we're actually uh the the we need to actually close out
44:03
this this channel to signal right to the further steps down the road that there's no more data coming yeah right this is a
44:11
great way to broadcast something this is the exactly when you close a a channel you're signaling to anybody who's
44:16
reading on that channel that there is no more data coming right that's why we actually perform this work here inside of a we defer the close of this uh of
44:24
this outbound channel here after having fed in after having done the work whatever work we need to do right sending in the scan into the outbound
44:31
channel right this happens concurrently and we return back the outbound channel so that any other step in the process down the road can actually pick up and
44:38
keep reading and know when there's no more work to be done and so we've kind of got a pattern here create a channel
44:43
yep go use our go keyword to schedule some work yep immediately return the channel after it's scheduled and then
44:48
then we kind of move on to the next step move and move on to the next step exactly exactly so let me actually quickly show
44:57
this version of the program which is our fin
45:06
out I think this is uh nine yep it is nine and we have that and we have the
45:13
ports here and then we run pretty pretty wide range of ports and then in this case here we found the the the one port
45:19
that is open and we can see how sort of again use using that that F out and F in pattern we're able to actually split the
45:25
work across multiple groups routines and actually merge all the steps back in so that we can get ourselves a result and
45:31
our program successfully ends and this is cool because we were really easily able to put filter on the end of our
45:36
pipeline so we don't have to par ex all these things yeah our filter function
45:42
which is exactly what we have here we basically said hey I want to filter out only the open channels right which is another step in our process right for uh
45:49
um in in our pipeline right basically say you only give me back the things that I care about right so one last
45:55
thing we're going to look at before before we before we wrap up here is uh
46:00
the are we looking at number is it number 10 we're looking at uh context I think we're looking at number yeah I
46:06
think we're looking at number 10 because one of the things that that that is useful right in in being able to sort of
46:13
launch all this all this work and being able to sort of uh uh know when when things are um when things are done being
46:20
able to signal to to to go routines that hey you're done you can you can go away now because go routines are not garbage
46:25
collected right they're going to hang around until you you return from them or you say hey I'm not using them anymore
46:31
right so we we're basically augmenting the the uh the um our pipeline by having
46:37
a mechanism of signaling right to to the every step of the pipeline when when we
46:42
no longer need the pipeline to to to be running right and this goes back to our you know always know when your gotin is
46:48
going to shut down always know when the gotin is going to shut down no when to signal for it to to to go away because
46:53
again these things sort of hang around so this is where we're actually using not a context but we're using an actual
46:59
a done Channel which is a very common you know idiom within the go Community basically having that done Channel
47:04
passed around just like you would a context right and this has given me since we've been talking about this has given me an idea how we should actually
47:10
have other another example um to this it actually uses the context right in addition to to the example here that
47:17
uses the dun Channel but this dun channel here you can think of it as the the the signaling mechanism that we hand
47:23
over to every step of the pipeline so that they can be made aware of when they should stop doing the work and so like
47:29
you said this is a really powerful and popular mechanism to broadcast something is done if you use context for doing for
47:36
signaling things or timed out or manually finishing something this is used under the scenes behind the scenes
47:42
behind the scenes absolutely absolutely so you know you you might see some some um some change some sophistication added
47:49
here before we're sort of manually doing our fan out here this is a more sophisticated way of actually
47:55
based on the number of workers creating the number basically doing that same F out work we're doing before and then
48:01
basically we're doing some filtering here p in this time every step of our
48:06
pipeline actually we receive as part of the invocation as part of the arguments
48:11
for for that function invocation it's going to receive the done Channel as a signaling mechanism meaning that if I
48:16
were to go to say the uh let's go to the filter open as one of the steps in our
48:23
in our in our as an examplar right you see that it rece the done Channel as you know as a as a as a readon Channel of
48:31
ever empty strs right the empty struck is the common commonly used for for for for signaling on channels and again we
48:38
have the same thing we're doing before we have the inbalance channel from which we're are actually going to read of
48:43
things things things things to do uh you see I'm telling you it's it's it's been a
48:48
long and then basically we turning the the uh channel of skin operations again again for whatever comes comes at the uh
48:55
next in in the PIP line and we're following the same exact procedure here but the only thing we've introduced here
49:01
that is worthy of note is the fact that we're actually now using a select block right to basically now not only sort of
49:08
doing the work we're doing before but listening in on this done channel so that if we receive a signal right that
49:14
says hey you should you should stop now you should you should be there's nothing more that's going to come through on this pipeline because sometimes you you
49:21
don't need the pipeline to keep operating anymore you know maybe maybe you received an error and and you really want want to sort of bail out of there
49:28
there there's nothing to be done anymore you want to shut things down so receiving a message off that done
49:33
channel is going to be our cue to return from this go routine right so that now
49:40
this go routine goes away the basically you know you can get garbage collected now we can actually you know sort of
49:45
shut down things cleanly and this is a really popular pattern you know production concurrent code exactly if
49:51
you are going to do some operation inside of goutine send or receive on a channel you probably want to put in
49:56
select and you probably want to have a done Channel as well exactly one of the cases in the select exactly so the FI
50:03
final final run of our of our example here what I'm going to do is I'm going to use this filter error uh function
50:08
right and and I'm going to uncomment some things above so so you can actually see this in action I'm going to use this filter function here as as that is part
50:14
of our pipeline to basically say hey if the the the I want to extract out of
50:20
that pipeline as things are going through that conveyor Bel I want to pluck out right the ones the scan
50:25
operations that i a out okay I'm basic I'm going to find the ones that basically when you do the scan and um
50:32
for whatever reason right there was an error but the errors I'm going to be looking for is one where there are too
50:37
many open files remember that that when we're in a foul of of of uh the the OS andc saying hey you're using up too many
50:43
resources when we had too many grow things happening running at the same time I want to pluck those out right and
50:49
the way I'm going to do that is by actually invoking this program you know basically asking for way too many grin
50:54
to be running at any given time lots of workers right so and but that filter function is actually going to give me
51:00
back right the one where I've actually run a file of of of my constraints here and I'm going to return that so let me
51:06
go back onto the top here and then we're GNA we're going to end up some with the with the with summary of the best
51:12
practices for this you know imagine we got different requirements from our PM or from the CTO or something we've got
51:17
our Pipeline and we're essentially going to comment out filter open and uncomment filter error and now we've got different
51:23
Behavior right so we're basically you know the person at the end of the the assembly line we're going to fire one of
51:29
them and we're going to bring in someone new who's going to do slightly different work and now we've got a different program so this is the power to me of of
51:35
pipelines is it's very composable but it's very orthogonal I suppose right
51:41
indeed indeed so the first time I invoked this program I just want to make sure it works not nothing's broken right
51:46
so I'm going to invoke this program as is right this is the one with our workers so it actually did something let
51:52
me make sure I've got my my uh um okay I've got the appropriate number of ports okay okay good good so remember we
51:58
actually the reason why we didn't get an output is because I'm filtering out right I'm only saying give me back the errors that happen right so in this case
52:05
I actually I scanned right basically this default number of ports uh and and there was no error so my program
52:12
actually ran successfully even though there was no output right so let's actually change that a little bit let's actually do this but this time around
52:18
I'm going to provide like ridiculously wide range a lot of file descrip and I'm
52:24
going to ask for a ridiculous number of work workers right let's actually see right if this goes according to how I
52:30
think it's going to go right the operating system is going to tell hey at some point this process has got to way
52:35
too many file Handles open right and our filter error function in our pipeline should be able to pluck that out of that
52:42
stream of Errors you know happening uh as as the scanning is happening our our filter error should be able to pick that
52:47
up let's see that's true it's Friday but we're asking our trusty machine and trusty programming language do extra
52:53
work let don't let me down even though it's almost the weekend ready ready here it goes here it goes
52:58
here it goes oh it's running it's running oh here we go all right to many
53:03
files there we go we actually plugged out right the too many open files and what else did we do when we actually
53:09
caught this right remember when we talking about the whole signaling mechanism so what we did here we were ringing over right plucking values out
53:16
of the filter error um uh step in our pipeline right and the moment we were able to read a value right we decided
53:23
you know what I don't want this program running anymore we got an error we didn't like we didn't we we're bailing
53:29
out of here we're done we send a signaling mechanism to any other steps that might be in the pipeline you know that map perhaps could be wrapping
53:35
filter whatever it is you know we're not doing that here but you can actually see hey if I had other things wrapping filter whatever it is I could send this
53:41
done signal right to to everybody say hey we're out of here we're bailing right and then we move on and then
53:46
basically our program can can now basically shut down cleanly right so these are a lot of of of cool tricks and
53:54
tips right let's see if we can actually summarize this into a handful of of of
54:00
uh of uh key takeaways let's do it all right so the first thing I would say is
54:06
start simple this one start simple I love it you saw us we just used the net. we found out one thing in the St library
54:12
that we needed to use for this we had a you know um a port that we just basically we just hardcoded into the
54:18
program we said hey what's the simplest thing I could work and we basically we could valid the idea and then basically
54:24
build from there we evolved from a V1 that you know we're putting in production on day one we evolved it
54:29
maybe we fast forward it through three years of production experience here you know we evolved it we got more complex
54:35
as we needed to as production bugs came out or as requirements grew
54:40
indeed should we do a one two identify the 1 2 3 embarrassingly
54:46
parallel exactly we to with this oh yeah yeah so yeah not every problem you
54:52
encounter right is going to be suitable right for concurrency right and when you do you need to identify the parts of of
54:59
your code that actually can run right in parallel that that can that's a good candidate for concurrency and a lot of
55:06
that is when they don't have dependencies with each other exactly if there's no dependency from yep from one to the next that's usually a good spot
55:13
right for for for contemplating concurrency what's the next one avoid
55:19
runaway go team yes yes yes yes that that goes back to like always make sure
55:25
you know when you're scheding them at least when you're scheduling them and when they're going to be done exactly
55:30
exactly and and as you saw us Bic saying throwing a bunch of girl routines you know at the task and the operating
55:36
system basically saying ah you're using way too many file Handles in that case that was a constraint that we were
55:41
violating right with our program so you need to understand and what is the context in which my program is going to
55:46
run what are the resources that I'm going to have available right what's the network throughput I'm working with what's the database constraints what's
55:52
the latency whatever the case may be identify the constraints that in within which you're program is actually going to to execute so that when you have gr
55:59
routines doing the work you actually know what what boundaries are they going to bulp up against and another thing to remember if you're coming from cc++ or
56:06
Java or other other operating systems that give you access to OS threads gotin
56:12
are not OS threads right so the operating system is not going to automatically schedule your go routines
56:18
the operating system is not going to automatically limit your go routines the go run time lets you have lots and lots
56:23
of go routines indeed indeed so lever just in a library right we
56:30
didn't need to go far yeah that that was that was pretty much it the only thing that was that was foreign if you want to call it that was the the semore that
56:36
came from the you know Golding X yeah but pretty much everything is is standard Library yeah use the more
56:43
advanced pattern as needed and that's key the last two words there as needed
56:48
we evolved from the simple thing and it kind of goes back to the first bullet point we evolved from the simplest
56:54
solution as we needed to as we found out that well a user might come and test in production and give us 10,000 ports to
57:02
scan well that's not going to work on our server or on our laptop so we need to add more patterns to to manage that
57:07
requirement exactly exactly so the uh repository with all the source code and
57:13
even the things we didn't have time to go over is all available at this URL and uh this is my Twitter handle and this is
57:19
uh errand you know feel free to to to get in touch with us especially those that are new to the community uh we
57:26
welcome you this is this is a safe place and definitely reach out to us um don't hesitate don't don't be don't be shy
57:32
right you know we as hopefully more than happy to talk about if you couldn't tell we love this stuff currency is our Jam
57:39
distributed systems is our Jam so please reach out if if you want to talk about this stuff if you have questions if you
57:44
you know want to work through a problem anything anything absolutely absolutely


