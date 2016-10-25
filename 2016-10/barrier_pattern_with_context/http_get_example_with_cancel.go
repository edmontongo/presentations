package main
import ("context"; "fmt"; "net/http"; "golang.org/x/sync/errgroup")

func main() {
	g, ctx := errgroup.WithContext(context.Background()) // HL

	for _, h := range []string{"a.nonexistent.name", "www.golang.org", "www.google.ca"} {
		req, _ := http.NewRequest("GET", "http://"+h, nil)
		g.Go(func() error { // HL
			fmt.Println("Fetching", h, "...")
			resp, err := http.DefaultClient.Do(req.WithContext(ctx)) // HL
			if err == nil {
				fmt.Printf("OK    %s: succeeded with %q\n", h, resp.Status)
			} else {
				fmt.Printf("ERROR %s: %s\n", h, err)
			}
			return err
		})
	}

	fmt.Println("Global error:", g.Wait())
}
