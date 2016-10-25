package main
import ("fmt"; "net/http"; "golang.org/x/sync/errgroup")

func main() {
	var g errgroup.Group // HL

	for _, h := range []string{"a.nonexistent.name", "www.golang.org", "www.google.ca"} {
		url := "http://" + h
		g.Go(func() error { // HL
			fmt.Println("Fetching", url, "...")
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("OK    %s: succeeded with %q\n", url, resp.Status)
			} else {
				fmt.Printf("ERROR %s: %s\n", url, err)
			}
			return err
		})
	}

	// Wait until all requests have either completed or erred
	fmt.Println("Global error:", g.Wait()) // HL
}
