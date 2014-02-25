# while loop

var condition = 1;
while (condition) {
    // ...
    condition = !condition;
}
// condition is False

//how many OKs ?
var count = 0;
while(confirm("OK???")) {
    count++;
}
// alert(count);

var x = 10;
while (x > 0) {
    x--;
}
// x is 0

// maybe you're not sure how many
// iterations you need?
var x = 100.0;
while (x > 1) {
    x /= 3;
}
// x is 0.41152263374485604

// for loop
var sum = 0;
for (var i = 0 ; i < 10; i++) {
    sum += i;
}
// sum is 45

function range(start,end) {
    var out = [];
    var j = 0;
    for (var i = start; i < end; i++) {
        out[j++] = i;
    }
    return out;
}
//alert(range(1,10).length);

for var i in range(1,10) {
    alert(i);
}

var s = "";
var v = ["a","b","c"];
var u = {"A":"a", "B":"b", "C":"c"};
for (var i in v) {
    s += v[i]; // iterates over keys
}
for (var i in u) {
    s += u[i]; // iterates over keys
}
alert(s);
// s starts with abc

// recursive way
function recsum(l) {
    if (l.length > 0) {
        var v = l.pop();
        return v + recsum(l);
    }
    return 0;
}
alert(recsum([1,2,3,4]));

var tree = {"a":{"b":{"c":1,"d":2,"e":3,"f":4},"g":5},"h":{"i":6}};
function treesum(tree) {
    if (typeof tree === "number") {
        return tree;
    } else {
        var sum = 0;
        for (var key in tree) {
            sum += treesum(tree[key]);
        }
        return sum;
    }
}
alert(treesum(tree));


// OO Iterator way
OnlyEvens = function(seq) {
    this.seq = seq;
    this.index = 0;
    self = this;
    this.hasNext = function() {
        return self.index < self.seq.length
    }
    this.next = function() {
        var v = self.seq[self.index];
        self.index += 2;
        return v;
    }
};
var oe = new OnlyEvens([0,1,2,3,4,5,6,7,8,9,10]);
var s = "";
while (oe.hasNext()) {
    s += oe.next();
}
alert(s);


// add 1 to a list
var v = [1,2,3,4,5];
function inc(x) { return 1 + x; }
var u = v.map(inc);
// alert(u); 2,3,4,5,6
// alert(v); 1,2,3,4,5

function basename(path) {
    var sp = path.split("/");
    return sp[sp.length - 1];
}
var v = ["/home","/file","/usr/local"];
var u = v.map(basename);
// u = ['home','file','local']


function GET(url) {
    var request = new XMLHttpRequest();
    request.open("GET", url, false);
    request.send(null);
    return request.responseText;
}
var urls = ["http://cbc.ca","http://gc.ca","http://alberta.ca"];
var pages = urls.map(GET);

var v = [1,2,3,4,5,6,7,8,9];
alert(v.reduce(function(x,y) { return x + y }));
function sum(l) {
    return l.reduce(function(x,y) {return x + y});
}
alert(sum(v));

function sqr(v) { return v * v; }
alert(sum(range(1,100).map(sqr)));
var p = new Parallel(range(1,100));//this could mean 100 workers!
p.map(sqr).then(function(d) { alert("what"+sum(d)); });



// this is why you want blocks with
// few dependencies!
// http://adambom.github.io/parallel.js/
<script src="parallel.js"></script>
<script>

function sum(l) {
    var sum = 0.0;
    for (var i in l) {
	sum += l[i];
    }
    return sum;
}
function range(start,end) {
    var out = [];
    var j = 0;
    for (var i = start; i < end; i++) {
        out[j++] = i;
    }
    return out;
}
var p = new Parallel([range(1,10000),range(10001,20000),range(20001,30000)]);
p.map(sum).reduce(sum).then(alert);

function GET(url) {
    var request = new XMLHttpRequest();
    request.open("GET", url, false);
    request.send(null);
    return request.responseText;
}
var urls = ["map.html","map.html","map.html"];
alert(urls.map(GET));
// You need proper URLs and permission!
var urls = ["file:///./map.html","file:///./map.html","file:///./map.html"];
var urls = ["http://cbc.ca","http://gc.ca","http://alberta.ca"];
new Parallel(urls).map(GET).then(alert);

</script>
