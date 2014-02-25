package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "net/http"

type Tree struct {
    Value int
    Branches map[string]Tree
}


func TreeSum( tree Tree ) int  {
	sum := tree.Value;
	for _, val := range tree.Branches {
		sum += TreeSum(val)
	}
	return sum;
}

type iInt int

func (max iInt) Iter () <-chan iInt {
	ch := make(chan iInt);
	go func () {
		m := int(max)
		for i := 0; i <= m; i++ {
			ch <- iInt(i)
		}
		close(ch)
	} ();
	return ch
}

func TreeWalker( tree Tree, f func(Tree) )  {
	f(tree)
	for _, val := range tree.Branches {
		TreeWalker(val,f)
	}
}

func TreeIter( tree Tree ) <-chan Tree {
	ch := make(chan Tree);
	go func () {
		TreeWalker( tree, func(t Tree) {
			ch <- t
		})
		close(ch)
	} ();
	return ch
}

func (tree Tree) Iter() <-chan Tree {
	return TreeIter( tree )
}

type StringIterator struct {
	current int
	s []rune
}

type EvenStringIterator struct {
	current int
	s []rune
}


// http://ewencp.org/blog/golang-iterators/
func (si *StringIterator) Next() bool {
	si.current++
	return (si.current < len(si.s))
}

func (si *StringIterator) Value() string {
	return string(si.s[si.current])
}


func Iterator(s string) *StringIterator {
	return &StringIterator{current: -1, s: []rune(s)}
}


func (si *EvenStringIterator) Next() bool {
	si.current += 2;
	return (si.current < len(si.s))
}

func (si *EvenStringIterator) Value() string {
	return string(si.s[si.current])
}


func EvenIterator(s string) *EvenStringIterator {
	return &EvenStringIterator{current: -1, s: []rune(s)}
}


func intIntMap( iarr []int, cb (func(int) int)) []int {
	out := make( []int, len(iarr))
	for i,v := range iarr {
		out[i] = cb( v )
	}
	return out
}

// MACROS??? GENERICS???
func strStrMap( iarr []string, cb (func(string) string)) []string {
	out := make( []string, len(iarr))
	for i,v := range iarr {
		out[i] = cb( v )
	}
	return out
}


func series(start int, end int) []int {
	length := end - start + 1
	out := make( []int, length )
	for i := 0; i < length; i++ {
		out[i] = start + i
	}
	return out
}

func parallelIntIntMap( l []int, f (func(int) int), workers int) []int {
	chans := make( [](chan []int),  workers )
	//var chans [workers]chan []int
	for i := range chans {
		chans[i] = make(chan []int)
	}

	unit := len(l)/workers
	for i := 0 ; i< workers; i++ {
		mychan := chans[i]
		start := i * unit
		end := (i + 1)*unit
		if end >= len(l) {
			end = len(l) 
		}
		subl := l[start:end]
		par := func(l []int) {
			mychan <- intIntMap(l, f)
			close(mychan)
		}
		go par(subl)
	}
	out := make([]int, len(l))
	for i := 0; i < workers; i++ {
		arr := <- chans[i]
		start := i*unit
		end := start + len(arr)
		copy(out[start:end], arr)
	}
	return out
}


func parallelStrStrMap( l []string, f (func(string) string), workers int) []string {
	chans := make( [](chan []string),  workers )
	//var chans [workers]chan []string
	for i := range chans {
		chans[i] = make(chan []string)
	}

	unit := len(l)/workers
	for i := 0 ; i< workers; i++ {
		mychan := chans[i]
		start := i * unit
		end := (i + 1)*unit
		if end >= len(l) {
			end = len(l) 
		}
		subl := l[start:end]
		par := func(l []string) {
			mychan <- strStrMap(l, f)
			close(mychan)
		}
		go par(subl)
	}
	out := make([]string, len(l))
	for i := 0; i < workers; i++ {
		arr := <- chans[i]
		start := i*unit
		end := start + len(arr)
		copy(out[start:end], arr)
	}
	return out
}

func intIntReduce( iarr []int, cb (func(int,int) int)) int {
	o := iarr[0]
	m := len(iarr)
	for i := 1; i < m; i++ {
		o = cb( iarr[i], o)
	}
	return o
}


func parallelIntIntMapReduce( l []int, mapper (func(int) int), reducer (func(int,int) int),workers int) int {
	chans := make( [](chan int),  workers )
	//var chans [workers]chan []int
	for i := range chans {
		chans[i] = make(chan int)
	}

	unit := len(l)/workers
	for i := 0 ; i< workers; i++ {
		mychan := chans[i]
		start := i * unit
		end := (i + 1)*unit
		if end >= len(l) {
			end = len(l) 
		}
		subl := l[start:end]
		par := func(l []int) {
			mychan <- intIntReduce(intIntMap(l, mapper), reducer)
			close(mychan)
		}
		go par(subl)
	}

	reductions := make( []int, workers)
	for i := 0; i < workers; i++ {
		r := <- chans[i]
		reductions[i] = r
	}
	return intIntReduce(reductions, reducer)
}





func main() {
	fmt.Printf("Hello, world.\n")
	condition := false

	for condition {
		// ...
		condition = !condition
	}
	reader := bufio.NewReader(os.Stdin)
	reads := 0
	for {
		reads++
		_, err := reader.ReadString('\n');
		if err!=nil {
			break;
		}
	}
	fmt.Printf("Reads: %d\n", reads)

	x := 10
	for x > 0 {
		x = x - 1
	}
	// x is 0
	fmt.Printf("x is %d\n", x)
	
	// maybe you're not sure how many
	// iterations you need?
	y := 100.0
	for y > 1 {
		y = y / 3;
	}
	fmt.Printf("y is %f\n", y)
	// y is 0.41152263374485604

	sum := 0;
	for i := 0 ; i < 10; i++ {
		sum += i;
	}
	fmt.Printf("Sum is %d\n", sum)

	s := ""
	s2 := "❤☀☆☂☻♞☯☭☢"
	v := []string{"a","b","c"}
	u := map[string]string{"A":"a", "B":"b","C":"c"}
	for i, val := range v {
		for j := 0; j <= i; j++ {
			s += val; 
		}
	}
	fmt.Printf("s is [%s]\n",s)
	for key, val := range u {
		s += key;
		s += val; //  over keys
	}
	fmt.Printf("s is [%s]\n",s)
	for i, val := range s2 {
		// i is the byte location
		for j := 0; j <= i; j++ {
			s += string(val);
		}
	}
	fmt.Printf("s is (note that the unique was base-4) [%s]\n",s)

	stump := map[string]Tree{}	

	tree := Tree{0,
		map[string]Tree{
			"a":Tree{1,stump},
			"b":Tree{2, map[string]Tree{
					"h":Tree{8,stump},
					"i":Tree{9,stump},
					"j":Tree{10,stump},
				},
			},
			"c":Tree{3,stump},
			"d":Tree{4, map[string]Tree{
					"e":Tree{5,stump},
					"f":Tree{6,stump},
					"g":Tree{7,stump},
				},
			},
		},
	}	

	fmt.Printf("Treesum %d\n", TreeSum( tree ))

	for i := range iInt(6).Iter() {
		fmt.Printf("Wow! %v\n", i)
	}

	// Call Back walker
	TreeWalker(tree, func(t Tree) { fmt.Printf("Node value %d\n",t.Value) } )

	for tree := range tree.Iter() {
		fmt.Printf("Now via Iter Node value %d\n",tree.Value)
	}

	si := Iterator(s)
	for si.Next() {
		fmt.Printf("String val! %s\n", si.Value())
	}

	si2 := EvenIterator(s)
	for si2.Next() {
		fmt.Printf("Even String val! %s\n", si2.Value())
	}


	v2 := []int{1,2,3,4,5,6,7,8}
	inc := func(x int) int { return 1 + x }
	sqr := func(x int) int { return x * x }
	// lack of generics
	v3 := intIntMap(v2, inc)
	fmt.Printf("inc v2: [%v] v3: [%v]\n",v2,v3)
	// lack of generics
	v3 = intIntMap(v2, sqr)
	fmt.Printf("sqr v2: [%v] v3: [%v]\n",v2,v3)
	

	basename := func(path string) string {
		sp := strings.Split(path,"/")
		return(sp[len(sp) - 1])
	}
	vs := []string{"/home","/file", "/usr/local"}
	vs2 := strStrMap( vs, basename )
	fmt.Printf("basename vs: [%v] vs2: [%v]\n",vs,vs2)

	urls := []string{"http://cbc.ca", "http://gc.ca", "http://alberta.ca"}	
	status := func( uri string ) string {
		resp, _ := http.Get(uri)		
		return(resp.Status)
	}

	statuses := strStrMap(urls, status);
	fmt.Printf("statuses: %v\n", statuses)
	
	mySum := func(l []int) int {
		sum := 0
		for _,v := range l {
			sum += v
		}
		return(sum)
	}
	
	sumres := mySum( intIntMap(series(1,1000),sqr))
	fmt.Printf("sumres: %v\n", sumres)
	psumres := mySum(parallelIntIntMap( series(1,1000), sqr, 4))
	fmt.Printf("psumres: %v\n", psumres)
	
	pgets := parallelStrStrMap( urls, status, 3)
	fmt.Printf("Stupid examples with URLs %v\n", pgets)

	vex := []int{1,2,3,4,5,6,7,8,9};
	iadd := func(x, y int) int { return(x + y) }
	vexr := intIntReduce(vex, iadd)
	fmt.Printf("Sum: %v\n", vexr)

	preduced := parallelIntIntMapReduce( series(1,100000000), sqr, iadd, 4)
	fmt.Printf("Reduced %v\n", preduced)
	
	
}

